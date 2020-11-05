package log

import "time"

type FilterableTextWriter struct {
	Filter Filter
	Writer Writer
}

func (w *FilterableTextWriter) Write(
	logType EventLogType,
	eventID string,
	category string,
	source string,
	version string,
	message string,
	timestamp time.Time) {

	if w.Filter.CanWrite(
		logType,
		eventID,
		category,
		source,
		version,
		message,
		timestamp) {

		w.Writer.Write(
			logType,
			eventID,
			category,
			source,
			version,
			message,
			timestamp)
	}
}

func (w *FilterableTextWriter) WriteEventLog(log EventLog) {
	if w.Filter.CanWriteEventLog(log) {
		w.Writer.WriteEventLog(log)
	}
}
