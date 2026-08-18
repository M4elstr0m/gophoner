package browser_profiles

import (
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/family"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/language"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/platform"
	_slices "github.com/M4elstr0m/gophoner/internal/utils/slices"
	"github.com/charmbracelet/log"
)

type PoolFilter struct {
	AllowedFamilySlice   []family.Family
	AllowedPlatformSlice []platform.Platform
	AllowedLanguageSlice []language.Language
}

func New(builder browserProfileBuilder) *BrowserProfile {
	log.Info("New browser profile created", "profile_builder", builder)
	return builder.build()
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

	if hasFilter && len(poolFilter.AllowedFamilySlice) != 0 {
		builder.Family = _slices.RandomFrom(poolFilter.AllowedFamilySlice)
	} else {
		builder.Family = family.Random()
	}

	allowedPlatforms := builder.Family.AllowedPlatforms()
	var hasPlatformRestriction bool = len(allowedPlatforms) != 0 // 'allowedPlatforms != nil &&' just discovered it was useless :P power of go!
	if hasFilter && len(poolFilter.AllowedPlatformSlice) != 0 {  // if there is hardcoded filter
		if hasPlatformRestriction { // if there is also platform restriction
			allowedPlatforms = _slices.Inter(allowedPlatforms, poolFilter.AllowedPlatformSlice)
			if len(allowedPlatforms) == 0 {
				log.Fatal("No platform compatible with both the family-allowed platforms and the platform pool filter",
					"family", builder.Family,
					"platform-pool-filter", poolFilter.AllowedPlatformSlice,
				)
			}
			builder.Platform = _slices.RandomFrom(allowedPlatforms)
		} else {
			builder.Platform = _slices.RandomFrom(poolFilter.AllowedPlatformSlice)
		}
	} else if hasPlatformRestriction { // platform restriction only
		builder.Platform = _slices.RandomFrom(allowedPlatforms)
	} else { // fully random
		builder.Platform = platform.Random()
	}

	if hasFilter && len(poolFilter.AllowedLanguageSlice) != 0 {
		builder.Language = _slices.RandomFrom(poolFilter.AllowedLanguageSlice)
	} else {
		builder.Language = language.Random()
	}

	return New(builder)
}
