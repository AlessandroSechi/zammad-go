package zammad

import (
	"fmt"
	"net/url"
	"time"
)

// User is a Zammad user. See https://docs.zammad.org/en/latest/api/user.html.
type User struct {
	ID             int       `json:"id"`
	OrganizationID int       `json:"organization_id"`
	Login          string    `json:"login"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	Email          string    `json:"email"`
	Web            string    `json:"web"`
	LastLogin      time.Time `json:"last_login"`
}

// UserMe returns the current authenticated user.
func (c *Client) UserMe() (User, error) {
	var user User

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/users/me"), nil)
	if err != nil {
		return user, err
	}

	if err = c.SendWithAuth(req, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) UserList() ([]User, error) {
	var users []User

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/users"), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) UserSearch(query string, limit int) ([]User, error) {
	var users []User

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/search?query=%s&limit=%d", url.QueryEscape(query), limit)), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) UserShow(userID int) (User, error) {
	var user User

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/%d", userID)), nil)
	if err != nil {
		return user, err
	}

	if err = c.SendWithAuth(req, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) UserCreate(u User) (User, error) {
	var user User

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/users"), u)
	if err != nil {
		return user, err
	}

	if err = c.SendWithAuth(req, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) UserUpdate(userID int, u User) (User, error) {
	var user User

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/%d", userID)), u)
	if err != nil {
		return user, err
	}

	if err = c.SendWithAuth(req, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) UserDelete(userID int) error {

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/%d", userID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
