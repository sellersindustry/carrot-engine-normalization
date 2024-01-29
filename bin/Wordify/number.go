package Wordify

import (
	"strconv"
	"strings"

	"github.com/brandenc40/romannumeral"
	"github.com/divan/num2words"
	"github.com/gertd/go-pluralize"
)


var ordinalNumberSuffixes = []string{
	"first", "second", "third", "fourth", "fifth", "sixth", "seventh",
	"eighth", "ninth", "tenth", "eleventh", "twelfth", "thirteenth", "fourteenth",
	"fifteenth", "sixteenth", "seventeenth", "eighteenth", "nineteenth",
	"twentieth",
}


func NumberOrdinal(text string) string {
	integer     := parseInt(text);
	lastDigit   := integer % 10;
	last2Digits := integer % 100;
	if lastDigit != 0 {
		if integer <= 20 {
			return ordinalNumberSuffixes[integer - 1];
		} else if last2Digits <= 20 {
			return Number(strconv.Itoa(integer - last2Digits)) + " and " + ordinalNumberSuffixes[last2Digits - 1];
		} else {
			return Number(strconv.Itoa(integer - lastDigit)) + " " + ordinalNumberSuffixes[lastDigit - 1];
		}
	}
	word := Number(strconv.Itoa(integer))
	if strings.HasSuffix(word, "y") {
		return strings.TrimSuffix(word, "y") + "ieth";
	}
	return word + "th";
}


func NumberPlural(text string) string {
	var numWord string;
	if len(text) == 4 {
		numWord = NumberYear(text);
	} else {
		numWord = Number(text);
	}
	return pluralize.NewClient().Plural(numWord);
}


// Fail - 257
// Fail - 278
func NumberYear(text string) string {
	if len(text) != 4 {
		return Number(text);
	}
	firstHalf  := parseInt(text[0:2]);
	secondHalf := parseInt(text[2:4]);
	if firstHalf == 20 && secondHalf >= 0 && secondHalf <= 9 {
		return Number(text);
	}
	if firstHalf > 10 && secondHalf >= 0 && secondHalf <= 9 {
		return Number(strconv.Itoa(firstHalf)) + " o " + Number(strconv.Itoa(secondHalf));
	}
	if firstHalf > 0 && secondHalf == 0 {
		return Number(strconv.Itoa(firstHalf)) + " hundred";
	}
	return Number(strconv.Itoa(firstHalf)) + " " + Number(strconv.Itoa(secondHalf));
}


func NumberCurrency(text string, currency string) string {
	//! parse number add the unit behind it dollars
	//! ensure cents are taken care of
	//! what are cents in other currencys? What currencies do we need to support??
	//! $ and £
	return "";
}


func Number(text string) string {
	parts := strings.Split(strings.ReplaceAll(text, ",", ""), ".")
	full  := strings.ReplaceAll(num2words.ConvertAnd(parseInt(parts[0])), "-", " ");
	if len(parts) > 1 {
		return full + " point " + NumberNominal(parts[1]);
	} else {
		return full;
	}
}


func NumberNominal(text string) string {
	numbers := []string{};
	for _, char := range text {
		numbers = append(numbers, num2words.Convert(parseInt(string(char))));
	}
	return strings.Join(numbers, " ");
}


func parseInt(text string) int {
	number, err := strconv.Atoi(text)
	if err == nil {
		return number;
	} else {
		return 0;
	}
}


func RomanNumeralPossessive(text string) string {
	integer, err := romannumeral.StringToInt(text)
	if err != nil {
		return "zero";
	}
	return "the " + NumberOrdinal(strconv.Itoa(integer));
}


func RomanNumeral(text string) string {
	integer, err := romannumeral.StringToInt(text)
	if err != nil {
		return "zero";
	}
	return Number(strconv.Itoa(integer));
}



