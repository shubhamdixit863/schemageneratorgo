package main

import (
	"github.com/shubh/json/pkg/converter"
)

const dir = "./schemas" // absolute paths
const outputDir = "./output"
const outputPackageName = "output"

func main() {
	converter.Convert(dir, outputDir, outputPackageName)

}
