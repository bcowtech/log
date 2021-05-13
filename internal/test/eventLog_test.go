package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/bcowtech/log"
)

func TestWriteMarshal(t *testing.T) {
	var buffer bytes.Buffer
	writer := &log.JsonTextWriter{
		Stream: &buffer,
	}

	var (
		logType   = log.PASS
		eventID   = "511c650b-bf04-4da3-9f96-c122ed5a1ce9"
		category  = "demo-service"
		source    = "192.168.56.5"
		version   = "v1.0.1-alpha"
		message   = "a pass message"
		timestamp = time.Date(2020, 11, 7, 13, 18, 12, 0, time.UTC)
	)

	writer.Write(
		logType,
		eventID,
		category,
		source,
		version,
		message,
		timestamp)

	fmt.Printf("> buffer\n< %#v\n", buffer.String())
	{
		model := make(map[string]interface{})
		err := json.Unmarshal(buffer.Bytes(), &model)
		if err != nil {
			t.Errorf("should be able to be serialized")
		}
		fmt.Printf("> model\n< %#v\n", model)

		if model["type"] != logType.Name() {
			t.Errorf("assert 'type':: expected '%v', got '%v'", logType.Name(), model["type"])
		}
		expectedSeverity := strconv.FormatInt(int64(logType.Severity()), 10)
		if strconv.FormatFloat(model["severity"].(float64), 'G', 20, 64) != expectedSeverity {
			t.Errorf("assert 'severity':: expected '%v', got '%v'", expectedSeverity, model["severity"])
		}
		if model["event_id"] != eventID {
			t.Errorf("assert 'event_id':: expected '%v', got '%v'", eventID, model["event_id"])
		}
		if model["category"] != category {
			t.Errorf("assert 'categoty':: expected '%v', got '%v'", category, model["category"])
		}
		if model["source"] != source {
			t.Errorf("assert 'source':: expected '%v', got '%v'", source, model["source"])
		}
		if model["version"] != version {
			t.Errorf("assert 'version':: expected '%v', got '%v'", version, model["version"])
		}
		if model["message"] != message {
			t.Errorf("assert 'message':: expected '%v', got '%v'", message, model["message"])
		}
		expectedTimestamp := strconv.FormatInt(timestamp.UnixNano()/int64(time.Millisecond), 10)
		if strconv.FormatFloat(model["timestamp"].(float64), 'G', 20, 64) != expectedTimestamp {
			t.Errorf("assert 'timestamp':: expected '%v', got '%v'", expectedTimestamp, model["timestamp"])
		}
	}

	{
		event := new(log.EventLog)
		err := json.Unmarshal(buffer.Bytes(), event)
		if err != nil {
			t.Errorf("should be able to be serialized")
		}
		fmt.Printf("> EventLog\n< %#v\n", event)

		if event.Type != logType {
			t.Errorf("assert 'EventLog.Type':: expected '%v', got '%v'", logType, event.Type)
		}
		if event.EventID != eventID {
			t.Errorf("assert 'EventLog.EventID':: expected '%v', got '%v'", eventID, event.EventID)
		}
		if event.Category != category {
			t.Errorf("assert 'EventLog.Category':: expected '%v', got '%v'", category, event.Category)
		}
		if event.Source != source {
			t.Errorf("assert 'EventLog.Source':: expected '%v', got '%v'", source, event.Source)
		}
		if event.Version != version {
			t.Errorf("assert 'EventLog.Version':: expected '%v', got '%v'", version, event.Version)
		}
		if event.Message != message {
			t.Errorf("assert 'EventLog.Message':: expected '%v', got '%v'", message, event.Message)
		}
		expectedTimestamp := timestamp.UnixNano() / int64(time.Millisecond)
		if event.Timestamp != expectedTimestamp {
			t.Errorf("assert 'EventLog.Timestamp':: expected '%v', got '%v'", expectedTimestamp, event.Timestamp)
		}
	}

}
