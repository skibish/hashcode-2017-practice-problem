package main

import (
	"flag"
	"log"

	"github.com/skibish/hashcode-2017-practice-problem/reader"
)

var (
	inputFile = flag.String("input-file", "file.in", "Input File")
)

var (
	fileReader *reader.Reader
)

func main() {
	// initialize file reader
	var fileReaderErr error
	fileReader, fileReaderErr = reader.New(*inputFile)
	if fileReaderErr != nil {
		log.Fatal("failed to read input file: " + fileReaderErr.Error())
	}

}
