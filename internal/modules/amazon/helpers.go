package amazon

import (
	"bytes"
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

func fetchLoginPage(browserProfile *browser_profiles.BrowserProfile, client tls_client.HttpClient) (*goquery.Document, *url.URL, error) {
	getReq, err := _http.NewBrowserRequest(
		_http.BrowserRequestOption{
			Method:         http.MethodGet,
			TargetURL:      LOGIN_PAGE_URL,
			BrowserProfile: browserProfile,
		},
	)
	if err != nil {
		log.Warn("Failed to create GET request", "error", err)
		return nil, nil, err
	}

	getResp, err := client.Do(getReq)
	if err != nil {
		log.Warn("Failed to get login page", "error", err)
		return nil, nil, err
	}
	defer getResp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(getResp.Body)
	if err != nil {
		log.Warn("Reader's data could not be parsed as HTML", "error", err)
		return nil, nil, err
	}

	return doc, getResp.Request.URL, nil
}

func resolveLoginForm(doc *goquery.Document, phoneNumber string, postUrl *url.URL) (*url.Values, string, error) {
	// gathers all inputs available in the form
	form := url.Values{}
	doc.Find("form input").Each(func(_ int, s *goquery.Selection) {
		if name, ok := s.Attr("name"); ok {
			value, _ := s.Attr("value")
			form.Set(name, value)
		}
	})

	loginField := doc.Find(LOGIN_INPUT_SELECTOR)
	loginFieldName, ok := loginField.Attr("name")
	if !ok {
		err := errors.New("login input field selector was not found or has no name attribute")
		log.Warn("Failed to resolve login input field", "selector", LOGIN_INPUT_SELECTOR, "error", err)
		return nil, "", err
	}

	form.Set(loginFieldName, phoneNumber)

	// resolves the form submit target
	var postUrlString string = postUrl.String()
	action, ok := doc.Find("form").First().Attr("action")
	if ok && action != "" {
		parsed, err := url.Parse(action)
		if err == nil {
			postUrlString = postUrl.ResolveReference(parsed).String()
		}
	}

	return &form, postUrlString, nil
}

func submitLoginForm(
	browserProfile *browser_profiles.BrowserProfile,
	postUrlString string,
	form *url.Values,
	client tls_client.HttpClient,
) (
	*goquery.Document,
	error,
) {
	postReq, err := _http.NewBrowserRequest(
		_http.BrowserRequestOption{
			Method:         http.MethodPost,
			TargetURL:      postUrlString,
			Body:           strings.NewReader(form.Encode()),
			BrowserProfile: browserProfile,
			AdditionalHeaders: [][2]string{
				{"Content-Type", "application/x-www-form-urlencoded"},
			},
		},
	)
	if err != nil {
		log.Warn("Failed to create POST request", "error", err)
		return nil, err
	}

	postResp, err := client.Do(postReq)
	if err != nil {
		log.Warn("Failed to submit login information", "error", err)
		return nil, err
	}
	defer postResp.Body.Close()

	bodyBytes, err := io.ReadAll(postResp.Body)
	if err != nil {
		log.Warn("Failed to read POST response body", "error", err)
		return nil, err
	}

	log.Debug("Successfully collected response body",
		"status", postResp.Status,
		"new-url", postResp.Request.URL.String(),
	)
	if postResp.StatusCode == http.StatusTooManyRequests {
		err := fmt.Errorf("service returned 429: %w", modules.ErrorLimitReached)
		log.Warn("Request denied by rate limit", "error", err)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bodyBytes))
	if err != nil {
		log.Warn("Reader's data could not be parsed as HTML", "error", err)
		return nil, err
	}

	return doc, nil
}

func parseResultPage(resultDoc *goquery.Document) (bool, error) {
	resultDoc.Find("script, noscript, style").Remove()

	var docTitle string = strings.TrimSpace(resultDoc.Find("title").Text())
	if docTitle != HAS_ACCOUNT_DOC_TITLE && docTitle != NO_ACCOUNT_DOC_TITLE {
		err := fmt.Errorf("endpoint page title is different than expected: %w",
			modules.ErrorLimitReached,
		)
		log.Warn("Endpoint page title is different than expected",
			"expected", HAS_ACCOUNT_DOC_TITLE+"||"+NO_ACCOUNT_DOC_TITLE,
			"received", docTitle,
		)
		return false, err
	}

	text := resultDoc.Text()
	text = strings.Join(strings.Fields(text), " ")

	return !strings.Contains(text, NO_ACCOUNT_TOKEN), nil
}
