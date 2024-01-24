package Detect

import (
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


var PATTERNS = []*Pattern {
	{
		Regexp: `^(\{\{)[\w\=\-\s]+(\}\})`,
		Class:  Token.Control,
	}, {
		Regexp: `^[\n\s\t]+`,
		Class:  Token.Space,
	}, {
		Regexp: `^[0-9]{1,4}[\.\/-][0-9]{1,2}[\.\/-]\d{1,4}\b`,
		Class:  Token.Date,
	}, {
		Regexp: `^[0-9]{1,2}:[0-9]{2}\b`,
		Class:  Token.Time,
	}, {
		Regexp: `^([\+][\s]?[0-9]{1,3}[\s]?)?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4}\b`,
		Class:  Token.Phone,
	}, {
		Regexp: `^[0-9]*(\,[0-9]{3})*(\.[0-9]+)?`,
		Class:  Token.Number,
	}, {
		Regexp: `^#[a-zA-Z0-9-_]*[a-zA-Z0-9]`,
		Class:  Token.Hashtag,
	}, {
		Regexp: `^[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,64}`,
		Class:  Token.Email,
	}, {
		Regexp: `^([A-Za-z0-9]+:\/\/)?([A-Za-z0-9]+(\.[A-Za-z0-9]+)+)(([\/\?])([\S]*[0-9A-Za-z\/])?)?`,
		Class:  Token.URL,
	}, {
		Regexp: `^([MDCLXVI]*[MDCLXV]+[MDCLXVI]*)\b`,
		Class:  Token.RomanNumeral,
	}, {
		Regexp: `^` + Utility.RegexWordListOr(Utility.GetWordsetKeys("./bin/wordsets/units.txt")) + `\b`,
		Class:  Token.Word,
	}, {
		Regexp: `^[a-zA-Z]+(\-[a-zA-Z]+)*(\'[a-zA-Z]+)?`,
		Class:  Token.Word,
	}, {
		Regexp: `^[,.!?]`,
		Class:  Token.Punctuation,
	},
}
