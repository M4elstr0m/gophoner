package google

import "github.com/M4elstr0m/gophoner/internal/modules"

var _ modules.IsRegisteredFunc = IsRegistered

func init() {
	modules.Register(modules.Google, IsRegistered)
	modules.RegisterColor(modules.Google, "#4285F4")
}
