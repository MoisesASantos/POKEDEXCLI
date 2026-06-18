package repl

import (
	"testing"
	"reflect"
)

/*If we want to dig in to the coverage report the go tool has several options to print the coverage report. We can use go tool cover -func to break down the coverage per function:

% go tool cover -func=c.out
split/split.go:8:       Split          100.0%
total:                  (statements)   100.0% */


func TestCleanInput(t *testing.T) {
    
	cases := []struct {
	input    string
	expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Moises Dos Santos",
			expected: []string{"Moises", "Dos", "Santos"},
		},
		{
			input: "Moises Dos Santos     ",
			expected: []string{"Moises", "Dos", "Santos"},
		},
		{
			input: "          Moises Dos Santos",
			expected: []string{"Moises", "Dos", "Santos"},
		},
		{
			input: "Moises | Dos | Santos",
			expected: []string{"Moises", "|", "Dos", "|", "Santos"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if !reflect.DeepEqual(c.expected, actual) {
        	t.Errorf("expected: %v, got: %v", c.expected, actual)
    	}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("The word don't match\n")
			}
			
		}
	}
}
