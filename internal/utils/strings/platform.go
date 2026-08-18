package _strings

import "strings"

func UserAgentTokenToAndroidVersion(token string) string {
	_, after, found := strings.Cut(token, "Android ")
	if !found {
		return ""
	}
	version, _, _ := strings.Cut(after, ";")
	return version
}
