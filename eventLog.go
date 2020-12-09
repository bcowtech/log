package log

import (
	"encoding/json"
)

type (
	// EventLog
	EventLog struct {
		Type      EventLogType `json:"type"`
		EventID   string       `json:"event_id"`
		Category  string       `json:"category"`
		Source    string       `json:"Source"`
		Version   string       `json:"version"`
		Message   string       `json:"message"`
		Details   Detail       `json:"details"`
		Metric    Metric       `json:"metric"`
		Timestamp int64        `json:"timestamp"` // using milliseconds
	}

	// Detail
	Detail struct {
		Request  HttpRequestDetail  `json:"request"`
		Response HttpResponseDetail `json:"response"`
		Error    string             `json:"error"`
	}

	// HttpRequestDetail
	HttpRequestDetail struct {
		Method      string `json:"method"`
		Path        string `json:"path"`
		QueryString string `json:"query_string"`
		Header      string `json:"header"`
		Body        string `json:"body"`
	}

	// HttpResponseDetail
	HttpResponseDetail struct {
		StatusCode string `json:"status_code"`
		Header     string `json:"header"`
		Body       string `json:"body"`
	}

	// Metric
	Metric struct {
		ElapsedTime       int32 `json:"elapsed_time"` // using milliseconds
		Depth             int32 `json:"depth"`
		ResponseBodyBytes int32 `json:"response_body_bytes"`
	}
)

func (l *EventLog) WriteTo(w Writer) {
	w.WriteEventLog(l)
}

func (l *EventLog) MarshalJSON() ([]byte, error) {
	type Alias EventLog
	return json.Marshal(&struct {
		Type     string `json:"type"`
		Severity int    `json:"severity"`
		*Alias
	}{
		Type:     l.Type.Name(),
		Severity: l.Type.Severity(),
		Alias:    (*Alias)(l),
	})
}

func (l *EventLog) UnmarshalJSON(data []byte) error {
	type Alias EventLog
	aux := &struct {
		Type string `json:"type"`
		*Alias
	}{
		Alias: (*Alias)(l),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	l.Type = ParseEventLogTypeName(aux.Type)
	return nil
}
