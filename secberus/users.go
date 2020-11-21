package secberus

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Id                         string    `json:"id"`
	Email                      string    `json:"email"`
	FullName                   string    `json:"full_name"`
	AccountOwner               bool      `json:"account_owner"`
	Mfa                        bool      `json:"mfa"`
	Roles                      []string  `json:"roles"`
	Orgs                       []string  `json:"orgs"`
	CreateDatetime             time.Time `json:"create_datetime"`
	DeactivateDatetime         time.Time `json:"deactivate_datetime"`
	LastLoginDatetime          time.Time `json:"last_login_datetime"`
	LastPasswordChangeDatetime time.Time `json:"last_password_change_datetime"`
}

func (c *Client) GetUsers() (users *[]User, err error) {
	resp, err := handleResponse(c.Get("/users", &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&users)

	return
}

func (c *Client) UserEmailToId(email string) (id string, err error) {
	users, err := c.GetUsers()
	if err != nil {
		err = handleError(err)
		return
	}

	for _, user := range *users {
		if user.Email == email {
			id = user.Id
			return
		}
	}

	err = fmt.Errorf("no user with email %s found", email)
	return
}

func (c *Client) GetCurrentUser() (user *User, err error) {
	resp, err := handleResponse(c.Get("/user/self", &RequestOptions{}))
	if err != nil {
		err = handleError(err)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&user)

	return
}
