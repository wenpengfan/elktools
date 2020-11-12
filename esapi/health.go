package esapi

import (
	"fmt"
	"strings"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/5.6/cat-health.html
func (es *Search) GetHealth(args []string) ([]byte, error) {
	var reqArgs = []string{"?v"}
	reqArgs = append(reqArgs, args...)
	path := fmt.Sprintf("_cat/health%s", strings.Join(reqArgs, "&"))
	return es.get(path)
}
