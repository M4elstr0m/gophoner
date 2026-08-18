package family

import (
	"fmt"
	"math/rand"

	"github.com/M4elstr0m/gophoner/internal/browser_profiles/enums/platform"
	_strings "github.com/M4elstr0m/gophoner/internal/utils/strings"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/charmbracelet/log"
)

//go:generate stringer -type=Family
type Family int

const (
	Chrome Family = iota
	Brave
	Firefox
	Safari
)

var LENGTH Family = Family(len(_Family_index) - 1)

func Random() Family {
	return Family(rand.Intn(int(LENGTH)))
}

func (self Family) TlsProfileSlice(p platform.Platform) []*profiles.ClientProfile {
	switch self {
	case Chrome:
		return []*profiles.ClientProfile{
			&profiles.Chrome_107,
			&profiles.Chrome_120,
			&profiles.Chrome_124,
			&profiles.Chrome_133,
			&profiles.Chrome_144,
			&profiles.Chrome_146,
		}
	case Brave:
		return []*profiles.ClientProfile{
			&profiles.Brave_146,
		}
	case Firefox:
		return []*profiles.ClientProfile{
			&profiles.Firefox_108,
			&profiles.Firefox_117,
			&profiles.Firefox_120,
			&profiles.Firefox_123,
			&profiles.Firefox_135,
			&profiles.Firefox_147,
			&profiles.Firefox_148,
		}
	case Safari:
		if p == platform.IOS {
			return []*profiles.ClientProfile{
				&profiles.Safari_IOS_16_0,
				&profiles.Safari_IOS_17_0,
				&profiles.Safari_IOS_18_0,
				&profiles.Safari_IOS_26_0,
			}
		}
		return []*profiles.ClientProfile{
			&profiles.Safari_15_6_1,
			&profiles.Safari_16_0,
		}
	default:
		log.Fatal("Bad value", "family", self)
		return nil
	}
}

func (self Family) UserAgent(p platform.Platform, tlsProfileHelloIdVersion string) string {
	switch self {
	case Chrome, Brave:
		mobileToken := ""
		if p.IsMobile() {
			mobileToken = "Mobile "
		}
		return fmt.Sprintf(
			"Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s.0.0.0 %sSafari/537.36",
			p.UserAgentToken(),
			tlsProfileHelloIdVersion,
			mobileToken,
		)
	case Firefox:
		if p.IsMobile() {
			androidVersion := _strings.UserAgentTokenToAndroidVersion(p.UserAgentToken())
			return fmt.Sprintf("Mozilla/5.0 (Android %[1]s; Mobile; rv:%[2]s.0) Gecko/%[2]s.0 Firefox/%[2]s.0",
				androidVersion,
				tlsProfileHelloIdVersion,
			)
		}
		return fmt.Sprintf("Mozilla/5.0 (%s; rv:%[2]s.0) Gecko/20100101 Firefox/%[2]s.0",
			p.UserAgentToken(),
			tlsProfileHelloIdVersion,
		)
	case Safari:
		if p == platform.IOS {
			return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%s Mobile/15E148 Safari/604.1",
				platform.IOSDeviceUserAgentToken(tlsProfileHelloIdVersion),
				tlsProfileHelloIdVersion,
			)
		}
		return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/%s Safari/605.1.15",
			p.UserAgentToken(),
			tlsProfileHelloIdVersion,
		)
	default:
		log.Fatal("Bad value", "family", self)
		return ""
	}
}

func (self Family) SecCHUA(tlsProfileHelloIdVersion string) string {
	switch self {
	case Chrome, Brave:
		return fmt.Sprintf("\"Chromium\";v=\"%[1]s\", \"Not=A?Brand\";v=\"24\", \"Google Chrome\";v=\"%[1]s\"",
			tlsProfileHelloIdVersion,
		)
	case Firefox, Safari:
		return ""
	default:
		log.Fatal("Bad value", "family", self)
		return ""
	}
}

func (self Family) AllowedPlatforms() []platform.Platform {
	switch self {
	case Safari:
		return []platform.Platform{
			platform.MacOS,
			platform.IOS,
		}
	default:
		return []platform.Platform{
			platform.Windows,
			platform.Linux,
			platform.MacOS,
			platform.ChromeOS,
			platform.Android,
		}
	}
}
