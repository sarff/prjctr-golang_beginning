package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFormatNumber(t *testing.T) {
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
			result, err := FormatNumber(test.input)
			if err != nil {
				t.Error(err)
			}
			if result != test.expected {
				t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
			}
		})
	}
}

func TestPhoneNormalize(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		input := strings.NewReader(`Zinadin Zidan - 050 123 45 67
Roberto Carlos - 044-456-78-90
Zinadin Zidan -`)
		want := `Zinadin Zidan - (050) 12-34-567
Roberto Carlos - (044) 45-67-890
`
		var buf bytes.Buffer
		err := PhoneNormalize(input, &buf)
		if err != nil {
			t.Fatal("unexpected error:", err)
		}

		got := buf.String()
		if got != want {
			t.Errorf("unexpected output:\n Want %q\nGot: %q", want, got)
		}
	})

	t.Run("invalid input", func(t *testing.T) {
		var buf bytes.Buffer
		test2Input := "Zinadin Zidan - "
		err := PhoneNormalize(strings.NewReader(test2Input), &buf)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
