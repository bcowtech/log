package log

import "time"

type SeverityFilter struct {
	Severity int
}

func (w *SeverityFilter) CanWrite(
	logType EventLogType,
	eventID string,
	category string,
	source string,
	version string,
	message string,
	timestamp time.Time) bool {

	return w.canWrite(logType)
}

func (w *SeverityFilter) CanWriteEventLog(log *EventLog) bool {
	return w.canWrite(log.Type)
}

func (w *SeverityFilter) canWrite(logType EventLogType) bool {
	return logType.Severity() == 0 || logType.Severity() >= w.Severity
}
