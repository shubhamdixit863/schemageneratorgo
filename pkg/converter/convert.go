package converter

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/shubh/json/pkg/inputs"
	"github.com/shubh/json/pkg/utils"
)

func readFilesFromDir(dir string) []fs.DirEntry {
	files, _ := os.ReadDir(dir)
	return files

}

func Convert(dir string, outputDir string, outputPackageName string) {

	path, _ := filepath.Abs(dir)

	for _, file := range readFilesFromDir(dir) {
		var inputFiles []string
		inputFiles = append(inputFiles, filepath.Join(path, file.Name()))

		schemas, err := inputs.ReadInputFiles(inputFiles, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}

		g := inputs.New(schemas...)

		err = g.CreateTypes(utils.SuffixFileExtension(file.Name())) // Passing file name if in case the title is not prsent in the document
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failure generating structs: ", err)
			os.Exit(1)
		}

		var w io.Writer = os.Stdout

		w, err = os.Create(filepath.Join(outputDir, filepath.Base(utils.FileNameCreation(file.Name()))))

		inputs.Output(w, g, outputPackageName, utils.SuffixFileExtension(file.Name()))

	}

}
