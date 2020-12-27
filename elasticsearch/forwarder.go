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

type Forwarder struct {
	address string

	pingClient  *http.Client
	queryClient *http.Client
	logger      *log.Logger
}

func NewForwarder(opt *Option) *Forwarder {
	queryClient := &http.Client{
		Timeout: opt.QueryTimeout,
	}

	return &Forwarder{
		address:     opt.Address,
		pingClient:  &http.Client{},
		queryClient: queryClient,
		logger:      defaultLogger,
	}
}

func (w *Forwarder) write(indexPartitioner IndexProvider, eventLog *bcowlog.EventLog) {
	defer func() {
		err := recover()
		if err != nil {
			w.logger.Printf("ERR cannot update EventLog (#%s), %v", eventLog.EventID, err)
		}
	}()

	url := makeDocumentAPIUri(w.address, indexPartitioner.IndexName(), eventLog.EventID)
	body, err := json.Marshal(eventLog)
	if err != nil {
		panic(err)
	}

	ok := update(w.queryClient, url, body)
	if !ok {
		w.logger.Printf("FAIL cannot update EventLog (#%s)", eventLog.EventID)
	}
}

func (w *Forwarder) Ping(timeout time.Duration) bool {
	w.pingClient.Timeout = timeout
	return ping(w.pingClient, w.address)
}
