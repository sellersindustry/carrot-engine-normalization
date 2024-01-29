package Detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/Token"
)


type Pattern struct {
	DetectRegexp   *regexp.Regexp
	DetectWords    []string
	Class          Token.Class
}

