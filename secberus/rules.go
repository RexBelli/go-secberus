package secberus

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Rule struct {
	Description      string       `json:"description"`
	Summary          string       `json:"summary"`
	Logic            string       `json:"logic"`
	RemediationSteps string       `json:"remediation_steps"`
	AlertSummaryTmpl string       `json:"alert_summary_tmpl"`
	PolicyID         string       `json:"policy_id"`
	Priority         float64      `json:"priority"`
	ID               string       `json:"id,omitempty"`
	OrgID            string       `json:"org_id"`
	SecberusManaged  bool         `json:"secberus_managed"`
	Resources        []Resource   `json:"resources"`
	Compliances      []Compliance `json:"compliances"`
	AlertCount       int          `json:"alert_count"`
	Score            float64      `json:"score"`
	Subscribed       bool         `json:"subscribed"`
}

func (c *Client) CreateRule(ruleIn *Rule) (rule *Rule, err error) {
	bodyBytes, err := json.Marshal(ruleIn)
	if err != nil {
		return
	}
	bodyReader := bytes.NewReader(bodyBytes)

	ro := &RequestOptions{
		Body:       bodyReader,
		BodyLength: int64(len(bodyBytes)),
	}

	resp, err := handleResponse(c.Post("/rules", ro))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rule)

	return
}

func (c *Client) SetRule(id string, ruleIn *Rule) (rule *Rule, err error) {
	bodyBytes, err := json.Marshal(ruleIn)
	if err != nil {
		return
	}
	bodyReader := bytes.NewReader(bodyBytes)

	ro := &RequestOptions{
		Body:       bodyReader,
		BodyLength: int64(len(bodyBytes)),
	}

	path := fmt.Sprintf("/rules/%s", id)

	resp, err := handleResponse(c.Put(path, ro))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rule)

	return
}

func (c *Client) DeleteRule(id string) (err error) {
	path := fmt.Sprintf("/rules/%s", id)

	_, err = handleResponse(c.Delete(path, &RequestOptions{}))
	if err != nil {
		err = handleError(err)
	}

	return
}

func (c *Client) GetRule(id string) (rule *Rule, err error) {
	path := fmt.Sprintf("/rules/%s", id)

	resp, err := handleResponse(c.Get(path, &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rule)

	return
}
