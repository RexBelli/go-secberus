package secberus

import "encoding/json"

type Compliance struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	URL          interface{} `json:"url"`
	ID           string      `json:"id"`
	Requirements []struct {
		Ordinal     string `json:"ordinal"`
		Identifier  string `json:"identifier"`
		Description string `json:"description"`
		ID          string `json:"id"`
		Controls    []struct {
			Ordinal     string `json:"ordinal"`
			Identifier  string `json:"identifier"`
			Description string `json:"description"`
			ID          string `json:"id"`
			Children    []struct {
				Ordinal     string        `json:"ordinal"`
				Identifier  string        `json:"identifier"`
				Description string        `json:"description"`
				ID          string        `json:"id"`
				Children    []interface{} `json:"children"`
				ParentID    string        `json:"parent_id"`
			} `json:"children"`
			ParentID interface{} `json:"parent_id"`
		} `json:"controls"`
		ComplianceID string `json:"compliance_id"`
	} `json:"requirements"`
}

func (c *Client) GetCompliances() (compliances *[]Compliance, err error) {
	resp, err := handleResponse(c.Get("/compliances", &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&compliances)

	return
}
