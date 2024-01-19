package main

import (
	"fmt"

	"github.com/sellersindustry/normalization-tts/bin/detect"
)


// var text = "Hello world! How are you're you today? That's pretty nice! sellers' and evan's but not nice'";
var text = "This is a great world sellersew@gmail.com 695.5435, $134"

func main() {
    fmt.Println(detect.Process(text));
}
