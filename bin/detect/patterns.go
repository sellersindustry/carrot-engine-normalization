package detect

import "github.com/sellersindustry/normalization-tts/bin/token"


var patterns = []*Pattern {
	{
		Regexp: `^(\{\{)[\w\=\-\s]+(\}\})`,
		Class:  token.Control,
	}, {
		Regexp: `^[\n\s\t]+`,
		Class:  token.Space,
	}, {
		Regexp: `^#[a-zA-Z0-9-_]*[a-zA-Z0-9]`,
		Class:  token.Hashtag,
	}, {
		Regexp: `^[A-Z0-9a-z._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,64}`,
		Class:  token.Email,
	}, {
		Regexp: `^([A-Za-z0-9]+:\/\/)?([A-Za-z0-9]+(\.[A-Za-z0-9]+)+)(([\/\?])([\S]*[0-9A-Za-z\/])?)?`,
		Class:  token.URL,
	}, {
		Regexp: `^([\+][\s]?[0-9]{1,3}[\s]?)?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4}\b`,
		Class:  token.Phone,
	}, {
		Regexp: `^[0-9]{1,4}[\.\/-][0-9]{1,2}[\.\/-]\d{1,4}\b`,
		Class:  token.Date,
	}, {
		Regexp: `^[0-9]{1,2}:[0-9]{2}\b`,
		Class:  token.Time,
	}, {
		Regexp: `^((\$)|(€)|(ƒ)|(¥)|(JP¥)|(CN¥))?[0-9]*(\,[0-9]{3})*(\.[0-9]+)*`,
		Class:  token.NumberCurrency,
	}, {
		Regexp: `^[0-9]*(\,[0-9]{3})*(\.[0-9]+)*`,
		Class:  token.Number,
	}, {
		Regexp: `^[a-zA-Z]+(\-[a-zA-Z]+)*(\'[a-zA-Z]+)?`,
		Class:  token.Word,
	},
}
