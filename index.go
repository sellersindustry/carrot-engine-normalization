package main

import (
	"fmt"

	"github.com/sellersindustry/normalization-tts/bin"
)

var text = "Hello world! \"Test Test Test\""

func main() {
	// Normalize.ExecuteTests();
	output := Normalize.Process(text, false);
	// for _, _token := range output.Tokens {
	// 	fmt.Println(_token.);
	// }
	// fmt.Println("\n");
	fmt.Println(output.Tokens);
	for _, sentence := range output.Sentences {
		fmt.Println(sentence);
	}
	// fmt.Println(strings.Join(output.Sentences, " "));
	// fmt.Println(Normalize.ProcessJSON(text, false))
}
