package test

import (
	"bytes"
	"testing"

	log "gitlab.bcowtech.de/bcow-go/log"
)

func TestJsonTextWriter(t *testing.T) {
	var buffer bytes.Buffer

	logger := log.NewLogger(&log.Config{
		Category: "demo-api-server",
		Source:   "192.168.56.51",
		Version:  "v1.0.1",
		Writer: &log.JsonTextWriter{
			Indent: "    ",
			Stream: log.MixStreams(&buffer),
		},
	})

	buffer.Reset()
	if len(buffer.String()) != 0 {
		t.Errorf("buffer should be empty")
	}
	logger.Write(log.PASS, "log message successful")
	/*
	 * {
	 *   "type": "PASS",
	 *   "event_id": "27c17c13-17f4-40d9-b76e-578c38918d42",
	 *   "category": "demo-api-server",
	 *   "source": "192.168.56.51",
	 *   "version": "v1.0.1",
	 *   "message": "log message successful",
	 *   "timestamp": 1604579769853
	 * }
	 */
	if len(buffer.String()) == 0 {
		t.Errorf("message should not be empty")
	}

	buffer.Reset()
	if len(buffer.String()) != 0 {
		t.Errorf("buffer should be empty")
	}
	logger.Write(log.DEBUG, "log debug message")
	/*
	 * {
	 *  "type": "DEBUG",
	 *  "event_id": "9414d5d8-b9e1-45c1-8251-aa4f82868374",
	 *  "category": "demo-api-server",
	 *  "source": "192.168.56.51",
	 *  "version": "v1.0.1",
	 *  "message": "log debug message",
	 *  "timestamp": 1604579769863
	 *}
	 */
	if len(buffer.String()) == 0 {
		t.Errorf("message should not be empty")
	}
}

func TestPlainTextWriter(t *testing.T) {
	var buffer bytes.Buffer

	logger := log.NewLogger(&log.Config{
		Category: "demo-api-server",
		Source:   "192.168.56.51",
		Version:  "v1.0.1",
		Writer: &log.PlainTextWriter{
			Stream: &buffer,
		},
	})

	buffer.Reset()
	if len(buffer.String()) != 0 {
		t.Errorf("buffer should be empty")
	}
	logger.Write(log.PASS, "log message successful")
	/*
	 * 2020-11-05T12:46:16.780Z PASS [demo-api-server/192.168.56.51 - v1.0.1] log message successful (#e79dd213-3778-434c-86fc-c645c6f5935b)
	 */
	if len(buffer.String()) == 0 {
		t.Errorf("message should not be empty")
	}

	buffer.Reset()
	if len(buffer.String()) != 0 {
		t.Errorf("buffer should be empty")
	}
	logger.Write(log.DEBUG, "log debug message")
	/*
	 * 2020-11-05T12:46:16.780Z DEBUG [demo-api-server/192.168.56.51 - v1.0.1] log debug message (#978d9988-a225-4e87-943d-5f13d555494f)
	 */
	if len(buffer.String()) == 0 {
		t.Errorf("message should not be empty")
	}
}

func TestMixStreams(t *testing.T) {
	var buffer1 bytes.Buffer
	var buffer2 bytes.Buffer

	logger := log.NewLogger(&log.Config{
		Category: "demo-api-server",
		Source:   "192.168.56.51",
		Version:  "v1.0.1",
		Writer: &log.PlainTextWriter{
			Stream: log.MixStreams(
				&buffer1,
				&buffer2,
			),
		},
	})

	buffer1.Reset()
	if len(buffer1.String()) != 0 {
		t.Errorf("buffer1 should be empty")
	}
	buffer2.Reset()
	if len(buffer2.String()) != 0 {
		t.Errorf("buffer2 should be empty")
	}
	logger.Write(log.PASS, "log message successful")
	/*
	 * 2020-11-05T12:46:16.780Z PASS [demo-api-server/192.168.56.51 - v1.0.1] log message successful (#e79dd213-3778-434c-86fc-c645c6f5935b)
	 */
	if len(buffer1.String()) == 0 {
		t.Errorf("message should not be empty")
	}
	if len(buffer2.String()) == 0 {
		t.Errorf("message should not be empty")
	}

	buffer1.Reset()
	if len(buffer1.String()) != 0 {
		t.Errorf("buffer1 should be empty")
	}
	buffer2.Reset()
	if len(buffer2.String()) != 0 {
		t.Errorf("buffer2 should be empty")
	}
	logger.Write(log.DEBUG, "log debug message")
	/*
	 * 2020-11-05T12:46:16.780Z DEBUG [demo-api-server/192.168.56.51 - v1.0.1] log debug message (#978d9988-a225-4e87-943d-5f13d555494f)
	 */
	if len(buffer1.String()) == 0 {
		t.Errorf("message should not be empty")
	}
	if len(buffer2.String()) == 0 {
		t.Errorf("message should not be empty")
	}
}

