package log

import (
	"encoding/json"
	"io"
	"time"
)

type JsonTextWriter struct {
	Stream io.Writer

	Indent string
}

func (w *JsonTextWriter) Write(
	logType EventLogType,
	eventID string,
	category string,
	source string,
	version string,
	message string,
	timestamp time.Time) {

	w.write(&struct {
		Type      string `json:"type"`
		Severity  int    `json:"severity"`
		EventId   string `json:"event_id"`
		Category  string `json:"category"`
		Source    string `json:"source"`
		Version   string `json:"version"`
		Message   string `json:"message"`
		Timestamp int64  `json:"timestamp"`
	}{
		Type:      logType.Name(),
		Severity:  logType.Severity(),
		EventId:   eventID,
		Category:  category,
		Source:    source,
		Version:   version,
		Message:   message,
		Timestamp: (timestamp.UnixNano() / int64(time.Millisecond)),
	})
}

func (w *JsonTextWriter) WriteEventLog(log *EventLog) {
	w.write(log)
}

func (w *JsonTextWriter) write(v interface{}) {
	encoder := json.NewEncoder(w.Stream)
	encoder.SetIndent("", w.Indent)
	encoder.Encode(v)
}
