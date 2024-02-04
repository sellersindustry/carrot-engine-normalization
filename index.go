package main

import (
	"fmt"
	"strings"

	Normalize "github.com/sellersindustry/normalization-tts/bin"
)

var text = ""
var isTest = true

func main() {
	if isTest {
		Normalize.ExecuteTests()
	} else {
		output := Normalize.Process(text, false)
		for _, _token := range output.Tokens {
			fmt.Println(_token)
		}

		fmt.Println("\n")
		fmt.Println(text)
		fmt.Println(strings.Join(output.Sentences, " "))
	}
}

// FAIL 116
