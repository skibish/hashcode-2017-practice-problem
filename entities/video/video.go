package video

// Video contains video data.
type Video struct {
	Size     int         // size in megabytes
	Requests map[int]int // map[endpointID]numberOfRequests
}
