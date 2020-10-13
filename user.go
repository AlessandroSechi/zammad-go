package zammad

import "fmt"

func (c *Client) UserMe() (*map[string]interface{}, error) {
	var user map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/users/me"), nil)
	if err != nil {
		return &user, err
	}

	if err = c.SendWithAuth(req, &user); err != nil {
		return &user, err
	}

	return &user, nil
}

func (c *Client) UserList() (*[]map[string]interface{}, error) {
	var users []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/users"), nil)
	if err != nil {
		return &users, err
	}

	if err = c.SendWithAuth(req, &users); err != nil {
		return &users, err
	}

	return &users, nil
}

func (c *Client) UserSearch(query string, limit int) (*[]map[string]interface{}, error) {
	var users []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/search?query=%s&limit=%d", query, limit)), nil)
	if err != nil {
		return &users, err
	}

	if err = c.SendWithAuth(req, &users); err != nil {
		return &users, err
	}

	return &users, nil
}

func (c *Client) UserShow(userID int) (*map[string]interface{}, error) {
	var user map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/%d", userID)), nil)
	if err != nil {
		return &user, err
	}

	if err = c.SendWithAuth(req, &user); err != nil {
		return &user, err
	}

	return &user, nil
}

func (c *Client) UserCreate(u *map[string]interface{}) (*map[string]interface{}, error) {
	var user map[string]interface{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/users"), u)
	if err != nil {
		return &user, err
	}

	if err = c.SendWithAuth(req, &user); err != nil {
		return &user, err
	}

	return &user, nil
}

func (c *Client) UserUpdate(userID int, u *map[string]interface{}) (*map[string]interface{}, error) {
	var user map[string]interface{}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/users/%d", userID)), u)
	if err != nil {
		return &user, err
	}

	if err = c.SendWithAuth(req, &user); err != nil {
		return &user, err
	}

	return &user, nil
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