func TestUseFilter(t *testing.T) {
	var buffer bytes.Buffer

	logger := log.NewLogger(&log.Config{
		Category: "demo-api-server",
		Source:   "192.168.56.51",
		Version:  "v1.0.1",
		Writer: &log.FilterableTextWriter{
			Writer: &log.PlainTextWriter{
				Stream: &buffer,
			},
			Filter: &log.SeverityFilter{
				Severity: log.ERR.Severity(),
			},
		},
	})

	buffer.Reset()
	if len(buffer.String()) != 0 {
		t.Errorf("buffer should be empty")
	}
	logger.Write(log.PASS, "log message successful")
	/*
	* 2020-11-05T13:07:01.571Z PASS [demo-api-server/192.168.56.51 - v1.0.1] log message successful (#60cf7b3b-a09c-43fe-aec0-e8ceb48750ba)
	 */
	if len(buffer.String()) == 0 {
		t.Errorf("message should not be empty")
	}

	buffer.Reset()
	if len(buffer.String()) != 0 {
		t.Errorf("buffer should be empty")
	}
	logger.Write(log.DEBUG, "log debug message")
	if len(buffer.String()) != 0 {
		t.Errorf("message should be empty")
	}

	buffer.Reset()
	if len(buffer.String()) != 0 {
		t.Errorf("buffer should be empty")
	}
	logger.Write(log.ERR, "log err message")
	/*
	 * 2020-11-05T13:07:01.571Z ERR [demo-api-server/192.168.56.51 - v1.0.1] log err message (#89bb871c-60e6-43c2-8399-af6c48a4770c)
	 */
	if len(buffer.String()) == 0 {
		t.Errorf("message should not be empty")
	}
}

func TestMixWriters(t *testing.T) {
	var buffer1, buffer2 bytes.Buffer

	logger := log.NewLogger(&log.Config{
		Category: "demo-api-server",
		Source:   "192.168.56.51",
		Version:  "v1.0.1",
		Writer: log.MixWriters(
			&log.JsonTextWriter{
				Indent: "    ",
				Stream: &buffer1,
			},
			&log.FilterableTextWriter{
				Writer: &log.PlainTextWriter{
					Stream: &buffer2,
				},
				Filter: &log.SeverityFilter{
					Severity: log.ERR.Severity(),
				},
			},
		),
	})

	buffer1.Reset()
	if len(buffer1.String()) != 0 {
		t.Errorf("buffer1 should be empty")
	}
	buffer2.Reset()
	if len(buffer2.String()) != 0 {
		t.Errorf("buffer2 should be empty")
	}
	logger.Write(log.PASS, "log message successful")
	/*
	 * {
	 *   "type": "PASS",
	 *   "event_id": "204dae8f-578e-4dcf-b26c-822e9388e18d",
	 *   "category": "demo-api-server",
	 *   "source": "192.168.56.51",
	 *   "version": "v1.0.1",
	 *   "message": "log message successful",
	 *   "timestamp": 1604581920443
	 * }
	 */
	if len(buffer1.String()) == 0 {
		t.Errorf("message should not be empty")
	}
	/*
	 * 2020-11-05T13:12:00.443Z PASS [demo-api-server/192.168.56.51 - v1.0.1] log message successful (#204dae8f-578e-4dcf-b26c-822e9388e18d)
	 */
	if len(buffer2.String()) == 0 {
		t.Errorf("message should not be empty")
	}

	buffer1.Reset()
	if len(buffer1.String()) != 0 {
		t.Errorf("buffer1 should be empty")
	}
	buffer2.Reset()
	if len(buffer2.String()) != 0 {
		t.Errorf("buffer2 should be empty")
	}
	logger.Write(log.DEBUG, "log debug message")
	/*
	 * {
	 *   "type": "DEBUG",
	 *   "event_id": "511c650b-bf04-4da3-9f96-c122ed5a1ce9",
	 *   "category": "demo-api-server",
	 *   "source": "192.168.56.51",
	 *   "version": "v1.0.1",
	 *   "message": "log debug message",
	 *   "timestamp": 1604581920453
	 * }
	 */
	if len(buffer1.String()) == 0 {
		t.Errorf("message should not be empty")
	}
	if len(buffer2.String()) != 0 {
		t.Errorf("message should be empty")
	}

	buffer1.Reset()
	if len(buffer1.String()) != 0 {
		t.Errorf("buffer1 should be empty")
	}
	buffer2.Reset()
	if len(buffer2.String()) != 0 {
		t.Errorf("buffer2 should be empty")
	}
	logger.Write(log.ERR, "log err message")
	/*
	 * {
	 *   "type": "ERR",
	 *   "event_id": "68dd82bf-4542-40c1-a892-0bd1cb217702",
	 *   "category": "demo-api-server",
	 *   "source": "192.168.56.51",
	 *   "version": "v1.0.1",
	 *   "message": "log err message",
	 *   "timestamp": 1604581920456
	 * }
	 */
	if len(buffer1.String()) == 0 {
		t.Errorf("message should not be empty")
	}
	/*
	 * 2020-11-05T13:12:00.456Z ERR [demo-api-server/192.168.56.51 - v1.0.1] log err message (#68dd82bf-4542-40c1-a892-0bd1cb217702)
	 */
	if len(buffer2.String()) == 0 {
		t.Errorf("message should not be empty")
	}
}

