package zammad

import "fmt"

func (c *Client) TicketList() (*[]map[string]interface{}, error) {
	var tickets []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/tickets"), nil)
	if err != nil {
		return &tickets, err
	}

	if err = c.SendWithAuth(req, &tickets); err != nil {
		return &tickets, err
	}

	return &tickets, nil
}

func (c *Client) TicketSearch(query string, limit int) (*[]map[string]interface{}, error) {
	var tickets []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/search?query=%s&limit=%d", query, limit)), nil)
	if err != nil {
		return &tickets, err
	}

	if err = c.SendWithAuth(req, &tickets); err != nil {
		return &tickets, err
	}

	return &tickets, nil
}

func (c *Client) TicketShow(ticketID int) (*map[string]interface{}, error) {
	var ticket map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/%d", ticketID)), nil)
	if err != nil {
		return &ticket, err
	}

	if err = c.SendWithAuth(req, &ticket); err != nil {
		return &ticket, err
	}

	return &ticket, nil
}

func (c *Client) TicketCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var ticket map[string]interface{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/tickets"), t)
	if err != nil {
		return &ticket, err
	}

	if err = c.SendWithAuth(req, &ticket); err != nil {
		return &ticket, err
	}

	return &ticket, nil
}

func (c *Client) TicketUpdate(ticketID int, t *map[string]interface{}) (*map[string]interface{}, error) {
	var ticket map[string]interface{}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/%d", ticketID)), t)
	if err != nil {
		return &ticket, err
	}

	if err = c.SendWithAuth(req, &ticket); err != nil {
		return &ticket, err
	}

	return &ticket, nil
}

func (c *Client) TicketDelete(ticketID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/%d", ticketID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
