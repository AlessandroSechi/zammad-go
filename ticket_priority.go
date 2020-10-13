package zammad

import "fmt"

func (c *Client) TicketPriorityList() (*[]map[string]interface{}, error) {
	var ticketPriorities []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_priorities"), nil)
	if err != nil {
		return &ticketPriorities, err
	}

	if err = c.SendWithAuth(req, &ticketPriorities); err != nil {
		return &ticketPriorities, err
	}

	return &ticketPriorities, nil
}

func (c *Client) TicketPriorityShow(ticketPriorityID int) (*map[string]interface{}, error) {
	var ticketPriority map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_priorities/%d", ticketPriorityID)), nil)
	if err != nil {
		return &ticketPriority, err
	}

	if err = c.SendWithAuth(req, &ticketPriority); err != nil {
		return &ticketPriority, err
	}

	return &ticketPriority, nil
}

func (c *Client) TicketPriorityCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var ticketPriority map[string]interface{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_priorities"), t)
	if err != nil {
		return &ticketPriority, err
	}

	if err = c.SendWithAuth(req, &ticketPriority); err != nil {
		return &ticketPriority, err
	}

	return &ticketPriority, nil
}

func (c *Client) TicketPriorityUpdate(ticketPriorityID int, t *map[string]interface{}) (*map[string]interface{}, error) {
	var ticketPriority map[string]interface{}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_priorities/%d", ticketPriorityID)), t)
	if err != nil {
		return &ticketPriority, err
	}

	if err = c.SendWithAuth(req, &ticketPriority); err != nil {
		return &ticketPriority, err
	}

	return &ticketPriority, nil
}

func (c *Client) TicketPriorityDelete(ticketPriorityID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_priorities/%d", ticketPriorityID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
