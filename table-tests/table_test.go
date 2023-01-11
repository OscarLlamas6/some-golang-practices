package main

import "testing"

func TestToSnakeCase(t *testing.T) {
	type testCase struct {
		description string
		input       string
		expected    string
	}

	testCases := []testCase{
		{
			description: "empty string",
			input:       "",
			expected:    "",
		},
		{
			description: "single word",
			input:       "Hello",
			expected:    "hello",
		},
		{
			description: "two words (camel case)",
			input:       "HelloWorld",
			expected:    "hello_world",
		},
		{
			description: "two words with space",
			input:       "Hello World",
			expected:    "hello_world",
		},
		{
			description: "two words with space and underscore",
			input:       "Hello_World",
			expected:    "hello_world",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := ToSnakeCase(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
