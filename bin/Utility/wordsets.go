package Utility

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func GetWordsetKeys(filename string) []string {
	key, _ := GetWordset(filename)
	return key
}


func GetWordsetValues(filename string) []string {
	_, value := GetWordset(filename)
	return value
}


func GetWordsetBoth(filename string) []string {
	key, value := GetWordset(filename)
	return append(key, value...)
}


func GetWordset(filename string) ([]string, []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: Failed to reading file \"", filename, "\"");
		os.Exit(1);
	}

	var key   []string;
	var value []string;

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t\t")
		key = append(key, parts[0])
		value = append(value, parts[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: Failed to reading file \"", filename, "\"");
		os.Exit(1);
	}
	return key, value;
}

