package endpoint

// Endpoint contains endpoint data.
type Endpoint struct {
	DatacenterLatency int         // latency in milliseconds
	CacheLatency      map[int]int //map[CacheID]latency
}
