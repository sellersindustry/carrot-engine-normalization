package Utility

import (
	"fmt"
	"os"
	"regexp"
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

