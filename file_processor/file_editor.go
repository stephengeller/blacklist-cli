package FileProcessor

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func AddLinesToFile(file string, sites []string) error {
	err := checkFileExists(file)
	if err != nil {
		return err
	}

	input, err := ioutil.ReadFile(file)
	lines := strings.Split(string(input), "\n")
	linesWithSites := append(lines, sites...)

	output := strings.Join(linesWithSites, "\n")
	err = ioutil.WriteFile(file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func RemoveLinesFromFile(filePath string, linesToRemove []string) error {
	fileContents, err := ReadFile(filePath)
	var newFileContents []string

	if err != nil {
		return err
	}

	fmt.Printf("\nRemoving lines %v\n", linesToRemove)

	match := false

	for _, fileLine := range fileContents {
		match = false
		for _, lineToRemove := range linesToRemove {
			if !isWhitespace(fileLine) && !isWhitespace(lineToRemove) && strings.Contains(fileLine, lineToRemove) {
				match = true
				continue
			}
		}
		if match == false {
			newFileContents = append(newFileContents, fileLine)
		}
	}

	err = ReplaceFileContents(filePath, newFileContents)

	if err != nil {
		return err
	}

	return nil
}

func isWhitespace(str string) bool {
	return strings.TrimSpace(str) == ""
}

func ReplaceFileContents(filePath string, newContents []string) error {
	output := strings.Join(newContents, "\n")
	err := ioutil.WriteFile(filePath, []byte(output), 0644)

	if err != nil {
		return err
	}

	return nil
}
