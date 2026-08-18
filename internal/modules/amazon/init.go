package amazon

import "github.com/M4elstr0m/gophoner/internal/modules"

var _ modules.IsRegisteredFunc = IsRegistered

func init() {
	modules.Register(modules.Amazon, IsRegistered)
	modules.RegisterColor(modules.Amazon, "#FF6201")
}
