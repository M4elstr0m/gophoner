package microsoft

import (
	"github.com/charmbracelet/log"

	browser_profiles "github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/platform"
	_http "github.com/M4elstr0m/gophoner/internal/utils/http"
)

const LOGIN_PAGE_URL string = "https://login.live.com/oauth20_authorize.srf?client_id=3fa91358-6f74-4525-b5df-da149652be36&scope=openid+profile+User.Read+email+offline_access&redirect_uri=https%3a%2f%2fwww.linkedin.com%2fmicrosoft-login%2fhandler&response_type=code&response_mode=form_post&msproxy=1&issuer=mso&tenant=consumers&ui_locales=en-GB"

func IsRegistered(phoneNumber string) (bool, error) {
	desktopPlatformsOnlyFilter := browser_profiles.PoolFilter{
		AllowedFamilySlice:   nil,
		AllowedPlatformSlice: platform.DesktopPlatforms,
		AllowedLanguageSlice: nil,
	}
	browserProfile := browser_profiles.Random(&desktopPlatformsOnlyFilter)

	client, err := _http.NewTlsClient(browserProfile, 15)
	if err != nil {
		log.Warn("Failed to initialize TLS client", "error", err)
		return false, err
	}

	session, err := fetchLoginSession(browserProfile, client)
	if err != nil {
		return false, err
	}

	isRegistered, err := checkCredentialType(browserProfile, client, session, phoneNumber)
	if err != nil {
		return false, err
	}

	return isRegistered, nil
}
