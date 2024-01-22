package main

import (
	"fmt"

	"github.com/sellersindustry/normalization-tts/bin/Classify"
	"github.com/sellersindustry/normalization-tts/bin/Detect"
)

// var text = "Hello world! How are you're you today? That's pretty nice! sellers' and evan's but not nice'";
// var text = "1980s and 1990s. 15th, 2nd, 1ST"
var text = "+ 12.23 - 30"

func main() {
	gtokens     := Detect.Process(text);
    fmt.Println(gtokens);
	fmt.Println("");
	Classify.Process(gtokens);
	fmt.Println(gtokens);
}
