package facebook

import (
	"github.com/charmbracelet/log"

	browser_profiles "github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/platform"
	"github.com/M4elstr0m/gophoner/internal/modules"
	_http "github.com/M4elstr0m/gophoner/internal/utils/http"
)

const (
	LOGIN_PAGE_URL string = "https://www.facebook.com/login/identify/"
	GRAPHQL_URL    string = "https://www.facebook.com/api/graphql/"

	SEARCH_DOC_ID        string = "28496659306608697"
	SEARCH_FRIENDLY_NAME string = "CAAFBAccountSearchViewQuery"
	ASBD_ID              string = "359341"
)

func IsRegistered(phoneNumber string) (bool, error) {
	desktopPlatformsOnlyFilter := browser_profiles.PoolFilter{
		AllowedFamilySlice:   nil,
		AllowedPlatformSlice: platform.DesktopPlatforms,
		AllowedLanguageSlice: nil,
	}
	browserProfile := browser_profiles.Random(&desktopPlatformsOnlyFilter)

	client, err := _http.NewTlsClient(browserProfile, modules.HTTP_CLIENT_TIMEOUT_INT_SECONDS)
	if err != nil {
		log.Warn("Failed to initialize TLS client", "error", err)
		return false, err
	}

	session, err := fetchSearchSession(browserProfile, client)
	if err != nil {
		return false, err
	}

	isRegistered, err := submitAccountSearch(browserProfile, client, session, phoneNumber)
	if err != nil {
		return false, err
	}

	return isRegistered, nil
}
