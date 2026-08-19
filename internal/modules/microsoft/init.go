package microsoft

import "github.com/M4elstr0m/gophoner/internal/modules"

var _ modules.IsRegisteredFunc = IsRegistered

func init() {
	modules.Register(modules.Microsoft, IsRegistered)
	modules.RegisterColor(modules.Microsoft, "#89d2ff")
}
