package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " helloworld  ",
			expected: []string{"helloworld"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("Actual output: %v", actual)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Length of output does not match expected length")
		}
		for i := range actual {
			word := actual[i]
			t.Logf("Actual word: %v", word)
			expectedWord := c.expected[i]
			t.Logf("Expected word: %v", expectedWord)
			if word != expectedWord {
				t.Errorf("Unexpected output.")
			}
		}
	}

}
