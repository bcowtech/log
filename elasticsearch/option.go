package elasticsearch

import "time"

type (
	Option struct {
		Address          string
		IndexPartitioner IndexPartitioner
		QueryTimeout     time.Duration // no-timeout(0) https://golang.org/pkg/net/http/#Client.Timeout
	}

	IndexPartitioner interface {
		IndexName() string
	}
)

type ImmutableIndexPartitioner string

func (p ImmutableIndexPartitioner) IndexName() string {
	return string(p)
}
