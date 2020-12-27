package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"gitlab.bcowtech.de/bcow-go/log"
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
			t.Errorf("assert 'type':: excepted '%v', got '%v'", logType.Name(), model["type"])
		}
		exceptedSeverity := strconv.FormatInt(int64(logType.Severity()), 10)
		if strconv.FormatFloat(model["severity"].(float64), 'G', 20, 64) != exceptedSeverity {
			t.Errorf("assert 'severity':: excepted '%v', got '%v'", exceptedSeverity, model["severity"])
		}
		if model["event_id"] != eventID {
			t.Errorf("assert 'event_id':: excepted '%v', got '%v'", eventID, model["event_id"])
		}
		if model["category"] != category {
			t.Errorf("assert 'categoty':: excepted '%v', got '%v'", category, model["category"])
		}
		if model["source"] != source {
			t.Errorf("assert 'source':: excepted '%v', got '%v'", source, model["source"])
		}
		if model["version"] != version {
			t.Errorf("assert 'version':: excepted '%v', got '%v'", version, model["version"])
		}
		if model["message"] != message {
			t.Errorf("assert 'message':: excepted '%v', got '%v'", message, model["message"])
		}
		exceptedTimestamp := strconv.FormatInt(timestamp.UnixNano()/int64(time.Millisecond), 10)
		if strconv.FormatFloat(model["timestamp"].(float64), 'G', 20, 64) != exceptedTimestamp {
			t.Errorf("assert 'timestamp':: excepted '%v', got '%v'", exceptedTimestamp, model["timestamp"])
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
			t.Errorf("assert 'EventLog.Type':: excepted '%v', got '%v'", logType, event.Type)
		}
		if event.EventID != eventID {
			t.Errorf("assert 'EventLog.EventID':: excepted '%v', got '%v'", eventID, event.EventID)
		}
		if event.Category != category {
			t.Errorf("assert 'EventLog.Category':: excepted '%v', got '%v'", category, event.Category)
		}
		if event.Source != source {
			t.Errorf("assert 'EventLog.Source':: excepted '%v', got '%v'", source, event.Source)
		}
		if event.Version != version {
			t.Errorf("assert 'EventLog.Version':: excepted '%v', got '%v'", version, event.Version)
		}
		if event.Message != message {
			t.Errorf("assert 'EventLog.Message':: excepted '%v', got '%v'", message, event.Message)
		}
		exceptedTimestamp := timestamp.UnixNano() / int64(time.Millisecond)
		if event.Timestamp != exceptedTimestamp {
			t.Errorf("assert 'EventLog.Timestamp':: excepted '%v', got '%v'", exceptedTimestamp, event.Timestamp)
		}
	}

}
