package platform

import (
	"fmt"
	"math/rand"
	"strings"

	_slices "github.com/M4elstr0m/gophoner/internal/utils/slices"
	"github.com/charmbracelet/log"
)

//go:generate stringer -type=Platform
type Platform int

const (
	// Desktop

	Windows Platform = iota
	MacOS
	Linux
	ChromeOS

	// Mobile

	Android
	IOS
)

const lastDesktopPlatform = ChromeOS

var LENGTH Platform = Platform(len(_Platform_index) - 1)

var DesktopPlatforms = platformRange(0, lastDesktopPlatform)
var MobilePlatforms = platformRange(lastDesktopPlatform+1, LENGTH-1)

func platformRange(from, to Platform) []Platform {
	out := make([]Platform, 0, int(to-from)+1)
	for p := from; p <= to; p++ {
		out = append(out, p)
	}
	return out
}

func Random() Platform {
	return Platform(rand.Intn(int(LENGTH)))
}

func (self Platform) UserAgentToken() string {
	switch self {
	case Windows:
		return "Windows NT 10.0; Win64; x64"
	case MacOS:
		return "Macintosh; Intel Mac OS X 10_15_7"
	case Linux:
		return "X11; Linux x86_64"
	case ChromeOS:
		return _slices.RandomFrom([]string{
			"X11; CrOS x86_64 16700.46.0",
			"X11; CrOS x86_64 15633.69.0",
			"X11; CrOS x86_64 15359.58.0",
			"X11; CrOS x86_64 15772.44.0",
		})
	case Android:
		return _slices.RandomFrom([]string{
			"Linux; Android 17; Pixel 10",
			"Linux; Android 17; Pixel 10 Pro",
			"Linux; Android 17; Pixel 9",
			"Linux; Android 16; SM-S948B",
			"Linux; Android 16; SM-S938B",
			"Linux; Android 15; SM-S928B",
		})
	case IOS:
		log.Fatal("IOS user agent token was not randomized properly")
		return ""
	default:
		log.Fatal("Bad value", "platform", self)
		return ""
	}
}

func IOSDeviceUserAgentToken(iosVersion string) string {
	return fmt.Sprintf("iPhone; CPU iPhone OS %s like Mac OS X", strings.ReplaceAll(iosVersion, ".", "_"))
}

func (self Platform) SecCHUAPlatform() string {
	switch self {
	case Windows:
		return "Windows"
	case MacOS:
		return "macOS"
	case Linux:
		return "Linux"
	case ChromeOS:
		return "Chrome OS"
	case Android:
		return "Android"
	case IOS:
		return "iOS"
	default:
		log.Fatal("Bad value", "platform", self)
		return ""
	}
}

func (self Platform) IsMobile() bool {
	return self > lastDesktopPlatform
}

func (self Platform) SecCHUAMobile() string {
	if self.IsMobile() {
		return "?1"
	} else {
		return "?0"
	}
}
