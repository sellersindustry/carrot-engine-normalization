package Classify

import "github.com/sellersindustry/normalization-tts/bin/Token"


type PatternScan struct {
	Exists       []string
	NotExists    []string
	Range        int
	IgnoreSpaces bool
}


type Pattern struct {
	Current         string
	Prefix          []string
	Suffix          []string
	ScanBefore      *PatternScan
	ScanAfter       *PatternScan
	SetSubclassTo   Token.Subclass
	SetIsInactiveTo bool
}

