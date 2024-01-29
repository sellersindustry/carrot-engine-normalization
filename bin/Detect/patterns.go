package Detect

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


var PATTERNS = []*Pattern {
	{
		DetectRegexp: Utility.CompileRegex(`/^(\{\{)[\w\=\-\s]+(\}\})/`),
		Class:        Token.Control,
	}, {
		DetectRegexp: Utility.CompileRegex(`/^[\n\s\t]+/`),
		Class:        Token.Space,
	}, {
		DetectRegexp: Utility.CompileRegex(`/^[0-9]{1,4}[\.\/-][0-9]{1,2}[\.\/-]\d{1,4}\b/`),
		Class:        Token.Date,
	}, {
		DetectRegexp: Utility.CompileRegex(`/^[0-9]{1,2}:[0-9]{2}\b/`),
		Class:        Token.Time,
	}, {
		DetectRegexp: Utility.CompileRegex(`/^([\+][\s]?[0-9]{1,3}[\s]?)?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4}\b/`),
		Class:        Token.Phone,
	}, {
		DetectRegexp: Utility.CompileRegex(`/^[0-9]*(\,[0-9]{3})*(\.[0-9]+)?/`),
		Class:        Token.Number,
	}, {
		// Hashtag
		DetectRegexp: Utility.CompileRegex(`/^#[a-zA-Z0-9-_]*[a-zA-Z0-9]/`),
		Class:        Token.Special,
	}, {
		// Email
		DetectRegexp: Utility.CompileRegex(`/^[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,64}/`),
		Class:        Token.Special,
	}, {
		// URL
		DetectRegexp: Utility.CompileRegex(`/^([A-Za-z0-9]+:\/\/)?([A-Za-z0-9]+(\.[A-Za-z0-9]+)+)(([\/\?])([\S]*[0-9A-Za-z\/])?)?/`),
		Class:        Token.Special,
	}, {
		// Units
		DetectWords: Utility.GetWordsetKeys("./bin/wordsets/units.txt"),
		Class:       Token.Word,
	}, {
		// Word
		DetectRegexp: Utility.CompileRegex(`/^[a-zA-Z]+(\-[a-zA-Z]+)*(\'[a-zA-Z]+)?/`),
		Class:        Token.Word,
	},
}
