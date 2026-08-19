package microsoft

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"

	"github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/M4elstr0m/gophoner/internal/modules"
	_http "github.com/M4elstr0m/gophoner/internal/utils/http"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/charmbracelet/log"
)

type serverData struct {
	UrlGetCredentialType string `json:"urlGetCredentialType"`
	SFTTag               string `json:"sFTTag"`
	SUnauthSessionID     string `json:"sUnauthSessionID"`
}

type loginSession struct {
	CredentialTypeURL string
	FlowToken         string
	UAID              string
}

var ppftValuePattern = regexp.MustCompile(`value="([^"]+)"`)

func extractServerData(body []byte) ([]byte, error) {
	idx := bytes.Index(body, []byte("var ServerData ="))
	if idx == -1 {
		return nil, errors.New("ServerData assignment was not found on login page")
	}

	start := bytes.IndexByte(body[idx:], '{')
	if start == -1 {
		return nil, errors.New("ServerData object start was not found")
	}
	start += idx

	depth := 0
	inString := false
	escaped := false
	for i := start; i < len(body); i++ {
		c := body[i]
		if inString {
			switch {
			case escaped:
				escaped = false
			case c == '\\':
				escaped = true
			case c == '"':
				inString = false
			}
			continue
		}

		switch c {
		case '"':
			inString = true
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return body[start : i+1], nil
			}
		}
	}

	return nil, errors.New("ServerData object was not properly closed")
}

func fetchLoginSession(browserProfile *browser_profiles.BrowserProfile, client tls_client.HttpClient) (*loginSession, error) {
	getReq, err := _http.NewBrowserRequest(
		_http.BrowserRequestOption{
			Method:         http.MethodGet,
			TargetURL:      LOGIN_PAGE_URL,
			BrowserProfile: browserProfile,
		},
	)
	if err != nil {
		log.Warn("Failed to create GET request", "error", err)
		return nil, err
	}

	getResp, err := client.Do(getReq)
	if err != nil {
		log.Warn("Failed to get login page", "error", err)
		return nil, err
	}
	defer getResp.Body.Close()

	bodyBytes, err := io.ReadAll(getResp.Body)
	if err != nil {
		log.Warn("Failed to read login page body", "error", err)
		return nil, err
	}

	rawServerData, err := extractServerData(bodyBytes)
	if err != nil {
		log.Warn("Failed to locate ServerData", "error", err)
		return nil, err
	}

	var data serverData
	if err := json.Unmarshal(rawServerData, &data); err != nil {
		log.Warn("Failed to parse ServerData", "error", err)
		return nil, err
	}

	ppft := ppftValuePattern.FindSubmatch([]byte(data.SFTTag))
	if ppft == nil {
		err := errors.New("PPFT token was not found in ServerData.sFTTag")
		log.Warn("Failed to extract flow token", "error", err)
		return nil, err
	}

	return &loginSession{
		CredentialTypeURL: data.UrlGetCredentialType,
		FlowToken:         string(ppft[1]),
		UAID:              data.SUnauthSessionID,
	}, nil
}

type credentialTypeRequest struct {
	CheckPhones                    bool   `json:"checkPhones"`
	Country                        string `json:"country"`
	FederationFlags                int    `json:"federationFlags"`
	FlowToken                      string `json:"flowToken"`
	ForceOtcLogin                  bool   `json:"forceotclogin"`
	IsCookieBannerShown            bool   `json:"isCookieBannerShown"`
	IsExternalFederationDisallowed bool   `json:"isExternalFederationDisallowed"`
	IsFederationDisabled           bool   `json:"isFederationDisabled"`
	IsFidoSupported                bool   `json:"isFidoSupported"`
	IsOtherIdpSupported            bool   `json:"isOtherIdpSupported"`
	IsReactLoginRequest            bool   `json:"isReactLoginRequest"`
	IsRemoteConnectSupported       bool   `json:"isRemoteConnectSupported"`
	IsRemoteNGCSupported           bool   `json:"isRemoteNGCSupported"`
	IsSignup                       bool   `json:"isSignup"`
	OriginalRequest                string `json:"originalRequest"`
	OtcLoginDisallowed             bool   `json:"otclogindisallowed"`
	UAID                           string `json:"uaid"`
	Username                       string `json:"username"`
}

type credentialTypeResponse struct {
	IfExistsResult int `json:"IfExistsResult"`
}

const (
	ifExistsAccountFound    int = 0
	ifExistsAccountNotFound int = 1
)

func checkCredentialType(
	browserProfile *browser_profiles.BrowserProfile,
	client tls_client.HttpClient,
	session *loginSession,
	phoneNumber string,
) (bool, error) {
	payload := credentialTypeRequest{
		CheckPhones:          true,
		Country:              "",
		FederationFlags:      3,
		FlowToken:            session.FlowToken,
		IsReactLoginRequest:  true,
		IsRemoteNGCSupported: true,
		UAID:                 session.UAID,
		Username:             phoneNumber,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		log.Warn("Failed to encode credential-type request", "error", err)
		return false, err
	}

	postReq, err := _http.NewBrowserRequest(
		_http.BrowserRequestOption{
			Method:         http.MethodPost,
			TargetURL:      session.CredentialTypeURL,
			Body:           bytes.NewReader(body),
			BrowserProfile: browserProfile,
			AdditionalHeaders: [][2]string{
				{"Content-Type", "application/json; charset=UTF-8"},
			},
		},
	)
	if err != nil {
		log.Warn("Failed to create POST request", "error", err)
		return false, err
	}

	postResp, err := client.Do(postReq)
	if err != nil {
		log.Warn("Failed to submit credential-type check", "error", err)
		return false, err
	}
	defer postResp.Body.Close()

	if postResp.StatusCode == http.StatusTooManyRequests {
		err := fmt.Errorf("service returned 429: %w", modules.ErrorLimitReached)
		log.Warn("Request denied by rate limit", "error", err)
		return false, err
	}

	var result credentialTypeResponse
	if err := json.NewDecoder(postResp.Body).Decode(&result); err != nil {
		log.Warn("Failed to decode credential-type response", "error", err)
		return false, err
	}

	switch result.IfExistsResult {
	case ifExistsAccountFound:
		return true, nil
	case ifExistsAccountNotFound:
		return false, nil
	default:
		err := fmt.Errorf("unexpected IfExistsResult value %d: %w", result.IfExistsResult, modules.ErrorLimitReached)
		log.Warn("Unexpected credential-type result", "value", result.IfExistsResult, "error", err)
		return false, err
	}
}
