package zammad

import (
	"fmt"
	"time"
)

// TicketPriority represent a Zammad ticket priority. See https://docs.zammad.org/en/latest/api/ticket/priorities.html.
type TicketPriority struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	DefaultCreate bool      `json:"default_create"`
	UIIcon        string    `json:"ui_icon"`
	UIColor       string    `json:"ui_color"`
	Note          any       `json:"note"`
	Active        bool      `json:"active"`
	UpdatedByID   int       `json:"updated_by_id"`
	CreatedByID   int       `json:"created_by_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

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
