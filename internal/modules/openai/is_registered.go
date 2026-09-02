package openai

import (
	"github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/platform"
	"github.com/M4elstr0m/gophoner/internal/modules"
	_chrome "github.com/M4elstr0m/gophoner/internal/utils/chrome"
)

const LOGIN_PAGE_URL string = "https://auth.openai.com/log-in"

func IsRegistered(phoneNumber string) modules.IsRegisteredFuncOutput {
	desktopPlatformsOnlyFilter := browser_profiles.PoolFilter{
		AllowedFamilySlice:   nil,
		AllowedPlatformSlice: platform.DesktopPlatforms,
		AllowedLanguageSlice: nil,
	}
	browserProfile := browser_profiles.Random(&desktopPlatformsOnlyFilter)

	ctx, cancel := _chrome.NewContext(browserProfile, modules.HTTP_CLIENT_TIMEOUT_SECONDS, true)
	defer cancel()

	return modules.NewIsRegisteredFuncOutput(checkPhoneRegistration(ctx, phoneNumber))
}
