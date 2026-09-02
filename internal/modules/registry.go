package modules

import (
	"github.com/charmbracelet/log"
)

type IsRegisteredFunc func(phoneNumber string) IsRegisteredFuncOutput

type IsRegisteredFuncOutput struct {
	IsRegistered          bool
	AdditionalInformation map[string]string
	Err                   error
}

func NewIsRegisteredFuncOutput(isRegistered bool, additionalInformation map[string]string, err error) IsRegisteredFuncOutput {
	return IsRegisteredFuncOutput{
		isRegistered,
		additionalInformation,
		err,
	}
}

func MinimalIsRegisteredFuncOutput(isRegistered bool, err error) IsRegisteredFuncOutput {
	return IsRegisteredFuncOutput{
		isRegistered,
		nil,
		err,
	}
}

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
