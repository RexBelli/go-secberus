package secberus

import (
	"encoding/json"
)

type Resource struct {
	ID           string      `json:"id"`
	ResourceID   string      `json:"resource_id,omitempty"`
	Name         string      `json:"name"`
	DataProvider string      `json:"dp"`
	Description  string      `json:"description,omitempty"`
	Score        int         `json:"score"`
	ExampleData  interface{} `json:"example_data,omitempty"`
	Required     bool        `json:"required"`
}

func (c *Client) GetResources() (resources *[]Resource, err error) {
	resp, err := handleResponse(c.Get("/resources", &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&resources)

	return
}
