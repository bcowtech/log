package log

import (
	"io"
	"os"
	"strconv"
	"sync"
)

var (
	once     sync.Once
	instance *Logger
)

func GetLogger() *Logger {
	innerSetUp(nil)
	return instance
}

func MixWriters(writers ...Writer) Writer {
	return &compositeWriter{
		writers: writers,
	}
}

func MixStreams(writers ...io.Writer) io.Writer {
	return &compositeStream{
		writers: writers,
	}
}

func MixFilters(filters ...Filter) Filter {
	return &compositeFilter{
		filters: filters,
	}
}

// Setup the global Logger. If argument logger is nil, the Logger
// will use PlainTextWriter with stdout stream.
func SetUp(logger *Logger) {
	if instance != nil {
		panic("cannot setup global Logger")
	}
	innerSetUp(logger)
}

func Write(
	logType EventLogType,
	message string) {

	GetLogger().Write(logType, message)
}

func WriteEventLog(log EventLog) {
	GetLogger().getWriter().WriteEventLog(log)
}

func innerSetUp(logger *Logger) {
	once.Do(func() {
		if logger == nil {
			setUpDefaultLogger()
		} else {
			instance = logger
		}
	})
}

func setUpDefaultLogger() {
	var category, source string
	{
		hostname, _ := os.Hostname()
		workdir, _ := os.Getwd()
		category = hostname + ":" + workdir
	}
	{
		pid := os.Getpid()
		source = strconv.FormatInt(int64(pid), 10)
	}

	instance = NewLogger(&Config{
		Category: category,
		Source:   source,
		Version:  "",
		Writer: &PlainTextWriter{
			Stream: os.Stdout,
		},
	})
}
