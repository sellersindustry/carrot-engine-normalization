package main

import (
	"fmt"

	"github.com/sellersindustry/normalization-tts/bin/detect"
)


func main() {
    fmt.Print(detect.Process("Hello world! How are you're you today? That's pretty nice!"));
}
