package Classify

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/Token"
)


type PatternScan struct {
	Exists       []string
	NotExists    []string
	Range        int
	IgnoreSpaces bool
}


type Pattern struct {
	CurrentByRegexp     *regexp.Regexp
	CurrentByClass		Token.Class
	CurrentBySubclass	Token.Subclass
	CurrentByWords      []string

	HasPrefix           []string
	HasSuffix           []string

	ScanBefore          *PatternScan
	ScanAfter           *PatternScan
	
	SetSubclassTo       Token.Subclass
	SetIsInactiveTo     bool
}

