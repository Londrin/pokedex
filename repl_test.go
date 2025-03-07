package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " testing words  with  spaces",
			expected: []string{"testing", "words", "with", "spaces"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words and received: %d words", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected: %s, Received: %s", expectedWord, word)
			}
		}
	}
}
