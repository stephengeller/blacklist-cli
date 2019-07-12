package FileProcessor

import (
	"io/ioutil"
	"os"
	"strings"
)

type ReaderErr string

const ErrNoFilePassed = ReaderErr("No file passed")

func (e ReaderErr) Error() string {
	return string(e)
}

func checkFileExists(fileLocation string) error {
	if strings.TrimSpace(fileLocation) == "" {
		return ErrNoFilePassed
	}

	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		errString := "File " + fileLocation + " not found"
		return ReaderErr(errString)
	}
	return nil
}

func ReadFile(fileLocation string) ([]string, error) {
	err := checkFileExists(fileLocation)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadFile(fileLocation)

	if err != nil {
		return nil, err
	}

	strContents := strings.Split(string(contents), "\n")
	if strContents[0] == "" {
		strContents = strContents[1:]
	}

	return strContents, nil
}
