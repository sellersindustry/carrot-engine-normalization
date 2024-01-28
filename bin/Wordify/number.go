package Wordify

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brandenc40/romannumeral"
	"github.com/divan/num2words"
	"github.com/gertd/go-pluralize"
)


var ordinalNumberSuffixes = []string{
	"first", "second", "third", "fourth", "fifth", "sixth", "seventh",
	"eighth", "ninth", "eleventh", "twelfth", "thirteenth", "fourteenth",
	"fifteenth", "sixteenth", "seventeenth", "eighteenth", "nineteenth",
	"twentieth",
}


func NumberOrdinal(text string) string {
	integer   := parseInt(text);
	lastDigit := integer % 10;
	fmt.Println(integer, lastDigit);
	if lastDigit != 0 {
		if integer <= 20 {
			return ordinalNumberSuffixes[integer - 1];
		}
		return Number(strconv.Itoa(integer - lastDigit)) + " " + ordinalNumberSuffixes[lastDigit - 1];
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


func NumberYear(text string) string {
	//! "2001" -> "two thousand and one", "2023" -> "twenty twenty-three"
	//! this below might help
	//! https://learningenglish.voanews.com/a/pronouncing-years-in-american-english/7045997.html
	return "";
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
	return Number(num2words.ConvertAnd(integer));
}



