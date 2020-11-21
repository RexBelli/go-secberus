package secberus

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
)

const DefaultEnvAPIKey = "SECBERUS_API_KEY"
const DefaultEnvUsername = "SECBERUS_USERNAME"
const DefaultEnvPassword = "SECBERUS_PASSWORD"

const DefaultEndpoint = "https://api2.secberus.com"

const DebugProxy = "http://127.0.0.1:8081"

var ProjectURL = "github.com/RexBelli/go-secberus"

var ProjectVersion = "0.0.1"

var UserAgent = fmt.Sprintf("go-secberus/%s (+%s; %s)",
	ProjectVersion, ProjectURL, runtime.Version())

type Client struct {
	Address      string
	HTTPClient   *http.Client
	username     string
	password     string
	accessToken  string
	refreshToken string
	url          *url.URL
}

// NewClientWithAPIKey creates a Secberus client using Env variables
func NewClientFromEnv() (*Client, error) {
	if apiKey := os.Getenv(DefaultEnvAPIKey); apiKey != "" {
		return NewClientWithAPIKey(apiKey)
	}

	username := os.Getenv(DefaultEnvUsername)
	password := os.Getenv(DefaultEnvPassword)
	if username != "" && password != "" {
		return NewClientWithCredentials(username, password)
	}

	return nil, errors.New("no apikey or username/password found in env")
}

// NewClientWithAPIKey creates a Secberus client using an API key
func NewClientWithAPIKey(key string) (*Client, error) {
	if key == "" {
		return nil, errors.New("no API key provided")
	}

	c := &Client{
		accessToken: key,
	}
	return c.init()
}

func NewClientWithCredentials(username, password string) (*Client, error) {
	if username == "" || password == "" {
		return nil, errors.New("missing username or password")
	}

	c := &Client{
		username: username,
		password: password,
	}
	c.init()
	return c.login()
}

func (c *Client) init() (*Client, error) {
	c.Address = DefaultEndpoint

	u, err := url.Parse(c.Address)
	if err != nil {
		return nil, err
	}
	c.url = u

	c.HTTPClient = &http.Client{}

	if os.Getenv("DEBUG") != "" {
		os.Setenv("HTTP_PROXY", DebugProxy)
		os.Setenv("HTTPS_PROXY", DebugProxy)
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return c, nil
}

func (c *Client) login() (*Client, error) {
	type loginIn struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		MfaRequired bool   `json:"mfaRequired"`
	}
	type loginOut struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	in := &loginIn{
		Username:    c.username,
		Password:    c.password,
		MfaRequired: false,
	}

	bodyBytes, err := json.Marshal(in)
	if err != nil {
		return c, err
	}
	bodyReader := bytes.NewReader(bodyBytes)

	ro := &RequestOptions{
		Body:       bodyReader,
		BodyLength: int64(len(bodyBytes)),
	}

	resp, err := handleResponse(c.Post("/auth/login", ro))
	if err != nil {
		return c, handleError(err)
	}

	var out *loginOut
	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		return c, handleError(err)
	}

	c.accessToken = out.AccessToken
	c.refreshToken = out.RefreshToken

	return c, nil
}
