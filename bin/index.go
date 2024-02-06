package Normalize

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sellersindustry/normalization-tts/bin/Classify"
	"github.com/sellersindustry/normalization-tts/bin/Detect"
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Transform"
	"github.com/sellersindustry/normalization-tts/bin/Utility"
)


type Output struct {
	Sentences []string
	Groups    [][]string
	Tokens    []Token.Model
}


func Process(buffer string, hasControls bool) Output {
	gtokens := Detect.Process(buffer)
	Classify.Process(gtokens)
	Transform.Process(gtokens)
	return Output{
		Sentences: sentenize(gtokens, hasControls),
		Groups:    [][]string{}, //! Add Groupings
		Tokens:    *gtokens,
	};
}


func ProcessJSON(buffer string, hasControls bool) string {
	jsonData, err := json.Marshal(Process(buffer, hasControls))
	if err != nil {
		fmt.Println("Error:", err);
		os.Exit(1);
		return "{}";
	}
	return string(jsonData);
}


func sentenize(tokens *[]Token.Model, hasControls bool) []string {
	sentences := []string{""};
	for index := range *tokens {
		if !hasControls && (*tokens)[index].Class == Token.Control {
			continue
		}
		if (*tokens)[index].IsInactive {
			continue
		}
		if len(sentences[len(sentences)-1]) != 0 && (*tokens)[index].Subclass != Token.Punctuation {
			sentences[len(sentences)-1] += " "
		}
		if isSentenceEnd(&(*tokens)[index]) {
			sentences[len(sentences)-1] = trimAllSpace(sentences[len(sentences)-1])
			sentences[len(sentences)-1] += (*tokens)[index].Text
			sentences = append(sentences, "")
		} else {
			sentences[len(sentences)-1] += (*tokens)[index].Text
		}
	}
	sentences[len(sentences)-1] = trimAllSpace(sentences[len(sentences)-1])
	return Utility.FilterNonEmptyStrings(sentences);
}


func isSentenceEnd(token *Token.Model) bool {
	if token.Subclass != Token.Punctuation {
		return false;
	}
	return token.Text == "." || token.Text == "!" || token.Text == "?"
}


func trimAllSpace(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

