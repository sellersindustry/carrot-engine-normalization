package main

import (
	Normalize "github.com/sellersindustry/normalization-tts/bin"
)


// var text = "foo nice 180 - 5 nice foo.\nThat's pretty nice!! {{HELLO}}"

func main() {
	Normalize.ExecuteTests();
	// output := Normalize.Process(text, false);
	// for _, _token := range output.Tokens {
	// 	fmt.Println(_token);
	// }
	// fmt.Println("\n");
	// for _, sentence := range output.Sentences {
	// 	fmt.Println("\"" + sentence + "\"");
	// }
	// fmt.Println(Normalize.ProcessJSON(text, false))
}
