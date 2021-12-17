package converter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/shubh/json/pkg/inputs"
	"github.com/shubh/json/pkg/utils"
)

type ConvertSchema struct {
	files            []os.DirEntry
	inputFiles       []string
	packageDirectory string
	packageName      string
}

//Reads the files From Directories
func (cs *ConvertSchema) readFilesFromDir(dir string) {
	cs.files, _ = os.ReadDir(dir)
}

// Package name formatting
func (cs *ConvertSchema) packageFormat(outputDir string, file os.DirEntry) {
	cs.packageDirectory = fmt.Sprintf("%s/%s", outputDir, utils.Sanitizestring(file.Name()))
	cs.packageName = utils.Sanitizestring(file.Name())
}

func (cs *ConvertSchema) Convert(dir string, outputDir string) {
	path, _ := filepath.Abs(dir) //Get the absolute path of the files
	cs.readFilesFromDir(dir)

	for _, file := range cs.files {
		cs.inputFiles = append(cs.inputFiles, filepath.Join(path, file.Name()))
		schemas, err := inputs.ReadInputFiles(cs.inputFiles, true)
		if err != nil {
			os.Exit(1)
		}
		g := inputs.New(schemas...)
		err = g.CreateTypes(utils.Sanitizestring(utils.SuffixFileExtension(file.Name()))) // Passing file name if in case the title is not prsent in the document
		if err != nil {
			os.Exit(1)
		}
		var w io.Writer = os.Stdout
		cs.packageFormat(outputDir, file)
		err = os.Mkdir(cs.packageDirectory, 0755)
		w, err = os.Create(filepath.Join(cs.packageDirectory, filepath.Base(utils.FileNameCreation(file.Name()))))
		inputs.Output(w, g, cs.packageName)

	}

}
