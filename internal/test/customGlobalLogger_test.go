package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	log "github.com/bcowtech/log"
)

func TestCustomGlobalLogger(t *testing.T) {
	var buffer bytes.Buffer
	log.SetUp(log.NewLogger(&log.Config{
		Category: "demo-test",
		Source:   "localhost",
		Version:  "v1.0.0b",
		Writer: &log.PlainTextWriter{
			Stream: &buffer,
		},
	}))

	buffer.Reset()
	if len(buffer.String()) != 0 {
		t.Errorf("buffer should be empty")
	}

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	log.Write(log.FAIL, "a failure message")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if len(out) != 0 {
		t.Errorf("Stdout message should be empty")
	}

	if len(buffer.String()) == 0 {
		t.Errorf("buffered message should not be empty")
	}
}
