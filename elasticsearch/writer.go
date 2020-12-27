package elasticsearch

import (
	"time"

	bcowlog "gitlab.bcowtech.de/bcow-go/log"
)

type Writer struct {
	forwarder     *Forwarder
	indexProvider IndexProvider
}

func NewWriter(forwarder *Forwarder, indexProvider IndexProvider) *Writer {
	if forwarder == nil {
		panic("argument 'forwarder' cannot be nil")
	}
	if indexProvider == nil {
		panic("argument 'indexPartitioner' cannot be nil")
	}

	return &Writer{
		forwarder:     forwarder,
		indexProvider: indexProvider,
	}
}

func (w *Writer) Write(
	logType bcowlog.EventLogType,
	eventID string,
	category string,
	source string,
	version string,
	message string,
	timestamp time.Time) {

	eventLog := &bcowlog.EventLog{
		Type:      logType,
		EventID:   eventID,
		Category:  category,
		Source:    source,
		Version:   version,
		Message:   message,
		Timestamp: timestamp.UnixNano() / int64(time.Millisecond),
	}
	w.WriteEventLog(eventLog)
}

func (w *Writer) WriteEventLog(eventLog *bcowlog.EventLog) {
	w.forwarder.write(w.indexProvider, eventLog)
}
