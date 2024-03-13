package main

import (
	"fmt"
	"github.com/Brandon689/bleve-subtitles/search"
	"testing"
)

func TestF(t *testing.T) {
	input := "ｂｒｉｄｇｅｃｄｆｓｔeddd"
	cleanedString := search.RemoveFullWidthCharacters(input)
	fmt.Println("Cleaned String:", cleanedString) // Output: Cleaned String:
}

func TestRemoveNumbers(t *testing.T) {
	// Test cases
	testCases := []struct {
		input    string
		expected string
	}{
		{"This is 123 an exampl000e 456 string with 789 numbers.", "This is  an example  string with  numbers."},
		{"1234567890", ""},
		{"No numbers here!", "No numbers here!"},
		{"", ""},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		// Call removeNumbers function
		cleanedString := search.RemoveNumbers(tc.input)
		fmt.Println(cleanedString)
		// Compare the result with the expected output
		if cleanedString != tc.expected {
			t.Errorf("removeNumbers(%q) = %q; want %q", tc.input, cleanedString, tc.expected)
		}
	}
}
