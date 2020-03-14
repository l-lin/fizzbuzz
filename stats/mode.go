//go:generate stringer -type=Mode
package stats

import "strings"

// Mode of the primitive to transform
type Mode int

const (
	// Memory mode to store the statistics
	Memory Mode = iota
	// Kafka mode to track requests
	Kafka
)

// Modes contains all the possible modes
var Modes = [...]Mode{Memory, Kafka}

// ModesToSlice transforms the modes into a slice
func ModesToSlice() []string {
	result := []string{}
	for _, m := range Modes {
		result = append(result, strings.ToLower(m.String()))
	}
	return result
}

// GetMode from given string
func GetMode(s string) Mode {
	for _, m := range Modes {
		if strings.ToLower(m.String()) == strings.ToLower(s) {
			return m
		}
	}
	return Memory
}
