package Wordify

import "strings"

func Time(text string) string {
	sections := strings.Split(text, ":")
	if len(sections) != 2 {
		return "";
	}
	return Number(sections[0]) + " " + Number(sections[1])
}
