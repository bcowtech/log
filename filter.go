package log

import "time"

type Filter interface {
	CanWrite(
		logType EventLogType,
		eventID string,
		category string,
		source string,
		version string,
		message string,
		timestamp time.Time) bool
	CanWriteEventLog(log *EventLog) bool
}
