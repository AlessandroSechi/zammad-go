package zammad

import "fmt"

// TODO
type TicketPriority struct{}

func (c *Client) TicketPriorityList() ([]TicketPriority, error) {
	var ticketPriorities []TicketPriority

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_priorities"), nil)
	if err != nil {
		return ticketPriorities, err
	}

	if err = c.SendWithAuth(req, &ticketPriorities); err != nil {
		return ticketPriorities, err
	}

	return ticketPriorities, nil
}

func (c *Client) TicketPriorityShow(ticketPriorityID int) (TicketPriority, error) {
	var ticketPriority TicketPriority

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_priorities/%d", ticketPriorityID)), nil)
	if err != nil {
		return ticketPriority, err
	}

	if err = c.SendWithAuth(req, &ticketPriority); err != nil {
		return ticketPriority, err
	}

	return ticketPriority, nil
}

func (c *Client) TicketPriorityCreate(t TicketPriority) (TicketPriority, error) {
	var ticketPriority TicketPriority

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_priorities"), t)
	if err != nil {
		return ticketPriority, err
	}

	if err = c.SendWithAuth(req, &ticketPriority); err != nil {
		return ticketPriority, err
	}

	return ticketPriority, nil
}

func (c *Client) TicketPriorityUpdate(ticketPriorityID int, t TicketPriority) (TicketPriority, error) {
	var ticketPriority TicketPriority

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_priorities/%d", ticketPriorityID)), t)
	if err != nil {
		return ticketPriority, err
	}

	if err = c.SendWithAuth(req, &ticketPriority); err != nil {
		return ticketPriority, err
	}

	return ticketPriority, nil
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
