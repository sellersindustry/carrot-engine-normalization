package Wordify

import (
	"regexp"
	"strings"
)

func Phone(text string) string {
	if strings.Contains(text, "+") || strings.Contains(text, "(") {
		text = text[1:]
		text = strings.ReplaceAll(text, " ", "-")
	}
	str := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(text, "-")
	sections := strings.Split(str, "-")
	var normalizedPhoneNumber []string
	for i := range sections {
		normalizedPhoneNumber = append(normalizedPhoneNumber, NumberNominal(sections[i]))
		if i < len(sections)-1 {
			normalizedPhoneNumber = append(normalizedPhoneNumber, ", ")
		}
	}
	return strings.TrimSpace(strings.Join(normalizedPhoneNumber, ""))
}
