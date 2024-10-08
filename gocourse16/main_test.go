package main

import "testing"

func TestTrimNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"050 123 45 67", "(050) 12-34-567"},
		{"(067) 891-22-33", "(067) 89-12-233"},
		{"+380(93)111-44-55", "(093) 11-14-455"},
		{"044-456-78-90", "(044) 45-67-890"},
		{"0992951212", "(099) 29-51-212"},
		{"097 123 4567", "(097) 12-34-567"},
		{"+38050 777 8888", "(050) 77-78-888"},
		{"Invalid Number", "Invalid phone number"},
		{"віа3в3ау3434вав", "Invalid phone number"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := TrimNumber(test.input)
			if result != test.expected {
				t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
			}
		})
	}
}
