package esapi

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// https://www.elastic.co/guide/en/elasticsearch/reference/5.6/cat-indices.html
func (es *Search) GetIndex(sort string, args []string) ([]byte, error) {
	return es.getIndex("0", sort, args)
}
func (es *Search) GetIndexDay(day, sort string, args []string) ([]byte, error) {
	return es.getIndex(day, sort, args)
}

func (es *Search) GetIndexAll(sort string, args []string) ([]byte, error) {
	var reqArgs = []string{"?v", fmt.Sprintf("s=store.size:%s", sort)}
	reqArgs = append(reqArgs, args...)
	path := fmt.Sprintf("_cat/indices%s", strings.Join(reqArgs, "&"))
	return es.get(path)
}

func (es *Search) getIndex(day, sort string, args []string) ([]byte, error) {
	var (
		layouts = []string{
			"2006.1.2",
			"2006.01.02",
			"2006/1/2",
			"2006/01/02",
			"2006-1-2",
			"2006-01-02",
		}
		searchLayout = "2006.01.02"
		searchDay    = time.Now().Format(searchLayout)
		reqArgs      = []string{"?v", fmt.Sprintf("s=store.size:%s", sort)}
	)
	reqArgs = append(reqArgs, args...)
	n, err := strconv.Atoi(day)
	if err == nil {
		searchDay = time.Now().AddDate(0, 0, n).Format(searchLayout)
	} else {
		for _, layout := range layouts {
			t, err := time.Parse(layout, day)
			if err == nil {
				searchDay = t.Format(searchLayout)
			}
		}
	}

	path := fmt.Sprintf("_cat/indices/*-%s%s", searchDay, strings.Join(reqArgs, "&"))
	return es.get(path)
}
