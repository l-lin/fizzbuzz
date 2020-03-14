package generator

import "testing"

func TestGenerate(t *testing.T) {
	type given struct {
		int1, int2, limit int
		str1, str2        string
	}
	var tests = map[string]struct {
		given         given
		expected      []string
		expectedError bool
	}{
		"classic": {
			given: given{
				int1:  3,
				int2:  5,
				limit: 30,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			expected:      []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz", "26", "Fizz", "28", "29", "FizzBuzz"},
			expectedError: false,
		},
		"limit 0": {
			given: given{
				int1:  3,
				int2:  5,
				limit: 0,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			expected:      []string{},
			expectedError: false,
		},
		"limit 1": {
			given: given{
				int1:  3,
				int2:  5,
				limit: 1,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			expected:      []string{"1"},
			expectedError: false,
		},
		"int1 & int2 with same value": {
			given: given{
				int1:  2,
				int2:  2,
				limit: 10,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			expected:      []string{"1", "FizzBuzz", "3", "FizzBuzz", "5", "FizzBuzz", "7", "FizzBuzz", "9", "FizzBuzz"},
			expectedError: false,
		},
		"empty str1": {
			given: given{
				int1:  3,
				int2:  5,
				limit: 10,
				str1:  "",
				str2:  "Buzz",
			},
			expected:      []string{"1", "2", "", "4", "Buzz", "", "7", "8", "", "Buzz"},
			expectedError: false,
		},
		"empty str2": {
			given: given{
				int1:  3,
				int2:  5,
				limit: 10,
				str1:  "Fizz",
				str2:  "",
			},
			expected:      []string{"1", "2", "Fizz", "4", "", "Fizz", "7", "8", "Fizz", ""},
			expectedError: false,
		},
		"empty str1 & str2": {
			given: given{
				int1:  2,
				int2:  2,
				limit: 10,
				str1:  "",
				str2:  "",
			},
			expected:      []string{"1", "", "3", "", "5", "", "7", "", "9", ""},
			expectedError: false,
		},
		// Edge cases
		"negative limit": {
			given: given{
				int1:  2,
				int2:  2,
				limit: -10,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			expectedError: true,
		},
		"negative int1": {
			given: given{
				int1:  -2,
				int2:  2,
				limit: 10,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			expectedError: true,
		},
		"negative int2": {
			given: given{
				int1:  2,
				int2:  -2,
				limit: 10,
				str1:  "Fizz",
				str2:  "Buzz",
			},
			expectedError: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual, actualErr := Generate(tt.given.int1, tt.given.int2, tt.given.limit, tt.given.str1, tt.given.str2)
			if tt.expectedError {
				if actualErr == nil {
					t.Error("expected a error, got nothing")
				}
			} else {
				if len(actual) != len(tt.expected) {
					t.Errorf("expected length %d, actual length %d", len(tt.expected), len(actual))
				} else {
					for i := 0; i < len(tt.expected); i++ {
						if actual[i] != tt.expected[i] {
							t.Errorf("%d: expected %v, actual %v", i, tt.expected[i], actual[i])
						}
					}
				}
			}
		})
	}
}

func TestToParameters(t *testing.T) {
	type expected struct {
		int1, int2, limit int
		str1, str2        string
	}
	var tests = map[string]struct {
		given       map[string]interface{}
		expected    expected
		expectedErr bool
	}{
		"happy path": {
			given: map[string]interface{}{
				"int1": 3, "int2": 5, "limit": 30,
				"str1": "Fizz", "str2": "Buzz",
			},
			expected: expected{
				int1: 3, int2: 5, limit: 30,
				str1: "Fizz", str2: "Buzz",
			},
			expectedErr: false,
		},
		"missing int1": {
			given: map[string]interface{}{
				"int2": 5, "limit": 30,
				"str1": "Fizz", "str2": "Buzz",
			},
			expected: expected{
				int1: 0, int2: 0, limit: 0,
				str1: "", str2: "",
			},
			expectedErr: true,
		},
		"missing int2": {
			given: map[string]interface{}{
				"int1": 5, "limit": 30,
				"str1": "Fizz", "str2": "Buzz",
			},
			expected: expected{
				int1: 0, int2: 0, limit: 0,
				str1: "", str2: "",
			},
			expectedErr: true,
		},
		"missing limit": {
			given: map[string]interface{}{
				"int1": 3, "int2": 5,
				"str1": "Fizz", "str2": "Buzz",
			},
			expected: expected{
				int1: 0, int2: 0, limit: 0,
				str1: "", str2: "",
			},
			expectedErr: true,
		},
		"missing str1": {
			given: map[string]interface{}{
				"int1": 3, "int2": 5, "limit": 30,
				"str1": "Fizz",
			},
			expected: expected{
				int1: 0, int2: 0, limit: 0,
				str1: "", str2: "",
			},
			expectedErr: true,
		},
		"missing str2": {
			given: map[string]interface{}{
				"int1": 3, "int2": 5, "limit": 30,
				"str2": "Buzz",
			},
			expected: expected{
				int1: 0, int2: 0, limit: 0,
				str1: "", str2: "",
			},
			expectedErr: true,
		},
		"using float64 in int1": {
			given: map[string]interface{}{
				"int1": float64(3), "int2": 5, "limit": 30,
				"str1": "Fizz", "str2": "Buzz",
			},
			expected: expected{
				int1: 3, int2: 5, limit: 30,
				str1: "Fizz", str2: "Buzz",
			},
			expectedErr: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actualInt1, actualInt2, actualLimit, actualStr1, actualStr2, actualErr := ToParameters(tt.given)
			if tt.expectedErr {
				if actualErr == nil {
					t.Error("expected an error, got no error")
				}
			} else {
				if actualErr != nil {
					t.Errorf("expected no error, got error: %v", actualErr)
				}
				if actualInt1 != tt.expected.int1 {
					t.Errorf("expected int1 %d, actual %d", tt.expected.int1, actualInt1)
				}
				if actualInt2 != tt.expected.int2 {
					t.Errorf("expected int2 %d, actual %d", tt.expected.int2, actualInt2)
				}
				if actualLimit != tt.expected.limit {
					t.Errorf("expected limit %d, actual %d", tt.expected.limit, actualLimit)
				}
				if actualStr1 != tt.expected.str1 {
					t.Errorf("expected str1 %s, actual %s", tt.expected.str1, actualStr1)
				}
				if actualStr2 != tt.expected.str2 {
					t.Errorf("expected str2 %s, actual %s", tt.expected.str2, actualStr2)
				}
			}
		})
	}

}
