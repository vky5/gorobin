package balancer

import "sync"

type RoundRobin struct {
	mu      sync.Mutex // to prevent race condition when index is accessed by multiple goroutine
	servers []string
	index   int
}

func NewRoundRobin(servers []string) *RoundRobin { // returning a pointer to it because we want it to be globally same and not a copy
	return &RoundRobin{servers: servers}
}



// receiver function that will take the input as the pointer of the RoundRobin type
func (rr *RoundRobin) Next() string {
	rr.mu.Lock()
	defer rr.mu.Unlock() // because if a panic occur in between it might never unlock that's why we use defer and not like what I did below
	server := rr.servers[rr.index]
	rr.index = (rr.index + 1) % len(rr.servers) // updating the index of the server to the next
	// rr.mu.Unlock() // unlocking mutex after update done but defer is much safer
	return server // returning the server to send the address to
}


// did this but it is better to use receiver function 
// func UpdateIndex(rr *RoundRobin) string {
// 	rr.mu.Lock()
// 	defer rr.mu.Unlock()

// 	server := rr.servers[rr.index]
// 	rr.index = (rr.index + 1) % len(rr.servers)

// 	return server
// }
