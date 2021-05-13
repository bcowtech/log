package test

import (
	"testing"

	log "github.com/bcowtech/log"
)

func TestEventLogType(t *testing.T) {
	if log.NONE.IsValid() {
		t.Errorf("assert 'EventLogType.NONE.IsValid()':: expected '%v', got '%v'", true, log.NONE.IsValid())
	}
	if !log.PASS.IsValid() {
		t.Errorf("assert 'EventLogType.PASS.IsValid()':: expected '%v', got '%v'", false, log.PASS.IsValid())
	}
	if !log.FAIL.IsValid() {
		t.Errorf("assert 'EventLogType.FAIL.IsValid()':: expected '%v', got '%v'", false, log.FAIL.IsValid())
	}
}
