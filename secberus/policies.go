package secberus

import "encoding/json"

type Policy struct {
	ID              string        `json:"id"`
	OrgID           interface{}   `json:"org_id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	SecberusManaged bool          `json:"secberus_managed"`
	Rules           []interface{} `json:"rules"`
	Status          interface{}   `json:"status"`
	Schedule        interface{}   `json:"schedule"`
	RiskTrend       struct {
		TimeSpan    float64 `json:"time_span"`
		TimeMeasure string  `json:"time_measure"`
		Data        []struct {
			Timestamp int     `json:"timestamp"`
			Value     float64 `json:"value"`
		} `json:"data"`
	} `json:"risk_trend"`
}

func (c *Client) GetPolicies() (policies *[]Policy, err error) {
	resp, err := handleResponse(c.Get("/policies", &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&policies)

	return
}
