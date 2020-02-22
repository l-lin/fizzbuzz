package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/l-lin/fizzbuzz/model"
)

func TestFizzBuzzHandler(t *testing.T) {
	type expected struct {
		responseBody   string
		responseStatus int
	}
	var tests = map[string]struct {
		given    model.Parameters
		expected expected
	}{
		"classic": {
			given: model.Parameters{
				Int1: 3, Int2: 5, Limit: 30,
				Str1: "Fizz", Str2: "Buzz",
			},
			expected: expected{
				responseBody: `{"result":["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz","16","17","Fizz","19","Buzz","Fizz","22","23","Fizz","Buzz","26","Fizz","28","29","FizzBuzz"]}
`,
				responseStatus: http.StatusOK,
			},
		},
		"limit not set": {
			given: model.Parameters{
				Int1: 3, Int2: 5,
				Str1: "Fizz", Str2: "Buzz",
			},
			expected: expected{
				responseBody: `{"error":"Key: 'Parameters.Limit' Error:Field validation for 'Limit' failed on the 'required' tag"}
`,
				responseStatus: http.StatusBadRequest,
			},
		},
		"no input set": {
			given: model.Parameters{},
			expected: expected{
				responseBody: `{"error":"Key: 'Parameters.Int1' Error:Field validation for 'Int1' failed on the 'required' tag\nKey: 'Parameters.Int2' Error:Field validation for 'Int2' failed on the 'required' tag\nKey: 'Parameters.Limit' Error:Field validation for 'Limit' failed on the 'required' tag\nKey: 'Parameters.Str1' Error:Field validation for 'Str1' failed on the 'required' tag\nKey: 'Parameters.Str2' Error:Field validation for 'Str2' failed on the 'required' tag"}
`,
				responseStatus: http.StatusBadRequest,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			router := NewRouter()
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