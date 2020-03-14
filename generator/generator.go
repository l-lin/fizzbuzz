package generator

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

// ToParameters convert a generic map into fizz buzz parameters
func ToParameters(m map[string]interface{}) (int, int, int, string, string, error) {
	pc := parametersConverter{parameters: m}
	int1 := pc.convertToInt("int1")
	int2 := pc.convertToInt("int2")
	limit := pc.convertToInt("limit")
	str1 := pc.convertToString("str1")
	str2 := pc.convertToString("str2")
	if pc.err != nil {
		return 0, 0, 0, "", "", pc.err
	}
	return int1, int2, limit, str1, str2, nil
}

type parametersConverter struct {
	parameters map[string]interface{}
	err        error
}

func (p *parametersConverter) convertToInt(key string) int {
	if p.err != nil {
		return 0
	}
	v, ok := p.parameters[key]
	if !ok {
		p.err = fmt.Errorf(`Parameter "%s" is required`, key)
		return 0
	}
	i, ok := v.(int)
	if !ok {
		f, ok := v.(float64)
		if !ok {
			p.err = fmt.Errorf(`Parameter "%s" is required`, key)
			return 0
		}
		return int(f)
	}
	return i
}

func (p *parametersConverter) convertToString(key string) string {
	if p.err != nil {
		return ""
	}
	v, ok := p.parameters[key]
	if !ok {
		p.err = fmt.Errorf(`Parameter "%s" is required`, key)
		return ""
	}
	s, ok := v.(string)
	if !ok {
		p.err = fmt.Errorf(`Parameter "%s" is required`, key)
		return ""
	}
	return s
}
