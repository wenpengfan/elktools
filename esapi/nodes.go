package esapi

import (
	"fmt"
	"strings"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/5.6/cat-nodes.html
func (es *Search) GetNodes(sort string, args []string) ([]byte, error) {
	reqArgs := []string{"?v", fmt.Sprintf("s=name:%s", sort)}
	reqArgs = append(reqArgs, args...)
	path := fmt.Sprintf("_cat/nodes%s", strings.Join(reqArgs, "&"))
	return es.get(path)
}
