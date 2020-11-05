package log

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Config struct {
		Category   string
		Source     string
		Version    string
		TextWriter Writer
	}

	Logger struct {
		category string
		source   string
		version  string
		writer   Writer
	}
)

func NewLogger(conf *Config) *Logger {
	return &Logger{
		category: conf.Category,
		source:   conf.Source,
		version:  conf.Version,
		writer:   conf.TextWriter,
	}
}

func (l *Logger) Write(
	logType EventLogType,
	message string) {

	var (
		eventID   = uuid.NewV4().String()
		timestamp = time.Now()
	)

	l.writer.Write(
		logType,
		eventID,
		l.category,
		l.source,
		l.version,
		message,
		timestamp)
}
