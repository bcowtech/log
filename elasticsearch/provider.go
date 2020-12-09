package elasticsearch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func makeDocumentAPIUri(address, indexName, documentId string) string {
	if len(documentId) > 0 {
		return address + "/" + indexName + "/_doc/" + documentId
	}
	return address + "/" + indexName + "/_doc"
}

func ping(client *http.Client, url string) bool {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	// # from elasticsearch 7.4.2
	// $ curl -XHEAD -I -s -v http://192.168.56.54:9200/
	// HTTP/1.1 200 OK
	// content-type: application/json; charset=UTF-8
	// content-length: 541
	//
	// *   Trying 192.168.56.54:9200...
	// * Connected to 192.168.56.54 (192.168.56.54) port 9200 (#0)
	// > HEAD / HTTP/1.1
	// > Host: 192.168.56.54:9200
	// > User-Agent: curl/7.71.1
	// > Accept: */*
	// >
	// * Mark bundle as not supporting multiuse
	// < HTTP/1.1 200 OK
	// < content-type: application/json; charset=UTF-8
	// < content-length: 541
	// <
	return (resp.StatusCode == 200)
}

func update(client *http.Client, url string, data []byte) bool {
	req, err := http.NewRequest("PUT", url, bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// https://elasticsearch-py.readthedocs.io/en/7.9.1/exceptions.html#exceptions
	// Exception raised when ES returns a non-OK (>=400) HTTP status code.
	// Or when an actual connection error happens; in that case the status_code
	// will be set to 'N/A'.
	if resp.StatusCode >= 400 {
		return false
	}

	// # from elasticsearch 7.4.2
	// # https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-index_.html
	// $ curl -X PUT -s -v "192.168.56.54:9200/service-access-log-000001/_doc/192.168.56.54-0001?pretty" -H 'Content-Type: application/json' -d'
	// > {
	// >   "timestamp": "1560973500123",
	// >   "event_id" : "192.168.56.54#0001",
	// >   "category" : "WalletService",
	// >   "source"   : "192.168.56.54",
	// >   ...
	// > }'
	// {
	//   "_index" : "service-access-log-000001",
	//   "_type" : "_doc",
	//   "_id" : "192.168.56.54-0001",
	//   "_version" : 4,
	//   "result" : "updated",
	//   "_shards" : {
	//     "total" : 2,
	//     "successful" : 1,
	//     "failed" : 0
	//   },
	//   "_seq_no" : 3,
	//   "_primary_term" : 2
	// }
	//   Trying 192.168.56.54:9200...
	// * Connected to 192.168.56.54 (192.168.56.54) port 9200 (#0)
	// > PUT /service-access-log-000001/_doc/192.168.56.54-0001?pretty HTTP/1.1
	// > Host: 192.168.56.54:9200
	// > User-Agent: curl/7.71.1
	// > Accept: */*
	// > Content-Type: application/json
	// > Content-Length: 785
	// >
	// } [785 bytes data]
	// * upload completely sent off: 785 out of 785 bytes
	// * Mark bundle as not supporting multiuse
	// < HTTP/1.1 200 OK
	// < content-type: application/json; charset=UTF-8
	// < content-length: 256
	// <
	// { [256 bytes data]

	var result = make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	if v, ok := result["_shards"]; ok {
		if _shards, ok := v.(map[string]interface{}); ok {
			if vv, ok := _shards["successful"]; ok {
				if successful, ok := vv.(float64); ok {
					return (successful > 0)
				}
			}
		}
	}
	return false
}
