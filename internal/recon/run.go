package recon

import (
	"encoding/json"
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
	Indicator             resultIndicator   `json:"indicator"`
	AdditionalInformation map[string]string `json:"additional_information,omitempty"`
	Module                modules.Module    `json:"module"`
	Error                 error             `json:"-"`
}

func (o Output) MarshalJSON() ([]byte, error) {
	type alias Output

	var errMessage string
	if o.Error != nil {
		errMessage = o.Error.Error()
	}

	return json.Marshal(struct {
		alias
		Error string `json:"error,omitempty"`
	}{alias(o), errMessage})
}

type resultIndicator int

const (
	Registered resultIndicator = iota
	NotRegistered
	LimitReached
	ModuleError
)

func (i resultIndicator) MarshalJSON() ([]byte, error) {
	switch i {
	case Registered:
		return json.Marshal("REGISTERED")
	case NotRegistered:
		return json.Marshal("NOT REGISTERED")
	case LimitReached:
		return json.Marshal("RATE LIMIT")
	case ModuleError:
		return json.Marshal("ERROR")
	default:
		return json.Marshal("UNKNOWN")
	}
}

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

func Collect(events <-chan Event) []Output {
	var out []Output

	for event := range events {
		if event.Kind == ModuleFinished {
			out = append(out, event.Output)
		}
	}

	return out
}
