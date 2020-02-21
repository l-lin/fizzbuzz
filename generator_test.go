package main

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
