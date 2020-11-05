package log

import (
	"fmt"
	"io"
	"time"
)

type DateTimeFormatter func(t time.Time) string

type PlainTextWriter struct {
	Stream io.Writer

	DateTimeFormatter DateTimeFormatter
}

func (w *PlainTextWriter) Write(
	logType EventLogType,
	eventID string,
	category string,
	source string,
	version string,
	message string,
	timestamp time.Time) {

	w.ensureDateTimeFormatter()

	//  datetime type [category/source - version] message (#event-id)
	w.write(
		"%s %s [%s/%s - %s] %s (#%s)\n",
		w.DateTimeFormatter(timestamp),
		logType.Name(),
		category,
		source,
		version,
		message,
		eventID)
}

func (w *PlainTextWriter) WriteEventLog(log EventLog) {
	timestamp := transformMillisecondTime(log.Timestamp)

	w.ensureDateTimeFormatter()

	/*
	 *  datetime type [category/source - version] message (#event-id)
	 *
	 *  Request:
	 *  method path?query-string
	 *  header
	 *  body
	 *
	 *  Response:
	 *  status-code
	 *  header
	 *  body
	 *
	 *  Error:
	 *  error
	 *
	 *  Metric:
	 *  elapsed time: elapsed-time
	 *  depth: depth
	 *  response body bytes: response-body-bytes
	 */
	w.write(
		"%s %s [%s/%s - %s] %s (#%s)\n"+
			"\n"+
			"Request:\n"+
			"%s %s?%s\n"+
			"%s\n"+
			"%s\n"+
			"\n"+
			"Response:\n"+
			"%s\n"+
			"%s\n"+
			"%s\n"+
			"\n"+
			"Error:\n"+
			"%s\n"+
			"\n"+
			"Metric:\n"+
			"elapsed time: %d\n"+
			"depth: %d\n"+
			"response body bytes: %d\n"+
			"\n",
		w.DateTimeFormatter(timestamp),
		log.Type.Name(),
		log.Category,
		log.Source,
		log.Version,
		log.Message,
		log.EventID,
		// Request
		log.Details.Request.Method,
		log.Details.Request.Path,
		log.Details.Request.QueryString,
		log.Details.Request.Header,
		log.Details.Request.Body,
		// Response
		log.Details.Response.StatusCode,
		log.Details.Response.Header,
		log.Details.Response.Body,
		// Error
		log.Details.Error,
		// Metric
		log.Metric.ElapsedTime,
		log.Metric.Depth,
		log.Metric.ResponseBodyBytes)
}

func (w *PlainTextWriter) write(format string, a ...interface{}) {
	fmt.Fprintf(w.Stream, format, a...)
}

func (w *PlainTextWriter) ensureDateTimeFormatter() {
	if w.DateTimeFormatter == nil {
		w.DateTimeFormatter = func(t time.Time) string {
			return t.UTC().Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
}
