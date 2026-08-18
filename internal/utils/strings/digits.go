package _strings

import (
	"errors"
	"strings"
	"unicode"
)

func DigitsOnly(s string) error {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return errors.New("digits only")
		}
	}
	return nil
}

func DigitsWithOptionalPlus(s string) error {
	return DigitsOnly(strings.TrimPrefix(s, "+"))
}
