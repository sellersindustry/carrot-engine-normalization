package utility

import (
	"regexp"

	"github.com/sellersindustry/normalization-tts/bin/token"
)


func DetectTokenGeneral(buffer string, regex string, class token.Class) *token.Model {
	r, _    := regexp.Compile(regex);
	match   := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.NewGeneral(buffer[match[0]:match[1]], class);
	}
	return nil;
}


func DetectTokenSubclass(buffer string, regex string, class token.Class, subclass token.Subclass) *token.Model {
	r, _    := regexp.Compile(regex);
	match   := r.FindStringIndex(buffer);
	if (len(match) > 0) {
		return token.NewSubclass(buffer[match[0]:match[1]], class, subclass);
	}
	return nil;
}


func DetectReturnFirstNonNullToken(buffer string, functions ...func(buffer string) *token.Model) *token.Model {
	for _, fn := range functions {
		token := fn(buffer)
		if (token != nil) {
			return token
		}
	}
	return nil
}
