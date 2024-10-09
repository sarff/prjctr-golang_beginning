package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	log := New(&buf)
	log.Logger.Info("TEST", "KEY", "VALUE")

	logOutput := buf.String()

	if !strings.Contains(logOutput, "TEST") {
		t.Errorf("Expected log to contain 'TEST', but got: %s", logOutput)
	}

	if !strings.Contains(logOutput, "KEY") {
		t.Errorf("Expected log to contain 'KEY', but got: %s", logOutput)
	}

	if !strings.Contains(logOutput, "VALUE") {
		t.Errorf("Expected log to contain 'VALUE', but got: %s", logOutput)
	}
}
