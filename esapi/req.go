package esapi

import (
	"elktools/cmd/utils"
	"fmt"
	"io"
)

func (es *Search) get(path string) ([]byte, error) {
	if utils.GetDebug() {
		fmt.Printf("[DEBUG] curl -XGET -u%s:%s %s/%s , timeout: %s\n", es.Req.Username, es.Req.Password, es.Req.URL, path, es.Req.Timeout)
	}
	return es.Req.Get(path)
}

func (es *Search) post(path string, body io.Reader) ([]byte, error) {
	if utils.GetDebug() {

		fmt.Printf("[DEBUG] curl -XPOST -u%s:%s %s/%s , timeout: %s\n", es.Req.Username, es.Req.Password, es.Req.URL, path, es.Req.Timeout)
	}
	return es.Req.Post(path, body)
}
