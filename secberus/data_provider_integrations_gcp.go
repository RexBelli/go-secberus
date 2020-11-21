package secberus

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type DataProviderIntegration struct {
	ID       string `json:"id,omitempty"`
	OrgID    string `json:"org_id"`
	Dp       string `json:"dp"`
	Name     string `json:"name"`
	Verified bool   `json:"verified"`
}

type CreateDPIGCPIn struct {
	Name string `json:"name"`
	Data struct {
		Creds struct {
			ClientX509CertURL       string `json:"client_x509_cert_url"`
			AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
			TokenURI                string `json:"token_uri"`
			AuthURI                 string `json:"auth_uri"`
			ClientID                string `json:"client_id"`
			ClientEmail             string `json:"client_email"`
			PrivateKey              string `json:"private_key"`
			PrivateKeyID            string `json:"private_key_id"`
			ProjectID               string `json:"project_id"`
			Type                    string `json:"type"`
		} `json:"creds"`
		Projects []string `json:"projects"`
		Zones    []string `json:"zones"`
	} `json:"data"`
}

func (c *Client) CreateDPIGCP(cdpiin *CreateDPIGCPIn) (dpi *DataProviderIntegration, err error) {
	bodyBytes, err := json.Marshal(cdpiin)
	if err != nil {
		return
	}
	bodyReader := bytes.NewReader(bodyBytes)

	ro := &RequestOptions{
		Body:       bodyReader,
		BodyLength: int64(len(bodyBytes)),
	}

	resp, err := handleResponse(c.Post("/dpi/gcp", ro))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&dpi)

	return
}

func (c *Client) DeleteDPI(id string) (err error) {
	path := fmt.Sprintf("/dpi/%s", id)

	_, err = handleResponse(c.Delete(path, &RequestOptions{}))
	if err != nil {
		err = handleError(err)
	}

	return
}
