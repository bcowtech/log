package log

import "time"

type compositeFilter struct {
	filters []Filter
}

func (w *compositeFilter) CanWrite(
	logType EventLogType,
	eventID string,
	category string,
	source string,
	version string,
	message string,
	timestamp time.Time) bool {

	for _, filter := range w.filters {
		ok := filter.CanWrite(
			logType,
			eventID,
			category,
			source,
			version,
			message,
			timestamp)
		if !ok {
			return false
		}
	}
	return true
}

func (w *compositeFilter) CanWriteEventLog(log EventLog) bool {
	for _, filter := range w.filters {
		ok := filter.CanWriteEventLog(log)
		if !ok {
			return false
		}
	}
	return true
}
