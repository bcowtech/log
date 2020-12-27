package elasticsearch

import "time"

type (
	Option struct {
		Address      string
		QueryTimeout time.Duration // no-timeout(0) https://golang.org/pkg/net/http/#Client.Timeout
	}

	IndexProvider interface {
		IndexName() string
	}
)

type ImmutableIndexProvider string

func (p ImmutableIndexProvider) IndexName() string {
	return string(p)
}
