package req

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type HTTPClient struct {
	URL      string
	Username string
	Password string
	Timeout  time.Duration
	client   *http.Client
}

func NewHTTPClient(url, username, password string, timeout time.Duration) *HTTPClient {
	jar, _ := cookiejar.New(nil)
	httpClient := &HTTPClient{
		URL:      url,
		Username: username,
		Password: password,
		Timeout:  timeout,
	}
	client := &http.Client{Jar: jar, Timeout: httpClient.Timeout}
	httpClient.client = client
	return httpClient
}

func (c *HTTPClient) Get(path string) ([]byte, error) {
	resp, err := c.request("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *HTTPClient) Post(path string, body io.Reader) ([]byte, error) {
	resp, err := c.request("POST", path, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *HTTPClient) request(method, path string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", strings.TrimPrefix(c.URL, "/"), path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Content-Type", "application/json")
	return c.client.Do(req)
}
