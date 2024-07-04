package zammad

import "fmt"

func (c *Client) TagList(ticketID int) (*[]map[string]interface{}, error) {
	var tags []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tags?object=Ticket&o_id=%d", ticketID)), nil)
	if err != nil {
		return &tags, err
	}

	if err = c.SendWithAuth(req, &tags); err != nil {
		return &tags, err
	}

	return &tags, nil
}

func (c *Client) TagSearch(term string) (*[]map[string]interface{}, error) {
	var tags []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_search?term=%s", term)), nil)
	if err != nil {
		return &tags, err
	}

	if err = c.SendWithAuth(req, &tags); err != nil {
		return &tags, err
	}

	return &tags, nil
}

func (c *Client) TagAdd(t *map[string]interface{}) error {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/tags/add"), t)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagRemove(t *map[string]interface{}) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, "/api/v1/tags/remove"), t)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminList() (*[]map[string]interface{}, error) {
	var tags []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/tag_list"), nil)
	if err != nil {
		return &tags, err
	}

	if err = c.SendWithAuth(req, &tags); err != nil {
		return &tags, err
	}

	return &tags, nil
}

func (c *Client) TagAdminCreate(o *map[string]interface{}) error {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/tag_list"), o)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminRename(tagID int, t *map[string]interface{}) error {

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_list/%d", tagID)), t)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminDelete(tagID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_list/%d", tagID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
