package stats

// Request stats data structure
type Request struct {
	Path       string                 `json:"path"`
	NbHits     int                    `json:"nbHits"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}
