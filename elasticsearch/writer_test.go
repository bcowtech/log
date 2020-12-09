package elasticsearch

import (
	"os"
	"testing"
	"time"

	"gitlab.bcowtech.de/bcow-go/log"
)

func TestPing(t *testing.T) {
	writer := NewWriter(&Option{
		Address:          "http://192.168.56.54:9200",
		IndexPartitioner: ImmutableIndexPartitioner("service-access-log-000001"),
		QueryTimeout:     15 * time.Second,
	})

	ok := writer.Ping(3 * time.Second)
	if !ok {
		t.Errorf("Ping() should be ok")
	}
}

var (
	elasticsearchAddress = os.Getenv("ELASTICSEARCH_ADDRESS")
)

func TestWriteEventLog(t *testing.T) {
	writer := NewWriter(&Option{
		Address:          elasticsearchAddress,
		IndexPartitioner: ImmutableIndexPartitioner("service-access-log-000001"),
		QueryTimeout:     15 * time.Second,
	})

	eventLog := &log.EventLog{
		EventID:  "192.168.56.54-0006",
		Type:     log.PASS,
		Category: "WalletService",
		Source:   "10.4.10.6",
		Version:  "v0.5.1b",
		Details: log.Detail{
			Request: log.HttpRequestDetail{
				Method:      "GET",
				Path:        "/Player/Balance",
				QueryString: "arg1=eins&arg2=zwei&arg3=drei",
				Header:      "Content-Type: application/json\r\nUser-Agent: Debian APT-HTTP/1.3 (0.9.7.9)\r\n",
				Body:        "some request body",
			},
			Response: log.HttpResponseDetail{
				StatusCode: "304",
				Header:     "X-Response-Header: f/twXyy",
				Body:       "some response body",
			},
		},
		Metric: log.Metric{
			ElapsedTime:       122,
			ResponseBodyBytes: 15,
		},
		Timestamp: time.Date(2020, 12, 10, 5, 18, 6, 0, time.UTC).UnixNano() / int64(time.Millisecond),
	}
	writer.WriteEventLog(eventLog)
}
