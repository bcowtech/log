package log

import (
	"io"
)

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
