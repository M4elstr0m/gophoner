package recon

import (
	"sync"

	"github.com/M4elstr0m/gophoner/internal/logs"
	"github.com/M4elstr0m/gophoner/internal/modules"
	_slices "github.com/M4elstr0m/gophoner/internal/utils/slices"
	"github.com/charmbracelet/log"
)

type Input struct {
	PhoneNumber string
	ModuleSlice modules.ModuleSlice
}

type Output struct {
	Indicator resultIndicator
	Module    modules.Module
	Error     error
}

type resultIndicator int

const (
	Registered resultIndicator = iota
	NotRegistered
	LimitReached
	ModuleError
)

func Run(input *Input) []Output {
	log.Info("New recon initiated",
		_slices.Merge(
			[]any{
				"module-list", input.ModuleSlice,
			},
			logs.OmitIfNotDebug("phonenumber", input.PhoneNumber),
		)...,
	)

	var out []Output = make([]Output, len(input.ModuleSlice))

	var wg sync.WaitGroup
	wg.Add(len(input.ModuleSlice))

	for i, module := range input.ModuleSlice {
		go func(i int, module modules.Module) {
			defer wg.Done()
			out[i] = checkModule(input.PhoneNumber, module)
		}(i, module)
	}

	wg.Wait()

	return out
}
