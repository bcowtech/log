package log

import "time"

type Writer interface {
	Write(
		logType EventLogType,
		eventID string,
		category string,
		source string,
		version string,
		message string,
		timestamp time.Time)
	WriteEventLog(log *EventLog)
}
