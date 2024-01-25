package Normalize

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sellersindustry/normalization-tts/bin/Classify"
	"github.com/sellersindustry/normalization-tts/bin/Detect"
	"github.com/sellersindustry/normalization-tts/bin/Token"
	"github.com/sellersindustry/normalization-tts/bin/Transform"
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
		if len(sentences[len(sentences)-1]) == 0 && (*tokens)[index].Class == Token.Space {
			continue
		}
		sentences[len(sentences)-1] += (*tokens)[index].Text
		if isSentenceEnd(&(*tokens)[index]) {
			sentences = append(sentences, "")
		}

	}
	return sentences;
}


func isSentenceEnd(token *Token.Model) bool {
	if token.Subclass != Token.Punctuation {
		return false;
	}
	return token.Text == "." || token.Text == "!" || token.Text == "?"
}

