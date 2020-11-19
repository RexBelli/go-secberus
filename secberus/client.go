package secberus

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
)

const APIKey = "STUB_KEY"

const DefaultEndpoint = "api2.seberus.com"

var ProjectURL = "github.com/RexBelli/go-secberus"

var ProjectVersion = "0.0.1"

var UserAgent = fmt.Sprintf("go-secberus/%s (+%s; %s)",
	ProjectVersion, ProjectURL, runtime.Version())

type Client struct {
	Address    string
	HTTPClient *http.Client
	apiKey     string
	url        *url.URL
}

// NewClient creates a Secberus client
func NewClient(key string) (*Client, error) {
	return (&Client{
		apiKey:  key,
		Address: DefaultEndpoint,
	}).init()
}

func (c *Client) init() (*Client, error) {
	u, err := url.Parse(c.Address)
	if err != nil {
		return nil, err
	}
	c.url = u

	return c, nil
}
