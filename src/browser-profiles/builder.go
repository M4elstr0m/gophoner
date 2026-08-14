package browser_profiles

import (
	"fmt"
	"math/rand"

	"github.com/M4elstr0m/gophoner/src/browser-profiles/enums/family"
	"github.com/M4elstr0m/gophoner/src/browser-profiles/enums/language"
	"github.com/M4elstr0m/gophoner/src/browser-profiles/enums/platform"
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

	tlsProfileSlice := builder.Family.TlsProfileSlice()

	out.TlsProfile = tlsProfileSlice[rand.Intn(len(tlsProfileSlice))]
	tlsProfileHelloIdVersion := out.TlsProfile.GetClientHelloId().Version

	out.UserAgent = fmt.Sprintf(
		"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s Safari/537.36",
		builder.Platform.UserAgentToken(),
		builder.Family.UserAgentToken(tlsProfileHelloIdVersion),
	)

	out.SecCHUA = builder.Family.SecCHUA(tlsProfileHelloIdVersion)
	out.SecCHUAMobile = builder.Platform.SecCHUAMobile()
	out.SecCHUAPlatform = builder.Platform.SecCHUAPlatform()

	out.AcceptLanguage = builder.Language.AcceptLanguageENPrimary()

	return out
}
