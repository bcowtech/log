package log

import "time"

type compositeWriter struct {
	writers []Writer
}

func (w *compositeWriter) Write(
	logType EventLogType,
	eventID string,
	category string,
	source string,
	version string,
	message string,
	timestamp time.Time) {

	for _, writer := range w.writers {
		writer.Write(
			logType,
			eventID,
			category,
			source,
			version,
			message,
			timestamp)
	}
}

func (w *compositeWriter) WriteEventLog(log EventLog) {
	for _, writer := range w.writers {
		writer.WriteEventLog(log)
	}
}
