package esapi

import (
	"fmt"
	"strings"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/5.6/cat-shards.html
func (es *Search) GetShards(sort string, args []string) ([]byte, error) {
	reqArgs := []string{"?v", fmt.Sprintf("s=state:%s", sort)}
	reqArgs = append(reqArgs, args...)
	path := fmt.Sprintf("_cat/shards%s", strings.Join(reqArgs, "&"))

	return es.get(path)
}
