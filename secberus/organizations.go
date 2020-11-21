package secberus

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Organization struct {
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

func (c *Client) GetOrganizations() (orgs *[]Organization, err error) {
	resp, err := handleResponse(c.Get("/accounts/organizations", &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&orgs)

	return
}

func (c *Client) GetOrganization(id string) (org *Organization, err error) {
	path := fmt.Sprintf("/accounts/organizations/%s", id)

	resp, err := handleResponse(c.Get(path, &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&org)

	return
}

func (c *Client) CreateOrganization(coi Organization) (org *Organization, err error) {
	bodyBytes, err := json.Marshal(coi)
	if err != nil {
		return
	}
	bodyReader := bytes.NewReader(bodyBytes)

	ro := &RequestOptions{
		Body:       bodyReader,
		BodyLength: int64(len(bodyBytes)),
	}

	resp, err := handleResponse(c.Post("/accounts/organizations", ro))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&org)

	return
}

func (c *Client) UpdateOrganization(org *Organization) (err error) {
	bodyBytes, err := json.Marshal(org)
	if err != nil {
		return
	}
	bodyReader := bytes.NewReader(bodyBytes)

	ro := &RequestOptions{
		Body:       bodyReader,
		BodyLength: int64(len(bodyBytes)),
	}

	path := fmt.Sprintf("/accounts/organizations/%s", org.Id)

	_, err = handleResponse(c.Put(path, ro))
	if err != nil {
		err = handleError(err)
		return
	}

	return
}

func (c *Client) DeleteOrganization(id string) (err error) {
	path := fmt.Sprintf("/accounts/organizations/%s", id)

	_, err = handleResponse(c.Delete(path, &RequestOptions{}))
	if err != nil {
		err = handleError(err)
	}

	return
}

func (c *Client) SetOrganizationUsers(orgid string, userid []string) (err error) {
	bodyBytes, err := json.Marshal(userid)
	if err != nil {
		return
	}
	bodyReader := bytes.NewReader(bodyBytes)

	ro := &RequestOptions{
		Body:       bodyReader,
		BodyLength: int64(len(bodyBytes)),
	}

	path := fmt.Sprintf("/organizations/%s/users", orgid)

	_, err = handleResponse(c.Put(path, ro))

	return
}

func (c *Client) GetOrganizationUsers(id string) (users *[]User, err error) {
	path := fmt.Sprintf("/organizations/%s/users", id)

	resp, err := handleResponse(c.Get(path, &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&users)

	return
}
