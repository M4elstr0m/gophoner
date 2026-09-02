package openai

import "github.com/M4elstr0m/gophoner/internal/modules"

var _ modules.IsRegisteredFunc = IsRegistered

func init() {
	modules.Register(modules.OpenAI, IsRegistered)
	modules.RegisterColor(modules.OpenAI, "#74aa9c")
}
