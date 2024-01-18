package detect

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
	"github.com/sellersindustry/normalization-tts/bin/utility"
)


func number(buffer string) *token.Model {
	return utility.DetectReturnFirstNonNullToken(buffer,
		numberRomanNumeral,
		numberCurrency,
		numberPlural,
		numberYear,
		numberOrdinal,
		numberGenernalOrNominal,
	);
}


// Roman Numerals
func numberRomanNumeral(buffer string) *token.Model {
	regexp := `^([MDCLXVI]*[MDCLXV]+[MDCLXVI]*)\b`;
	return utility.DetectTokenSubclass(buffer, regexp, token.Number, token.NumberRoman);
}


// Currency Numbers
func numberCurrency(buffer string) *token.Model {
	regexp := `^((\$)|(€)|(ƒ)|(¥)|(JP¥)|(CN¥))(\s)*[0-9]+(,[0-9]+)*(\.[0-9]+)?`;
	return utility.DetectTokenSubclass(buffer, regexp, token.Number, token.NumberCurrency);
}


// Plural Numbers
func numberPlural(buffer string) *token.Model {
	regexp := `^(([0-9]){1,4}(s|S))\b`;
	return utility.DetectTokenSubclass(buffer, regexp, token.Number, token.NumberPlural);
}


// Year Numbers
func numberYear(buffer string) *token.Model {
	regexp := `^((2[0-1])|(1[0-9]))[0-9]{2}\b`;
	return utility.DetectTokenSubclass(buffer, regexp, token.Number, token.NumberYear);
}


// Ordinal Number - 1st, 1,324th
func numberOrdinal(buffer string) *token.Model {
	regexp := `^(([0-9])+(\,[0-9]{3})*(st|ST|nd|ND|rd|RD|th|TH))\b`;
	return utility.DetectTokenSubclass(buffer, regexp, token.Number, token.NumberOrdinal);
}


// Nominal Numbers - each digit seperated
func numberGenernalOrNominal(buffer string) *token.Model {
	rGeneral, _  := regexp.Compile(`^[0-9]*(\,[0-9]{3})*(\.[0-9]+)*`);
	match        := rGeneral.FindStringIndex(buffer);
	if (len(match) > 0) {
		buffer = buffer[match[0]:match[1]];
		if (len(buffer) == 0) {
			return nil;
		}
		rNominal, _  := regexp.Compile(`^[0-5]{5,}$`);
		isNominal    := rNominal.FindStringIndex(buffer);
		if (len(isNominal) > 0) {
			return token.NewSubclass(buffer, token.Number, token.NumberNominal);
		}
		return token.NewGeneral(buffer, token.Number);
	}
	return nil;
}

