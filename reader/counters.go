package reader

import (
	"errors"
	"strconv"
	"strings"
)

type counters struct {
	videos    int
	endpoints int
	requests  int
	caches    int
	capacity  int
}

func readCounters(line string) (c counters, err error) {
	numData := strings.Split(line, " ")

	// validate
	if len(numData) != 5 {
		err = errors.New("invalid data received in counters parser")
		return
	}

	// extract counters
	c.videos, err = strconv.Atoi(numData[0])
	if err != nil {
		return
	}

	c.endpoints, err = strconv.Atoi(numData[1])
	if err != nil {
		return
	}

	c.requests, err = strconv.Atoi(numData[2])
	if err != nil {
		return
	}

	c.caches, err = strconv.Atoi(numData[3])
	if err != nil {
		return
	}

	c.capacity, err = strconv.Atoi(numData[4])
	if err != nil {
		return
	}

	return
}
