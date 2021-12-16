package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func FileNameCreation(fileName string) string {
	return fmt.Sprintf("%s %s", fileName[:len(fileName)-len(filepath.Ext(fileName))], ".go")
}

func SuffixFileExtension(fileName string) string {

	return strings.ToUpper(strings.TrimSuffix(fileName, filepath.Ext(fileName)))
}
