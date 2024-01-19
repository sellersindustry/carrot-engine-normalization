package Classify

import "github.com/sellersindustry/normalization-tts/bin/Token"


var PATTERNS = []*Pattern {
	{ 
		// Plural Numbers - Before (Number)
		Current:       []string{ CLASS_NUMBER },
		Suffix:        []string{ "s" },
		SetSubclassTo: Token.NumberPlural,
	}, {
		// Plural Numbers - After (s)
		Current:       []string{ "s" },
		Prefix:        []string{ CLASS_NUMBER },
		SetIsSilentTo: true,
	}, {
		// Ordinal Numbers - Before (Number)
		Current:       []string{ CLASS_NUMBER },
		Suffix:        []string{ "/^(st|nd|rd|th)$/" },
		SetSubclassTo: Token.NumberOrdinal,
	}, {
		// Ordinal Numbers - After (st, nd, rd, th)
		Current:       []string{ "/^(st|nd|rd|th)$/" },
		Prefix:        []string{ CLASS_NUMBER },
		SetIsSilentTo: true,
	},
}

