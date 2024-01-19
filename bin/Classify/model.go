package Classify

import "github.com/sellersindustry/normalization-tts/bin/Token"


type PatternScan struct {
	Keywords []string
	Range    int
}


type Pattern struct {
	Current       []string
	Prefix        []string
	Suffix        []string
	ScanBefore    PatternScan
	ScanAfter     PatternScan
	SetSubclassTo Token.Subclass
	SetIsSilentTo bool
}


var CLASS_NUMBER = "{{CLASS:NUMBER}}";

