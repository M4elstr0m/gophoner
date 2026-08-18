package language

import (
	"math/rand"
	"strings"
)

//go:generate stringer -type=Language
type Language int

const (
	EN Language = iota
	ES
	FR
	DE
	ZH
	HI
	PT
	RU
	VI
	JA
	KO
	AR
	TL
	IT
	PL
	FA
	UR
	TH
	HE
)

var LENGTH Language = Language(len(_Language_index) - 1) // did that to prevent stringer from using it

func Random() Language {
	return Language(rand.Intn(int(LENGTH)))
}

func (self Language) acceptLanguage() string {
	switch self {
	case EN:
		return "en-US,en;q=0.9"
	case ES:
		return "es;q=0.8,es-419;q=0.7"
	case FR:
		return "fr;q=0.8,fr-CA;q=0.7"
	case ZH:
		return "zh-CN;q=0.8,zh;q=0.7"
	case PT:
		return "pt-BR;q=0.8,pt;q=0.7"
	default:
		return strings.ToLower(self.String()) + ";q=0.8"
	}
}

func (self Language) AcceptLanguageENPrimary() string {
	if self == EN {
		return EN.acceptLanguage()
	}
	return EN.acceptLanguage() + "," + self.acceptLanguage()
}
