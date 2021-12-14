package utils

import (
	"fmt"
	"path/filepath"
)

func FileNameCreation(fileName string) string {
	return fmt.Sprintf("%s %s", fileName[:len(fileName)-len(filepath.Ext(fileName))], ".go")
}
