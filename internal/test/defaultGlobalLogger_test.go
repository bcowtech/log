package test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	log "github.com/bcowtech/log"
)

func TestDefaultGlobalLogger(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	log.Write(log.FAIL, "a failure message")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if len(out) == 0 {
		t.Errorf("message should not be empty")
	}
}

func TestSetupGlobalLoggerWhenExisted(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("global Logger should not be set again")
		}
	}()

	var buffer bytes.Buffer
	log.SetUp(log.NewLogger(&log.Config{
		Category: "demo-test",
		Source:   "localhost",
		Version:  "v1.0.0b",
		Writer: &log.PlainTextWriter{
			Stream: &buffer,
		},
	}))
}
