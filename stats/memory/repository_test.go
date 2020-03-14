package memory

import (
	"sync"
	"testing"
)

const (
	nbHits = 500
	path   = "/foobar"
)

var parameters = map[string]interface{}{
	"Int1":  3,
	"Int2":  5,
	"Limit": 10,
	"Str1":  "Fizz",
	"Str2":  "Buzz",
}

func TestIncrement(t *testing.T) {
	r := NewRepository()
	var wg sync.WaitGroup
	for i := 0; i < nbHits; i++ {
		wg.Add(1)
		go func(path string) {
			r.Increment(path, parameters)
			wg.Done()
		}(path)
	}
	wg.Wait()

	req := r.find(path, parameters)
	if req.NbHits != nbHits {
		t.Errorf("expected nb hits %d, got %d", nbHits, req.NbHits)
	}
}
