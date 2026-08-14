package family

import (
	"fmt"
	"math/rand"

	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/charmbracelet/log"
)

//go:generate stringer -type=Family
type Family int

const (
	Chrome Family = iota
)

var LENGTH Family = Family(len(_Family_index) - 1)

func Random() Family {
	return Family(rand.Intn(int(LENGTH)))
}

func All() []Family {
	out := make([]Family, 0, LENGTH)
	for i := range LENGTH {
		out = append(out, i)
	}
	return out
}

func (self Family) TlsProfileSlice() []*profiles.ClientProfile {
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
	default:
		log.Fatal("Bad value", "family", self)
		return nil
	}
}

func (self Family) UserAgentToken(tlsProfileHelloIdVersion string) string {
	switch self {
	case Chrome:
		return "Chrome/" + tlsProfileHelloIdVersion + ".0.0.0"
	default:
		log.Fatal("Bad value", "family", self)
		return ""
	}
}

// tlsProfileHelloIdVersion := tlsProfile.GetClientHelloId().Version
func (self Family) SecCHUA(tlsProfileHelloIdVersion string) string {
	switch self {
	case Chrome:
		return fmt.Sprintf("\"Chromium\";v=\"%s\", \"Not=A?Brand\";v=\"24\", \"Google Chrome\";v=\"%s\"",
			tlsProfileHelloIdVersion,
			tlsProfileHelloIdVersion,
		)
	default:
		log.Fatal("Bad value", "family", self)
		return ""
	}
}
