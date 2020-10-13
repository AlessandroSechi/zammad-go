package zammad

import "fmt"

func (c *Client) OrganizationList() (*[]map[string]interface{}, error) {
	var organizations []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/organizations"), nil)
	if err != nil {
		return &organizations, err
	}

	if err = c.SendWithAuth(req, &organizations); err != nil {
		return &organizations, err
	}

	return &organizations, nil
}

func (c *Client) OrganizationSearch(query string, limit int) (*[]map[string]interface{}, error) {
	var organizations []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/organizations/search?query=%s&limit=%d", query, limit)), nil)
	if err != nil {
		return &organizations, err
	}

	if err = c.SendWithAuth(req, &organizations); err != nil {
		return &organizations, err
	}

	return &organizations, nil
}

func (c *Client) OrganizationShow(organizationID int) (*map[string]interface{}, error) {
	var organization map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/organizations/%d", organizationID)), nil)
	if err != nil {
		return &organization, err
	}

	if err = c.SendWithAuth(req, &organization); err != nil {
		return &organization, err
	}

	return &organization, nil
}

func (c *Client) OrganizationCreate(o *map[string]interface{}) (*map[string]interface{}, error) {
	var organization map[string]interface{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/organizations"), o)
	if err != nil {
		return &organization, err
	}

	if err = c.SendWithAuth(req, &organization); err != nil {
		return &organization, err
	}

	return &organization, nil
}

func (c *Client) OrganizationUpdate(organizationID int, o *map[string]interface{}) (*map[string]interface{}, error) {
	var organization map[string]interface{}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/organizations/%d", organizationID)), o)
	if err != nil {
		return &organization, err
	}

	if err = c.SendWithAuth(req, &organization); err != nil {
		return &organization, err
	}

	return &organization, nil
}

func (c *Client) OrganizationDelete(organizationID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/organizations/%d", organizationID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
