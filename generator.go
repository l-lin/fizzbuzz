package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Generate returns a list of strings with numbers from 1 to limit, where:
// - all multiples of int1 are replaced by str1
// - all multiples of int2 are replaced by str2
// - all multiples of int1 and int2 are replaced by str1str2
func Generate(int1, int2, limit int, str1, str2 string) ([]string, error) {
	if limit < 0 || int1 < 0 || int2 < 0 {
		return nil, fmt.Errorf("Numbers must be positive")
	}
	result := make([]string, limit)
	for i := 0; i < limit; i++ {
		var b strings.Builder
		// isMultiple is used to handle edge case where str1 & str2 are empty strings
		isMultiple := false
		if (i+1)%int1 == 0 {
			b.WriteString(str1)
			isMultiple = true
		}
		if (i+1)%int2 == 0 {
			b.WriteString(str2)
			isMultiple = true
		}
		if !isMultiple {
			b.WriteString(strconv.Itoa(i + 1))
		}
		result[i] = b.String()
	}
	return result, nil
}
