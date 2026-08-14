package browser_profiles

import (
	"github.com/M4elstr0m/gophoner/src/browser-profiles/enums/family"
	"github.com/M4elstr0m/gophoner/src/browser-profiles/enums/language"
	"github.com/M4elstr0m/gophoner/src/browser-profiles/enums/platform"
	gp_slices "github.com/M4elstr0m/gophoner/src/utils/slices"
	"github.com/charmbracelet/log"
)

type PoolFilter struct {
	allowedFamilySlice   []family.Family
	allowedPlatformSlice []platform.Platform
	allowedLanguageSlice []language.Language
}

var latest *BrowserProfile = nil

func New(builder browserProfileBuilder) *BrowserProfile {
	latest = builder.build()
	log.Debug("New browser profile created", "profile_builder", builder)
	return latest
}

func Random(optionalFilter ...*PoolFilter) *BrowserProfile {
	if len(optionalFilter) == 0 {
		return random_with_filter(nil)
	}
	return random_with_filter(optionalFilter[0])
}

func random_with_filter(poolFilter *PoolFilter) *BrowserProfile {
	builder := browserProfileBuilder{}

	var hasFilter bool = poolFilter != nil

	if hasFilter && len(poolFilter.allowedFamilySlice) != 0 {
		builder.Family = gp_slices.RandomFrom(poolFilter.allowedFamilySlice)
	} else {
		builder.Family = family.Random()
	}

	if hasFilter && len(poolFilter.allowedPlatformSlice) != 0 {
		builder.Platform = gp_slices.RandomFrom(poolFilter.allowedPlatformSlice)
	} else {
		builder.Platform = platform.Random()
	}

	if hasFilter && len(poolFilter.allowedPlatformSlice) != 0 {
		builder.Platform = gp_slices.RandomFrom(poolFilter.allowedPlatformSlice)
	} else {
		builder.Platform = platform.Random()
	}

	return New(builder)
}

func Latest() *BrowserProfile {
	if latest == nil {
		return Random()
	}
	return latest
}
