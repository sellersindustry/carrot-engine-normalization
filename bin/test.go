package Normalize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codycollier/wer"
)


func ExecuteTests() {
	wers     := []float64{};
	pass     := 0;
	fail     := 0;
	failText := [][]string{};
	for _, entry := range getEntries() {
		wer, output, target := processEntry(entry);
		if wer == 0.0 {
			pass += 1;
		} else {
			fail += 1;
			failText = append(failText, []string{output, target})
		}
		wers = append(wers, wer);
	}

	for _, entry := range failText {
		fmt.Println("OUTPUT: ", entry[0]);
		fmt.Println("TARGET: ", entry[1]);
		fmt.Println("");
	}

	fmt.Println("---------")
	fmt.Println("PASS: ", pass)
	fmt.Println("FAIL: ", fail)
	fmt.Println("WER:  ", avgFloats(wers))
}


func processEntry(entry []string) (float64, string, string) {
	input  := entry[0];
	output := strings.Join(Process(input, true).Sentences, " ");
	target := entry[1];
	wer, _ := wer.WER(strings.Split(target, " "), strings.Split(output, " "))
	return wer, output, target
}


func getEntries() [][]string {
	entries := [][]string{}
	for _, filename := range getFiles("./tests") {
		inputs := strings.Split(openTextFile(filename), "------INPUT------");
		for _, input := range inputs {
			sections := strings.Split(input, "------OUTPUT------")
			if len(sections) != 2 {
				continue
			}
			sections[0] = strings.Trim(sections[1], "\n\r")
			sections[1] = strings.Trim(sections[1], "\n\r")
			entries = append(entries, []string{sections[0], sections[1]})
		}
	}
	return entries;
}


func getFiles(dir string) []string {
	filenames := []string{};
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading test directory \"", dir, "\":", err)
		return nil;
	}
	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, filepath.Join(dir, file.Name()))
		}
	}
	return filenames
}


func openTextFile(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file \"", filepath, "\":", err)
		return ""
	}
	return string(content)
}


func avgFloats(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

