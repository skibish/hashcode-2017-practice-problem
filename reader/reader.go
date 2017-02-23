package reader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/skibish/hashcode-2017-practice-problem/entities/video"
)

// Reader contains reader data.
type Reader struct {
	filePath string
	data     string
}

// New returns new Reader.
func New(filePath string) (*Reader, error) {
	r := new(Reader)

	// validate file
	stat, fileErr := os.Stat(filePath)
	switch {
	case os.IsNotExist(fileErr):
		// doesn't exist
		return nil, fileErr
	case stat.IsDir():
		// is directory
		return nil, errors.New(filePath + " is a directory")
	default:
		// valid file
		data, dataErr := ioutil.ReadFile(filePath)
		if dataErr != nil {
			return nil, dataErr
		}

		r.data = string(data)
	}

	r.filePath = filePath

	return r, nil
}

// Parse parses incoming data.
func (r *Reader) Parse() (err error) {
	// define counters
	var countersData counters

	// define videos
	var videos []video.Video

	// loop through data
	for i, line := range strings.Split(r.data, "\n") {
		switch i {
		case 0:
			// the first line of the input contains the numbers
			countersData, err = readCounters(line)
			if err != nil {
				return
			}

		default:
			return

		}

		fmt.Println(line)
	}

	fmt.Println(countersData)

	return
}
