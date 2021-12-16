package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

func FileNameCreation(fileName string) string {
	return fmt.Sprintf("%s %s", fileName[:len(fileName)-len(filepath.Ext(fileName))], ".go")
}

func SuffixFileExtension(fileName string) string {

	return strings.ToUpper(strings.TrimSuffix(fileName, filepath.Ext(fileName)))
}

// Removes extra special characters from string and the

func Sanitizestring(fileName string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(strings.TrimSuffix(fileName, filepath.Ext(fileName)), "")

}
