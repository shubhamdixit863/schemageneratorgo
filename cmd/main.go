package main

import (
	"github.com/shubh/json/pkg/converter"
)

const dir = "./schemas" // absolute paths
const outputDir = "./output"

func main() {
	cs := new(converter.ConvertSchema)
	cs.Convert(dir, outputDir)

}
