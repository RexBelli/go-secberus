package secberus

import (
	"io"
	"net/http"
)

type RequestOptions struct {
	Params     map[string]string
	Headers    map[string]string
	Body       io.Reader
	BodyLength int64
}

func (c *Client) Request(verb, path string, ro *RequestOptions) (*http.Response, error) {
	return nil, nil
}
