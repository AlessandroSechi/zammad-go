package zammad

import "fmt"

func (c *Client) TicketStateList() (*[]map[string]interface{}, error) {
	var ticketStates []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_states"), nil)
	if err != nil {
		return &ticketStates, err
	}

	if err = c.SendWithAuth(req, &ticketStates); err != nil {
		return &ticketStates, err
	}

	return &ticketStates, nil
}

func (c *Client) TicketStateShow(ticketStateID int) (*map[string]interface{}, error) {
	var ticketState map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_states/%d", ticketStateID)), nil)
	if err != nil {
		return &ticketState, err
	}

	if err = c.SendWithAuth(req, &ticketState); err != nil {
		return &ticketState, err
	}

	return &ticketState, nil
}

func (c *Client) TicketStateCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var ticketState map[string]interface{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_states"), t)
	if err != nil {
		return &ticketState, err
	}

	if err = c.SendWithAuth(req, &ticketState); err != nil {
		return &ticketState, err
	}

	return &ticketState, nil
}

func (c *Client) TicketStateUpdate(ticketStateID int, t *map[string]interface{}) (*map[string]interface{}, error) {
	var ticketState map[string]interface{}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_states/%d", ticketStateID)), t)
	if err != nil {
		return &ticketState, err
	}

	if err = c.SendWithAuth(req, &ticketState); err != nil {
		return &ticketState, err
	}

	return &ticketState, nil
}

func (c *Client) TicketStateDelete(ticketStateID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_states/%d", ticketStateID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
