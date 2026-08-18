package modules

import (
	"github.com/charmbracelet/log"
)

type IsRegisteredFunc func(phoneNumber string) (bool, error)

var registry = map[Module]IsRegisteredFunc{}

func Register(m Module, fn IsRegisteredFunc) {
	_, exists := registry[m]
	if exists {
		log.Fatal("Attempt to register module multiple times",
			"module", m,
		)
	}

	registry[m] = fn
}

func Get(m Module) (IsRegisteredFunc, bool) {
	fn, ok := registry[m]
	return fn, ok
}
