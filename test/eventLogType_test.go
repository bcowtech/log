package test

import (
	"testing"

	log "gitlab.bcowtech.de/bcow-go/log"
)

func TestEventLogType(t *testing.T) {
	if log.NONE.IsValid() {
		t.Errorf("assert 'EventLogType.NONE.IsValid()':: excepted '%v', got '%v'", true, log.NONE.IsValid())
	}
	if !log.PASS.IsValid() {
		t.Errorf("assert 'EventLogType.PASS.IsValid()':: excepted '%v', got '%v'", false, log.PASS.IsValid())
	}
	if !log.FAIL.IsValid() {
		t.Errorf("assert 'EventLogType.FAIL.IsValid()':: excepted '%v', got '%v'", false, log.FAIL.IsValid())
	}
}
