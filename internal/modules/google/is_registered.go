package google

import (
	"time"

	"github.com/M4elstr0m/gophoner/internal/browser_profiles"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/platform"
)

const LOGIN_PAGE_URL string = "https://accounts.google.com/"

func IsRegistered(phoneNumber string) (bool, error) {
	desktopPlatformsOnlyFilter := browser_profiles.PoolFilter{
		AllowedFamilySlice:   nil,
		AllowedPlatformSlice: platform.DesktopPlatforms,
		AllowedLanguageSlice: nil,
	}
	browserProfile := browser_profiles.Random(&desktopPlatformsOnlyFilter)

	ctx, cancel := newChromeContext(browserProfile, 45*time.Second)
	defer cancel()

	return checkIdentifierRegistration(ctx, phoneNumber)
}
