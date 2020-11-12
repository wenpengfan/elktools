package esapi

import (
	"elktools/req"
	"time"
)

type Search struct {
	Req *req.HTTPClient
}

func NewSearch(url, username, password string, timeout time.Duration) *Search {
	client := req.NewHTTPClient(url, username, password, timeout)
	return &Search{Req: client}
}
