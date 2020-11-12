package esapi

import (
	"fmt"
	"strings"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/5.6/cluster-reroute.html
func (es *Search) RetryFailedRoute(args []string) ([]byte, error) {
	reqArgs := []string{"?retry_failed"}
	reqArgs = append(reqArgs, args...)

	path := fmt.Sprintf("_cluster/reroute%s", strings.Join(reqArgs, "&"))
	return es.post(path, nil)
}
