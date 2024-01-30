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


func NumberYear(text string) string {
	year       := parseInt(text);
	secondHalf := year % 100;
	firstHalf  := (year - secondHalf) / 100;
	if firstHalf == 20 && secondHalf >= 0 && secondHalf <= 9 {
		return Number(text);
	}
	if firstHalf > 10 && secondHalf > 0 && secondHalf <= 9 {
		return Number(strconv.Itoa(firstHalf)) + " o " + Number(strconv.Itoa(secondHalf));
	}
	if firstHalf > 0 && secondHalf == 0 {
		return Number(strconv.Itoa(firstHalf)) + " hundred";
	}
	if firstHalf == 0 {
		return Number(strconv.Itoa(secondHalf));
	}
	return Number(strconv.Itoa(firstHalf)) + " " + Number(strconv.Itoa(secondHalf));

	// if len(text) != 4 {
	// 	return Number(text);
	// }
	// firstHalf  := parseInt(text[0:2]);
	// secondHalf := parseInt(text[2:4]);
	// if firstHalf == 20 && secondHalf >= 0 && secondHalf <= 9 {
	// 	return Number(text);
	// }
	// if firstHalf > 10 && secondHalf >= 0 && secondHalf <= 9 {
	// 	return Number(strconv.Itoa(firstHalf)) + " o " + Number(strconv.Itoa(secondHalf));
	// }
	// if firstHalf > 0 && secondHalf == 0 {
	// 	return Number(strconv.Itoa(firstHalf)) + " hundred";
	// }
	// return Number(strconv.Itoa(firstHalf)) + " " + Number(strconv.Itoa(secondHalf));
}


func NumberCurrency(text string, currency string) string {
	if currency == "$" {
		if strings.Contains(text, ".") {
			dollars := strings.Split(text, ".")[0]
			cents   := strings.Split(text, ".")[1]
			return Number(dollars) + " dollars and " + Number(cents) + " cents";
		} else {
			return Number(text) + " dollars";
		}
	} else if currency == "£" {
		if strings.Contains(text, ".") {
			pounds := strings.Split(text, ".")[0]
			pence  := strings.Split(text, ".")[1]
			return Number(pounds) + " pounds and " + Number(pence) + " pence";
		} else {
			return Number(text) + " pounds";
		}
	}
	return Number(text);
}


func Number(text string) string {
	parts := strings.Split(strings.ReplaceAll(text, ",", ""), ".")
	full  := num2words.ConvertAnd(parseInt(parts[0]));
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



