package _http

import (
	browser_profiles "github.com/M4elstr0m/gophoner/internal/browser_profiles"
	tls_client "github.com/bogdanfinn/tls-client"
)

func NewTlsClient(profile *browser_profiles.BrowserProfile, timeout int) (tls_client.HttpClient, error) {
	jar := tls_client.NewCookieJar()
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(),
		tls_client.WithClientProfile(*profile.TlsProfile),
		tls_client.WithTimeoutSeconds(timeout),
		tls_client.WithCookieJar(jar),
	)
	if err != nil {
		return nil, err
	}

	return client, nil
}
