package Lib

import (
	"strings"
	"time"
)

func FillVarsToTemplate(content string, vars map[string]string) string {
	for key, item := range vars {
		content = strings.ReplaceAll(content, "{{"+key+"}}", item)
	}
	return content
}

func DateFormat(date time.Time) string {
	return date.Format("02 Jan 2006")
}

func Between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}
