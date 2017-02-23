package reader

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/skibish/hashcode-2017-practice-problem/entities/endpoint"
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
func (r *Reader) Parse() (videos []video.Video, endpoints []endpoint.Endpoint, err error) {
	var currentStage int

	var endpointsStage int
	var endpointsTmpData []int
	var endpointsLatencyData map[int]int // map[cacheID]latency

	// define counters
	var countersData counters

	// loop through data
	for _, line := range strings.Split(r.data, "\n") {
		switch currentStage {
		case 0:
			// the first line of the input contains the numbers
			countersData, err = readCounters(line)
			if err != nil {
				return
			}

			currentStage++

		case 1:
			// read
			for _, videoSize := range strings.Split(line, " ") {
				// get size
				var sizeInt int
				sizeInt, err = strconv.Atoi(videoSize)
				if err != nil {
					return
				}

				videos = append(videos, video.Video{
					Size: sizeInt,
				})
			}

			currentStage++

		case 2:
			// read endpoint data
			switch endpointsStage {
			case 0:
				// read first line of endpoints data
				endpointsTmpData, err = stringSliceToInt(strings.Split(line, " "))
				if err != nil {
					return
				}

				if len(endpointsTmpData) != 2 {
					err = errors.New("invalid endpoints data: " + line)
					return
				}

				endpointsLatencyData = make(map[int]int)

				// increase stage
				endpointsStage++

			case 1:
				// read endpoint latency data
				var endpointsLatencyTmpData []int
				endpointsLatencyTmpData, err = stringSliceToInt(strings.Split(line, " "))
				if err != nil {
					return
				}

				if len(endpointsLatencyTmpData) != 2 {
					err = errors.New("invalid endpoints latency data")
					return
				}

				// fill the map
				endpointsLatencyData[endpointsLatencyTmpData[0]] = endpointsLatencyTmpData[1]

				endpointsTmpData[1]--

				if endpointsTmpData[1] == 0 {
					endpointsStage = 0

					// append endpoint
					endpoints = append(endpoints, endpoint.Endpoint{
						CacheLatency: endpointsLatencyData,
					})

				}

				if len(endpoints) == countersData.endpoints {
					currentStage++
				}

			default:
				err = errors.New("invalid endpointsStage")
				return
			}

		case 3:
			if line == "" {
				return
			}

			// read requests data
			var requestsTmpData []int
			requestsTmpData, err = stringSliceToInt(strings.Split(line, " "))
			if err != nil {
				return
			}

			if len(requestsTmpData) != 3 {
				err = errors.New("invalid requests data: " + line)
				return
			}

			videoID := requestsTmpData[0]
			endpointID := requestsTmpData[1]
			numRequests := requestsTmpData[2]

			// write requests to endpoints
			if videos[videoID].Requests == nil {
				videos[videoID].Requests = make(map[int]int)
			}

			videos[videoID].Requests[endpointID] = numRequests

		default:
			// printMarshaled("videos", videos)
			// printMarshaled("endpoints", endpoints)

			err = errors.New("invalid currentStage")
			return

		}

	}

	return
}
