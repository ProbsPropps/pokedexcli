package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases:= []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		
		{
			input: "BULBASAUR CHARMANDER SQUIRTLE",
			expected: []string{"bulbasaur", "charmander", "squirtle"},
		},

		{
			input: "weaVILe     mamoswine electrivire    ",
			expected: []string{"weavile", "mamoswine", "electrivire"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("clean input did not return the length expected")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {

				t.Errorf("word: %s did not match the expected output: %s", word, expectedWord)
			}
		}
	}
}
