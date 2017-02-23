package reader

import (
	"errors"
	"io/ioutil"
	"os"
)

// Reader contains reader data.
type Reader struct {
	filePath string
	data     []byte
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
		var dataErr error
		r.data, dataErr = ioutil.ReadFile(filePath)
		if dataErr != nil {
			return nil, dataErr
		}
	}

	r.filePath = filePath

	return r, nil
}
