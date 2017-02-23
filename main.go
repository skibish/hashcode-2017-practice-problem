package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/skibish/hashcode-2017-practice-problem/reader"
)

var (
	inputFile = flag.String("input-file", "_data/me_at_the_zoo.in", "Input File")
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

	// parse incoming data
	videos, endpoints, parseErr := fileReader.Parse()
	if parseErr != nil {
		log.Fatal("failed to read input data: " + parseErr.Error())
	}

	fmt.Printf("videos: %+v\n", videos)
	fmt.Printf("endpoints: %+v\n", endpoints)
}
