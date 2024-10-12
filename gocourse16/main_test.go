package main

import (
	"os"
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
	tempDir := t.TempDir()
	pathInputFile := tempDir + "/input.txt"
	pathOutputFile := tempDir + "/output.txt"

	inputData := "Zinadin Zidan - 050 123 45 67\nRoberto Carlos - 044-456-78-90\n"
	outputData := "Zinadin Zidan - (050) 12-34-567\nRoberto Carlos - (044) 45-67-890\n"
	inputFile, err := os.Create(pathInputFile)
	if err != nil {
		t.Error(err)
	}
	_, err = inputFile.WriteString(inputData)
	if err != nil {
		t.Error(err)
	}
	err = inputFile.Close()
	if err != nil {
		t.Fatal(err)
	}
	defer inputFile.Close()
	outputFile, err := os.Create(pathOutputFile)
	if err != nil {
		t.Error(err)
	}
	defer outputFile.Close()
	inputFile, err = os.Open(pathInputFile)
	if err != nil {
		t.Fatal(err)
	}
	defer inputFile.Close()
	err = PhoneNormalize(inputFile, outputFile)
	if err != nil {
		t.Error(err)
	}
	outputRead, err := os.ReadFile(outputFile.Name())
	if err != nil {
		t.Error(err)
	}
	if outputData != string(outputRead) {
		t.Errorf("unexpected output:\n Want %q\nGot: %q", outputData, string(outputRead))
	}
}
