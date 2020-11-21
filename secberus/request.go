package secberus

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const AuthHeader = "Authorization"

const AuthHeaderFormat = "Bearer %s"

type RequestOptions struct {
	//	Params     map[string]string
	Headers    map[string]string
	Body       io.Reader
	BodyLength int64
}

func (c *Client) Request(verb, path string, ro *RequestOptions) (*http.Response, error) {
	url := strings.TrimRight(c.url.String(), "/") + "/" + strings.TrimLeft(path, "/")

	req, err := http.NewRequest(verb, url, ro.Body)
	if err != nil {
		return nil, err
	}

	if c.accessToken != "" {
		req.Header.Set(AuthHeader, fmt.Sprintf(AuthHeaderFormat, c.accessToken))
	}

	req.Header.Set("User-Agent", UserAgent)

	for k, v := range ro.Headers {
		req.Header.Add(k, v)
	}

	return c.HTTPClient.Do(req)
}

func (c *Client) Get(path string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("GET", path, ro)
}

func (c *Client) Post(path string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("POST", path, ro)
}

func (c *Client) Put(path string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("PUT", path, ro)
}

func (c *Client) Delete(path string, ro *RequestOptions) (*http.Response, error) {
	return c.Request("DELETE", path, ro)
}

func handleResponse(respIn *http.Response, errIn error) (resp *http.Response, err error) {
	resp = respIn
	err = errIn

	if err != nil {
		return
	}

	switch resp.StatusCode {
	case 200, 201, 202, 204, 205, 206:
		err = nil
	default:
		err = fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	return
}
