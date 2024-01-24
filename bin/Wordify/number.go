package Wordify

import (
	"strconv"
	"strings"

	"github.com/divan/num2words"
)



func NumberOrdinal(text string) string {
	//! "6" -> "sixth", "8" -> "eighth", "1" -> "first"
	//! parse number
	//! based on number st, th, nd ify the number ...
	word := Number(text);
	return word;
}


func NumberPlural(text string) string {
	//! "8s" -> "eights", "80s" -> "eighties", "1970s" -> "1970s"
	//! if its a 4 digit year just assume it's a year and parse it using year
	//! pluralify the word given
	word := Number(text);
	return word;
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
	//! take in number, parse number, return words
	//! just add the before the number and call the other one
	return "";
}


func RomanNumeral(text string) string {
	//! take in number, parse number, return words
	return "";
}



