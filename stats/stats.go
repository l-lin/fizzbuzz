package stats

import "github.com/l-lin/fizzbuzz/model"

// Request stats data structure
type Request struct {
	Path             string `json:"path"`
	NbHits           int    `json:"nbHits"`
	model.Parameters `json:"parameters"`
}
