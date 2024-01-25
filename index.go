package main

import (
	"fmt"

	"github.com/sellersindustry/normalization-tts/bin/Classify"
	"github.com/sellersindustry/normalization-tts/bin/Detect"
	"github.com/sellersindustry/normalization-tts/bin/Transform"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)

// var text = "Hello world! How are you're you today? That's pretty nice! sellers' and evan's but not nice'";
// var text = "1980s and 1990s. 15th, 2nd, 1ST"
var text = "foo nice 180 - 5 nice foo"


func main() {
	fmt.Println(Utility.RegexWordListOr([]string{"in."}));
	gtokens     := Detect.Process(text);
	Classify.Process(gtokens);
	fmt.Println(gtokens);
	Transform.Process(gtokens);
	fmt.Println(gtokens);
}
