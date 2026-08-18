package browser_profiles

import (
	"math/rand"

	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/family"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/language"
	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/platform"
	"github.com/bogdanfinn/tls-client/profiles"
)

type browserProfileBuilder struct {
	Family   family.Family
	Platform platform.Platform
	Language language.Language
}

type BrowserProfile struct {
	TlsProfile      *profiles.ClientProfile
	UserAgent       string
	SecCHUA         string
	SecCHUAMobile   string
	SecCHUAPlatform string
	AcceptLanguage  string
}

func (builder browserProfileBuilder) build() *BrowserProfile {
	var out *BrowserProfile = &BrowserProfile{}

	tlsProfileSlice := builder.Family.TlsProfileSlice(builder.Platform)

	out.TlsProfile = tlsProfileSlice[rand.Intn(len(tlsProfileSlice))]
	tlsProfileHelloIdVersion := out.TlsProfile.GetClientHelloId().Version

	out.UserAgent = builder.Family.UserAgent(
		builder.Platform,
		tlsProfileHelloIdVersion,
	)

	out.SecCHUA = builder.Family.SecCHUA(tlsProfileHelloIdVersion)
	out.SecCHUAMobile = builder.Platform.SecCHUAMobile()
	out.SecCHUAPlatform = builder.Platform.SecCHUAPlatform()

	out.AcceptLanguage = builder.Language.AcceptLanguageENPrimary()

	return out
}
