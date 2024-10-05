package sensor

import (
	"bytes"
	"strings"
	"sync"
	"testing"

	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/storage"

	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/centralsystem"
	"github.com/sarff/prjctr-golang_beginning/gocourse07/birdaviary/internal/logger"
)

func TestMainSensor_Start(t *testing.T) {
	var buf bytes.Buffer

	log := logger.New(&buf)
	db := storage.NewStorage(log)
	cs := centralsystem.NewCentralSystem(db, log)
	s := NewSensor(cs, log, "TestSensor")

	var wg sync.WaitGroup
	wg.Add(1)
	go s.Start(&wg)

	wg.Wait()

	logOutput := buf.String()

	if !strings.Contains(logOutput, "Starting sensor") {
		t.Errorf("Expected log to contain 'Starting sensor', but got: %s", logOutput)
	}

	if !strings.Contains(logOutput, "Stopping sensor") {
		t.Errorf("Expected log to contain 'Stopping sensor', but got: %s", logOutput)
	}
}
