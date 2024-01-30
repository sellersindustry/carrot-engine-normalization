package main

import (
	"fmt"
	"strings"

	"github.com/sellersindustry/normalization-tts/bin"
)


var text = "Zombie cookie. Some cookies are automatically recreated after a user has deleted them; these are called zombie cookies."
var isTest = true;


func main() {
	if isTest {
		Normalize.ExecuteTests();
	} else {
		output := Normalize.Process(text, false);
		for _, _token := range output.Tokens {
			fmt.Println(_token);
		}

		fmt.Println("\n");
		fmt.Println(text);
		fmt.Println(strings.Join(output.Sentences, " "));
	}
}


// PASS 565
// PASS 579
// PASS 580