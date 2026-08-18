package logs

import "github.com/charmbracelet/log"

func OmitIfNotDebug(key string, value any) []any {
	if log.GetLevel() != log.DebugLevel {
		return nil
	}
	return []any{key, value}
}
