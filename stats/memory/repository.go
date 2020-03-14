package memory

import (
	"reflect"
	"sync"

	"github.com/l-lin/fizzbuzz/stats"
)

// Repository stores stats in memory
// /!\ Do not use this implementation if you want to use the web server in cluster
// because this resource is not shared among all the web servers (each web server
// will have their own Repository).
// If you want to have it in cluster, consider using other approaches by implementing
// the "stats.Repository" interface:
// - log the requests -> use logstash to centralize -> use elasticsearch for metrics
// - use 3rd party locking system (e.g. zookeeper), store metrics in a database
// - others...
type Repository struct {
	requests []*stats.Request
	mutex    sync.Mutex
}

// NewRepository returns a new instance the repository to store stats in memory
func NewRepository() *Repository {
	return &Repository{requests: []*stats.Request{}}
}

// GetAll stats from memory
func (r *Repository) GetAll() []*stats.Request {
	return r.requests
}

// Increment number of hits in memory
func (r *Repository) Increment(path string, input map[string]interface{}) {
	// Using mutex to prevent when concurrent requests to write this shared resource
	// at the same time, hence having a race condition problem
	r.mutex.Lock()
	{
		req := r.find(path, input)
		if req == nil {
			req = &stats.Request{
				Path:       path,
				Parameters: input,
			}
			r.requests = append(r.requests, req)
		}
		req.NbHits++
	}
	r.mutex.Unlock()
}

func (r *Repository) find(path string, p map[string]interface{}) *stats.Request {
	var result *stats.Request
	for _, req := range r.requests {
		if req.Path == path && reflect.DeepEqual(req.Parameters, p) {
			result = req
		}
	}
	return result
}
