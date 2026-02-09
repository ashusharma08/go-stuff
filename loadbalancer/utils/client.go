package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	Timeout time.Duration
	client  *http.Client
}

type Options func(*Client)

func NewClient(opts ...Options) *Client {
	c := &Client{
		client: &http.Client{},
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

func WithTimeout(t time.Duration) Options {
	return func(c *Client) {
		c.client.Timeout = t
	}
}

func (c *Client) Get(url string) ([]byte, error) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	bts, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bts, nil
}

func (c *Client) Do(host string, r *http.Request) (*http.Response, error) {
	fmt.Println("___________ host, ", host)
	u := new(url.URL)
	*u = *r.URL
	u.Scheme = r.URL.Scheme
	u.Host = host
	u.Scheme = "http"
	fmt.Println("___________ u.host, ", u.Host, "____", u.String())

	newreq, err := http.NewRequest(r.Method, u.String(), r.Body)
	if err != nil {
		fmt.Println("______________", err)
		return nil, err
	}
	return c.client.Do(newreq)
}
