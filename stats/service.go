package stats

import "io"

// Repository to store stats
type Repository interface {
	GetAll() []*Request
	Increment(path string, input map[string]interface{}) error
	io.Closer
}

// Service to store stats
type Service struct {
	r Repository
}

// NewService creates a service to handle request stats
func NewService(r Repository) *Service {
	return &Service{r}
}

// GetAll request stats from the storage
func (s *Service) GetAll() []*Request {
	return s.r.GetAll()
}

// GetMostUsed request
func (s *Service) GetMostUsed() *Request {
	requests := s.GetAll()
	var mostUsed *Request
	for _, req := range requests {
		if mostUsed == nil {
			mostUsed = req
		} else {
			if req.NbHits > mostUsed.NbHits {
				mostUsed = req
			}
		}
	}
	return mostUsed
}

// Increment number of hits for a given path
func (s *Service) Increment(path string, input map[string]interface{}) error {
	return s.r.Increment(path, input)
}
