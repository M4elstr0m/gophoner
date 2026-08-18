package _http

import (
	"errors"
	"io"

	browser_profiles "github.com/M4elstr0m/gophoner/internal/browser_profiles"
	http "github.com/bogdanfinn/fhttp"
	"github.com/charmbracelet/log"
)

type BrowserRequestOption struct {
	Method            string
	TargetURL         string
	Body              io.Reader
	BrowserProfile    *browser_profiles.BrowserProfile
	AdditionalHeaders [][2]string
}

func NewBrowserRequest(options BrowserRequestOption) (*http.Request, error) {
	if options.BrowserProfile == nil {
		log.Error("nil browser profile is not valid for new browser request, please contact M4elstr0m on Github")
		return nil, errors.New("nil browser profile is not valid for new browser request")
	}

	req, err := http.NewRequest(options.Method, options.TargetURL, options.Body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", options.BrowserProfile.UserAgent)
	if options.BrowserProfile.SecCHUA != "" {
		req.Header.Set("Sec-CH-UA", options.BrowserProfile.SecCHUA)
		req.Header.Set("Sec-CH-UA-Mobile", options.BrowserProfile.SecCHUAMobile)
		req.Header.Set("Sec-CH-UA-Platform", options.BrowserProfile.SecCHUAPlatform)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", options.BrowserProfile.AcceptLanguage)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")

	if options.AdditionalHeaders != nil {
		for _, pair := range options.AdditionalHeaders {
			req.Header.Set(pair[0], pair[1])
		}
	}

	return req, nil
}
