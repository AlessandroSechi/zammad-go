package zammad

import "fmt"

func (c *Client) GroupList() (*[]map[string]interface{}, error) {
	var groups []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/groups"), nil)
	if err != nil {
		return &groups, err
	}

	if err = c.SendWithAuth(req, &groups); err != nil {
		return &groups, err
	}

	return &groups, nil
}

func (c *Client) GroupShow(groupID int) (*map[string]interface{}, error) {
	var group map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/groups/%d", groupID)), nil)
	if err != nil {
		return &group, err
	}

	if err = c.SendWithAuth(req, &group); err != nil {
		return &group, err
	}

	return &group, nil
}

func (c *Client) GroupCreate(g *map[string]interface{}) (*map[string]interface{}, error) {
	var group map[string]interface{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/groups"), g)
	if err != nil {
		return &group, err
	}

	if err = c.SendWithAuth(req, &group); err != nil {
		return &group, err
	}

	return &group, nil
}

func (c *Client) GroupUpdate(groupID int, g *map[string]interface{}) (*map[string]interface{}, error) {
	var group map[string]interface{}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/groups/%d", groupID)), g)
	if err != nil {
		return &group, err
	}

	if err = c.SendWithAuth(req, &group); err != nil {
		return &group, err
	}

	return &group, nil
}

func (c *Client) GroupDelete(groupID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/groups/%d", groupID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
