package Wordify

import (
	"strconv"
	"strings"
)

func Time(text string) string {
	sections := strings.Split(text, ":")
	num := parseInt(sections[0]) % 12
	if num == 0 {
		num = 12
	}
	sections[0] = strconv.Itoa(num)
	if len(sections) != 2 {
		return "";
	}
	if sections[1] == "00" {
		return Number(sections[0]) + " o'clock"
	}
	return Number(sections[0]) + " " + Number(sections[1])
}
