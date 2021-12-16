package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/shubh/json/pkg/inputs"
	"github.com/shubh/json/pkg/utils"
)

const dir = "./schemas"
const outputDir = "./output"
const outputPackageName = "output"

func main() {
	var inputFiles []string

	// Reading file recursively from the directory
	files, _ := os.ReadDir(dir)

	path, _ := filepath.Abs(dir)

	for _, file := range files {
		inputFiles = append(inputFiles, filepath.Join(path, file.Name()))

		schemas, err := inputs.ReadInputFiles(inputFiles, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}

		g := inputs.New(schemas...)

		err = g.CreateTypes()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failure generating structs: ", err)
			os.Exit(1)
		}

		var w io.Writer = os.Stdout

		w, err = os.Create(filepath.Join(outputDir, filepath.Base(utils.FileNameCreation(file.Name()))))

		inputs.Output(w, g, outputPackageName, utils.SuffixFileExtension(file.Name()))

	}

}
