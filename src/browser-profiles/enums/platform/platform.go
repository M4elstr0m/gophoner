package platform

import (
	"math/rand"

	"github.com/charmbracelet/log"
)

//go:generate stringer -type=Platform
type Platform int

const (
	Windows Platform = iota
	MacOS
	Linux
	ChromeOS
)

var LENGTH Platform = Platform(len(_Platform_index) - 1)

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
		return "X11; CrOS x86_64 15633.69.0"
	default:
		log.Fatal("Bad value", "platform", self)
		return ""
	}
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
	default:
		log.Fatal("Bad value", "platform", self)
		return ""
	}
}

func (self Platform) IsMobile() bool {
	switch self {
	default:
		return false
	}
}

func (self Platform) SecCHUAMobile() string {
	if self.IsMobile() {
		return "?1"
	} else {
		return "?0"
	}
}
