package recon

import (
	"errors"
	"fmt"

	"github.com/M4elstr0m/gophoner/internal/logs"
	"github.com/M4elstr0m/gophoner/internal/modules"
	_slices "github.com/M4elstr0m/gophoner/internal/utils/slices"
	"github.com/charmbracelet/log"
)

func checkModule(phoneNumber string, module modules.Module) Output {
	checkFunc, ok := modules.Get(module)
	if !ok {
		log.Error("Failed to find a module in registry",
			"module", module,
		)
		return Output{
			Indicator: ModuleError,
			Module:    module,
			Error:     fmt.Errorf("failed to find a module in registry: %v", module),
		}
	}

	log.Info("Checking phone number registration status",
		_slices.Merge(
			logs.OmitIfNotDebug("phonenumber", phoneNumber),
			[]any{
				"module", module,
			},
		)...,
	)

	var indicator resultIndicator = NotRegistered
	isRegistered, err := checkFunc(phoneNumber)
	if err != nil {
		log.Error("Failed to check a phone number",
			_slices.Merge(
				logs.OmitIfNotDebug("phonenumber", phoneNumber),
				[]any{
					"module", module,
					"error", err,
				},
			)...,
		)

		if errors.Is(err, modules.ErrorLimitReached) {
			indicator = LimitReached
		} else {
			indicator = ModuleError
		}

	} else {
		log.Info("Successfully checked a phone number",
			_slices.Merge(
				logs.OmitIfNotDebug("phonenumber", phoneNumber),
				[]any{
					"module", module,
				},
			)...,
		)

		if isRegistered {
			indicator = Registered
		}
	}

	return Output{
		Indicator: indicator,
		Module:    module,
		Error:     err,
	}
}
