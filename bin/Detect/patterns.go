package Detect

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


var PATTERNS = []*Pattern {
	{
		// Control
		DetectRegexp: Utility.CompileRegex(`/^(\{\{)[\w\=\-\s]+(\}\})/`),
		Class:        Token.Control,
	}, {
		// Space
		DetectRegexp: Utility.CompileRegex(`/^[\n\s\t]+/`),
		Class:        Token.Space,
	}, {
		// Date
		DetectRegexp: Utility.CompileRegex(`/^[0-9]{1,4}[\.\/-][0-9]{1,2}[\.\/-]\d{1,4}\b/`),
		Class:        Token.Date,
	}, {
		// Time
		DetectRegexp: Utility.CompileRegex(`/^[0-9]{1,2}:[0-9]{2}\b/`),
		Class:        Token.Time,
	}, {
		// Phone
		DetectRegexp: Utility.CompileRegex(`/^([\+][\s]?[0-9]{1,3}[\s]?)?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4}\b/`),
		Class:        Token.Phone,
	}, {
		// Number
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
		// Initialism
		DetectRegexp: Utility.CompileRegex(`/^([A-Z]\.)+([A-Z])(\.)/`),
		Class:        Token.Word,
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
		DetectRegexp: Utility.CompileRegex(`/^[A-Za-zÀ-ÖØ-öø-ÿ]+(\-[A-Za-zÀ-ÖØ-öø-ÿ]+)*(\'[A-Za-zÀ-ÖØ-öø-ÿ]+)?/`),
		Class:        Token.Word,
	},
}
