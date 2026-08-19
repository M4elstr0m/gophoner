package facebook

import "github.com/M4elstr0m/gophoner/internal/modules"

var _ modules.IsRegisteredFunc = IsRegistered

func init() {
	modules.Register(modules.Facebook, IsRegistered)
	modules.RegisterColor(modules.Facebook, "#0866FF")
}
