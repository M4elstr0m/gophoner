package facebook

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/M4elstr0m/gophoner/internal/modules"
	_http "github.com/M4elstr0m/gophoner/internal/utils/http"
	"github.com/PuerkitoBio/goquery"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/charmbracelet/log"
)

type eqmcData struct {
	U string `json:"u"`
	L string `json:"l"`
}

type searchSession struct {
	LSD     string
	Jazoest string
}

func fetchSearchSession(browserProfile *browser_profiles.BrowserProfile, client tls_client.HttpClient) (*searchSession, error) {
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

	doc, err := goquery.NewDocumentFromReader(getResp.Body)
	if err != nil {
		log.Warn("Reader's data could not be parsed as HTML", "error", err)
		return nil, err
	}

	raw := doc.Find("script#__eqmc").First().Text()
	if raw == "" {
		err := errors.New("__eqmc script tag was not found on login page")
		log.Warn("Failed to locate session data", "error", err)
		return nil, err
	}

	var data eqmcData
	if err := json.Unmarshal([]byte(raw), &data); err != nil {
		log.Warn("Failed to parse __eqmc data", "error", err)
		return nil, err
	}

	beaconURL, err := url.Parse(data.U)
	if err != nil {
		log.Warn("Failed to parse embedded beacon URL", "error", err)
		return nil, err
	}

	jazoest := beaconURL.Query().Get("jazoest")
	if data.L == "" || jazoest == "" {
		err := errors.New("session data was missing lsd or jazoest")
		log.Warn("Failed to extract session tokens", "error", err)
		return nil, err
	}

	return &searchSession{LSD: data.L, Jazoest: jazoest}, nil
}

func newUUIDv4() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	buf[6] = (buf[6] & 0x0f) | 0x40
	buf[8] = (buf[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:16]), nil
}

type searchParams struct {
	CipherText     *string `json:"cipher_text"`
	Context        string  `json:"context"`
	EventRequestID string  `json:"event_request_id"`
	FriendName     string  `json:"friend_name"`
	SearchQuery    string  `json:"search_query"`
	WaterfallID    string  `json:"waterfall_id"`
}

type searchVariables struct {
	Params searchParams `json:"params"`
}

type accountSearchResult struct {
	Accounts        []json.RawMessage `json:"accounts"`
	NumResultsShown int               `json:"num_results_shown"`
}

type graphqlResponse struct {
	Data struct {
		CaaArFbAccountSearch *accountSearchResult `json:"caa_ar_fb_account_search"`
	} `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

func submitAccountSearch(
	browserProfile *browser_profiles.BrowserProfile,
	client tls_client.HttpClient,
	session *searchSession,
	phoneNumber string,
) (bool, error) {
	eventRequestID, err := newUUIDv4()
	if err != nil {
		log.Warn("Failed to generate event request id", "error", err)
		return false, err
	}
	waterfallID, err := newUUIDv4()
	if err != nil {
		log.Warn("Failed to generate waterfall id", "error", err)
		return false, err
	}

	variables := searchVariables{
		Params: searchParams{
			CipherText:     nil,
			Context:        "recover",
			EventRequestID: eventRequestID,
			FriendName:     "",
			SearchQuery:    phoneNumber,
			WaterfallID:    waterfallID,
		},
	}
	variablesJSON, err := json.Marshal(variables)
	if err != nil {
		log.Warn("Failed to encode search variables", "error", err)
		return false, err
	}

	form := url.Values{}
	form.Set("fb_api_caller_class", "RelayModern")
	form.Set("fb_api_req_friendly_name", SEARCH_FRIENDLY_NAME)
	form.Set("variables", string(variablesJSON))
	form.Set("doc_id", SEARCH_DOC_ID)
	form.Set("server_timestamps", "true")
	form.Set("lsd", session.LSD)
	form.Set("jazoest", session.Jazoest)
	form.Set("__user", "0")
	form.Set("__a", "1")

	postReq, err := _http.NewBrowserRequest(
		_http.BrowserRequestOption{
			Method:         http.MethodPost,
			TargetURL:      GRAPHQL_URL,
			Body:           strings.NewReader(form.Encode()),
			BrowserProfile: browserProfile,
			AdditionalHeaders: [][2]string{
				{"Content-Type", "application/x-www-form-urlencoded"},
				{"X-FB-Friendly-Name", SEARCH_FRIENDLY_NAME},
				{"X-FB-LSD", session.LSD},
				{"X-ASBD-ID", ASBD_ID},
				{"Origin", "https://www.facebook.com"},
				{"Referer", LOGIN_PAGE_URL},
			},
		},
	)
	if err != nil {
		log.Warn("Failed to create POST request", "error", err)
		return false, err
	}

	postResp, err := client.Do(postReq)
	if err != nil {
		log.Warn("Failed to submit account search", "error", err)
		return false, err
	}
	defer postResp.Body.Close()

	if postResp.StatusCode == http.StatusTooManyRequests {
		err := fmt.Errorf("service returned 429: %w", modules.ErrorLimitReached)
		log.Warn("Request denied by rate limit", "error", err)
		return false, err
	}

	bodyBytes, err := io.ReadAll(postResp.Body)
	if err != nil {
		log.Warn("Failed to read search response body", "error", err)
		return false, err
	}

	var result graphqlResponse
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Warn("Failed to decode search response", "error", err)
		return false, err
	}

	if len(result.Errors) > 0 {
		err := fmt.Errorf("graphql returned an error: %s: %w", result.Errors[0].Message, modules.ErrorLimitReached)
		log.Warn("Account search returned errors", "message", result.Errors[0].Message, "error", err)
		return false, err
	}

	if result.Data.CaaArFbAccountSearch == nil {
		err := fmt.Errorf("unexpected graphql response shape: %w", modules.ErrorLimitReached)
		log.Warn("Unexpected account search response", "body", string(bodyBytes), "error", err)
		return false, err
	}

	return result.Data.CaaArFbAccountSearch.NumResultsShown > 0, nil
}
