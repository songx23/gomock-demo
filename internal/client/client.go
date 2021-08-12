package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	httpClient *http.Client
	baseURL    url.URL
}

func NewClient(c *http.Client, u url.URL) *Client {
	return &Client{
		httpClient: c,
		baseURL:    u,
	}
}

func (c *Client) Magic(something string) (string, error) {
	c.baseURL.Path = path.Join(c.baseURL.Path, fmt.Sprintf("/%s", something))
	r, err := http.NewRequest(http.MethodGet, c.baseURL.String(), nil)
	if err != nil {
		return "", err
	}
	res, err := c.httpClient.Do(r)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
