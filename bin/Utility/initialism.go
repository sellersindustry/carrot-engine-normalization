package Utility

import (
	"os"
	"strings"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
)


var WORDS = getWords();
var ACRONYM_OVERRIDE    = []string{ "DARPA", "OPEC", "NATO", "UNESCO", "UNICEF" };
var INITIALISM_OVERRIDE = []string{ "AD", "US", "AI", "IT" };


func IsInitialism(text string) bool {
	if strings.HasSuffix(text, "'s") {
		text = strings.TrimSuffix(text, "'s")
	}
	if strings.HasSuffix(text, "s") {
		text = strings.TrimSuffix(text, "s")
	}
	if Contains(INITIALISM_OVERRIDE, strings.ToUpper(text)) {
		return true;
	}
	if Contains(ACRONYM_OVERRIDE, strings.ToUpper(text)) {
		return false;
	}
	return !wordExists(lemmatize(text));
}


func wordExists(word string) bool {
	return binarySearch(WORDS, word);
}


func binarySearch(words []string, target string) bool {
	low, high := 0, len(words)-1
	for low <= high {
		mid := (low + high) / 2
		if words[mid] == target {
			return true
		} else if words[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}


func lemmatize(text string) string {
	lemmatizer, err := golem.New(en.New())
	if err != nil {
		return text
	}
	return lemmatizer.Lemma(strings.ToLower(text))
}


func getWords() []string {
	content, _ := os.ReadFile("./bin/wordsets/words.txt")
	words      := strings.Split(string(content), "\n")
	for i := 0; i < len(words); i++ {
		words[i] = strings.TrimSpace(words[i])
	}
	return words;
}

