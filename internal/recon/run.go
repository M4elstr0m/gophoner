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
	Indicator             resultIndicator
	AdditionalInformation map[string]string
	Module                modules.Module
	Error                 error
}

type resultIndicator int

const (
	Registered resultIndicator = iota
	NotRegistered
	LimitReached
	ModuleError
)

type EventKind int

const (
	ModuleStarted EventKind = iota
	ModuleFinished
)

type Event struct {
	Kind   EventKind
	Module modules.Module
	Output Output
}

func Run(input *Input) <-chan Event {
	log.Info("New recon initiated",
		_slices.Merge(
			[]any{
				"module-list", input.ModuleSlice,
			},
			logs.OmitIfNotDebug("phonenumber", input.PhoneNumber),
		)...,
	)

	out := make(chan Event)

	go func() {
		defer close(out)

		var wg sync.WaitGroup
		wg.Add(len(input.ModuleSlice))

		for _, module := range input.ModuleSlice {
			go func(module modules.Module) {
				defer wg.Done()
				out <- Event{Kind: ModuleStarted, Module: module}
				out <- Event{Kind: ModuleFinished, Module: module, Output: checkModule(input.PhoneNumber, module)}
			}(module)
		}

		wg.Wait()
	}()

	return out
}