func TestEventLog(t *testing.T) {

	var buffer1, buffer2 bytes.Buffer

	writer := log.MixWriters(
		&log.JsonTextWriter{
			Indent: "  ",
			Stream: &buffer1,
		},
		&log.FilterableTextWriter{
			Writer: &log.PlainTextWriter{
				Stream: &buffer2,
			},
			Filter: &log.SeverityFilter{
				Severity: log.ERR.Severity(),
			},
		},
	)

	event := log.EventLog{
		Type:     log.PASS,
		EventID:  "68dd82bf-4542-40c1-a892-0bd1cb217702",
		Category: "demo-api-server",
		Source:   "192.168.56.51",
		Version:  "v1.0.1",
		Message:  "log message successful",
		Details: log.Detail{
			Request: log.HttpRequestDetail{
				Method:      "GET",
				Path:        "/echo",
				QueryString: "foo=bar&baz=1",
				Header:      "Content-Type: application/json\r\nUser-Agent: Debian APT-HTTP/1.3 (0.9.7.9)",
				Body:        "some request body",
			},
			Response: log.HttpResponseDetail{
				StatusCode: "200",
				Header:     "X-Response-Header: f/twXyy",
				Body:       "some response body",
			},
			Error: "no errors",
		},
		Metric: log.Metric{
			ElapsedTime:       6,
			Depth:             1,
			ResponseBodyBytes: 18,
		},
	}

	buffer1.Reset()
	if len(buffer1.String()) != 0 {
		t.Errorf("buffer1 should be empty")
	}
	buffer2.Reset()
	if len(buffer2.String()) != 0 {
		t.Errorf("buffer2 should be empty")
	}
	event.WriteTo(writer)
	/*
	 * {
	 *   "type": "PASS",
	 *   "severity": 0,
	 *   "event_id": "68dd82bf-4542-40c1-a892-0bd1cb217702",
	 *   "category": "demo-api-server",
	 *   "Source": "192.168.56.51",
	 *   "version": "v1.0.1",
	 *   "message": "log message successful",
	 *   "details": {
	 *     "request": {
	 *       "method": "GET",
	 *       "path": "/echo",
	 *       "query_string": "foo=bar\u0026baz=1",
	 *       "header": "Content-Type: application/json\r\nUser-Agent: Debian APT-HTTP/1.3 (0.9.7.9)",
	 *       "body": "some request body"
	 *     },
	 *     "response": {
	 *       "status_code": "200",
	 *       "header": "X-Response-Header: f/twXyy",
	 *       "body": "some response body"
	 *     },
	 *     "error": "no errors"
	 *   },
	 *   "metric": {
	 *     "elapsed_time": 6,
	 *     "depth": 1,
	 *     "response_body_bytes": 18
	 *   },
	 *   "timestamp": 0
	 * }
	 */
	if len(buffer1.String()) == 0 {
		t.Errorf("message should not be empty")
	}
	/*
	 *1970-01-01T00:00:00.000Z PASS [demo-api-server/192.168.56.51 - v1.0.1] log message successful (#68dd82bf-4542-40c1-a892-0bd1cb217702)
	 *
	 *Request:
	 *GET /echo?foo=bar&baz=1
	 *Content-Type: application/json
	 *User-Agent: Debian APT-HTTP/1.3 (0.9.7.9)
	 *some request body
	 *
	 *Response:
	 *200
	 *X-Response-Header: f/twXyy
	 *some response body
	 *
	 *Error:
	 *no errors
	 *
	 *Metric:
	 *elapsed time: 6
	 *depth: 1
	 *response body bytes: 18
	 *
	 */
	if len(buffer2.String()) == 0 {
		t.Errorf("message should not be empty")
	}
}
