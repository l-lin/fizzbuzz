package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/l-lin/fizzbuzz/stats/memory"
)

func TestFizzBuzzHandler(t *testing.T) {
	type expected struct {
		responseBody   string
		responseStatus int
	}
	var tests = map[string]struct {
		given    map[string]interface{}
		expected expected
	}{
		"classic": {
			given: map[string]interface{}{
				"int1": 3, "int2": 5, "limit": 30,
				"str1": "Fizz", "str2": "Buzz",
			},
			expected: expected{
				responseBody: `{"result":["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz","16","17","Fizz","19","Buzz","Fizz","22","23","Fizz","Buzz","26","Fizz","28","29","FizzBuzz"]}
`,
				responseStatus: http.StatusOK,
			},
		},
		"limit not set": {
			given: map[string]interface{}{
				"int1": 3, "int2": 5,
				"str1": "Fizz", "str2": "Buzz",
			},
			expected: expected{
				responseBody: `{"error":"Parameter \"limit\" is required"}
`,
				responseStatus: http.StatusBadRequest,
			},
		},
		"no input set": {
			given: map[string]interface{}{},
			expected: expected{
				responseBody: `{"error":"Parameter \"int1\" is required"}
`,
				responseStatus: http.StatusBadRequest,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			router := NewRouter(memory.NewRepository())
			w := httptest.NewRecorder()
			c, err := json.Marshal(tt.given)
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest("POST", fizzBuzzRoute, bytes.NewReader(c))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)
			if w.Code != tt.expected.responseStatus {
				t.Errorf("expected %v, actual %v", tt.expected.responseStatus, w.Code)
			}
			actualResponseBody := w.Body.String()
			if actualResponseBody != tt.expected.responseBody {
				t.Errorf("expected:\n%v\nactual:\n%v", tt.expected.responseBody, actualResponseBody)
			}
		})
	}
}
