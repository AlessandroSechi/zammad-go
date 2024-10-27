package zammad

import (
	"fmt"
	"net/http"
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

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, "/api/v1/users/me"), nil)
	if err != nil {
		return user, err
	}

	if err = c.sendWithAuth(req, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) UserListResult(opts ...Option) *Result[User] {
	return &Result[User]{
		res:     nil,
		resFunc: c.UserListWithOptions,
		opts:    NewRequestOptions(opts...),
	}
}

func (c *Client) UserList() ([]User, error) {
	return c.UserListResult().FetchAll()
}

func (c *Client) UserListWithOptions(ro RequestOptions) ([]User, error) {
	var users []User

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, "/api/v1/users"), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = ro.URLParams()

	if err = c.sendWithAuth(req, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) UserSearch(query string, limit int) ([]User, error) {
	var users []User

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/search?query=%s&limit=%d", url.QueryEscape(query), limit)), nil)
	if err != nil {
		return nil, err
	}

	if err = c.sendWithAuth(req, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (c *Client) UserShow(userID int) (User, error) {
	var user User

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/%d", userID)), nil)
	if err != nil {
		return user, err
	}

	if err = c.sendWithAuth(req, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) UserCreate(u User) (User, error) {
	var user User

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/users"), u)
	if err != nil {
		return user, err
	}

	if err = c.sendWithAuth(req, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) UserUpdate(userID int, u User) (User, error) {
	var user User

	req, err := c.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/%d", userID)), u)
	if err != nil {
		return user, err
	}

	if err = c.sendWithAuth(req, &user); err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) UserDelete(userID int) error {

	req, err := c.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/%d", userID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
