package amazon

import (
	"github.com/charmbracelet/log"

	browser_profiles "github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/platform"
	"github.com/M4elstr0m/gophoner/internal/modules"
	_http "github.com/M4elstr0m/gophoner/internal/utils/http"
)

const (
	LOGIN_PAGE_URL       string = "https://www.amazon.com/ap/signin?openid.return_to=https%3A%2F%2Fwww.amazon.com%2F%3Fref_%3Dnav_ya_signin&openid.identity=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.assoc_handle=usflex&openid.mode=checkid_setup&openid.claimed_id=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0%2Fidentifier_select&openid.ns=http%3A%2F%2Fspecs.openid.net%2Fauth%2F2.0"
	LOGIN_INPUT_SELECTOR string = "input#ap_email_login"

	NO_ACCOUNT_DOC_TITLE  string = "Amazon Intent Confirmation"
	HAS_ACCOUNT_DOC_TITLE string = "Amazon Sign-In"

	NO_ACCOUNT_TOKEN string = "Looks like you're new to Amazon"
)

func IsRegistered(phoneNumber string) modules.IsRegisteredFuncOutput {
	desktopPlatformsOnlyFilter := browser_profiles.PoolFilter{
		AllowedFamilySlice:   nil,
		AllowedPlatformSlice: platform.DesktopPlatforms,
		AllowedLanguageSlice: nil,
	}
	browserProfile := browser_profiles.Random(&desktopPlatformsOnlyFilter)

	client, err := _http.NewTlsClient(browserProfile, modules.HTTP_CLIENT_TIMEOUT_INT_SECONDS)
	if err != nil {
		log.Warn("Failed to initialize TLS client", "error", err)
		return modules.MinimalIsRegisteredFuncOutput(false, err)
	}

	doc, postUrl, err := fetchLoginPage(browserProfile, client)
	if err != nil {
		return modules.MinimalIsRegisteredFuncOutput(false, err)
	}

	form, postUrlString, err := resolveLoginForm(doc, phoneNumber, postUrl)
	if err != nil {
		return modules.MinimalIsRegisteredFuncOutput(false, err)
	}

	resultDoc, err := submitLoginForm(
		browserProfile,
		postUrlString,
		form,
		client,
	)
	if err != nil {
		return modules.MinimalIsRegisteredFuncOutput(false, err)
	}

	isRegistered, err := parseResultPage(resultDoc)
	if err != nil {
		return modules.MinimalIsRegisteredFuncOutput(false, err)
	}

	return modules.MinimalIsRegisteredFuncOutput(isRegistered, nil)
}
