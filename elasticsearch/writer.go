package elasticsearch

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	bcowlog "gitlab.bcowtech.de/bcow-go/log"
)

var (
	defaultLogger = log.New(os.Stdout, "[bcow-go/log/elasticsearch]", log.LstdFlags|log.Lmicroseconds|log.Llongfile|log.Lmsgprefix)
)

type Writer struct {
	address          string
	indexPartitioner IndexPartitioner

	pingClient  *http.Client
	queryClient *http.Client
	logger      *log.Logger
}

func NewWriter(opt *Option) *Writer {
	queryClient := &http.Client{
		Timeout: opt.QueryTimeout,
	}

	return &Writer{
		address:          opt.Address,
		indexPartitioner: opt.IndexPartitioner,
		pingClient:       &http.Client{},
		queryClient:      queryClient,
		logger:           defaultLogger,
	}
}

func (w *Writer) Write(
	logType bcowlog.EventLogType,
	eventID string,
	category string,
	source string,
	version string,
	message string,
	timestamp time.Time) {

	eventLog := &bcowlog.EventLog{
		Type:      logType,
		EventID:   eventID,
		Category:  category,
		Source:    source,
		Version:   version,
		Message:   message,
		Timestamp: timestamp.UnixNano() / int64(time.Millisecond),
	}
	w.WriteEventLog(eventLog)
}

func (w *Writer) WriteEventLog(eventLog *bcowlog.EventLog) {
	defer func() {
		err := recover()
		if err != nil {
			w.logger.Printf("ERR cannot update EventLog (#%s), %v", eventLog.EventID, err)
		}
	}()

	url := makeDocumentAPIUri(w.address, w.indexPartitioner.IndexName(), eventLog.EventID)
	body, err := json.Marshal(eventLog)
	if err != nil {
		panic(err)
	}

	ok := update(w.queryClient, url, body)
	if !ok {
		w.logger.Printf("FAIL cannot update EventLog (#%s)", eventLog.EventID)
	}
}

func (w *Writer) Ping(timeout time.Duration) bool {
	w.pingClient.Timeout = timeout
	return ping(w.pingClient, w.address)
}
