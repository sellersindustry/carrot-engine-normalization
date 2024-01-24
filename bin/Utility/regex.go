package Utility

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)


func CompileRegex(regex string) *regexp.Regexp {
	r, err := regexp.Compile(regex)
	if err != nil {
		fmt.Printf("Error: Failed to compile regular expression: \"%s\"\n", regex)
		os.Exit(1)
		return nil
	}
	return r
}


func RegexWordListOr(keywords []string) string {
	for i := range keywords {
		keywords[i] = strings.Replace(keywords[i], "/", "\\/", -1);
		keywords[i] = strings.Replace(keywords[i], ".", "\\.", -1);
	}
	return "((" + strings.Join(keywords, ")|(") + "))";
}

