package Wordify

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


var NON_INITIALISMS = []string{ "com", "net", "org", "co" }


func Special(text string) string {
	buffer := []string{}
	for _, chunk := range splitSpecial(text) {
		if unicode.IsNumber(rune(chunk[0])) {
			buffer = append(buffer, Number(chunk))
		} else if unicode.IsLetter(rune(chunk[0])) {
			if Utility.Contains(NON_INITIALISMS, strings.ToLower(chunk)) {
				buffer = append(buffer, chunk)
			} else if Utility.IsInitialism(chunk) {
				buffer = append(buffer, strings.Join(strings.Split(chunk, ""), " "))
			} else {
				buffer = append(buffer, Utility.Decamelize(chunk))
			}
		} else {
			keys, values := Utility.GetWordset("./bin/wordsets/symbols-general.txt")
			buffer = append(buffer, Utility.KeywordMapping(chunk, keys, values, false))
		}
	}
	return strings.Join(buffer, " ")
}


func splitSpecial(s string) []string {
	re := regexp.MustCompile(`([A-Za-z]+)|([0-9]+)|([^A-Za-z0-9])`)
	return re.FindAllString(s, -1)
}

