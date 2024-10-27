package zammad

import (
	"fmt"
	"net/http"
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

func (c *Client) TicketPriorityListResult(opts ...Option) *Result[TicketPriority] {
	return &Result[TicketPriority]{
		res:     nil,
		resFunc: c.TicketPriorityListWithOptions,
		opts:    NewRequestOptions(opts...),
	}
}

func (c *Client) TicketPriorityList() ([]TicketPriority, error) {
	return c.TicketPriorityListResult().FetchAll()
}

func (c *Client) TicketPriorityListWithOptions(ro RequestOptions) ([]TicketPriority, error) {
	var ticketPriorities []TicketPriority

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_priorities"), nil)
	if err != nil {
		return ticketPriorities, err
	}

	req.URL.RawQuery = ro.URLParams()

	if err = c.sendWithAuth(req, &ticketPriorities); err != nil {
		return ticketPriorities, err
	}

	return ticketPriorities, nil
}

func (c *Client) TicketPriorityShow(ticketPriorityID int) (TicketPriority, error) {
	var ticketPriority TicketPriority

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_priorities/%d", ticketPriorityID)), nil)
	if err != nil {
		return ticketPriority, err
	}

	if err = c.sendWithAuth(req, &ticketPriority); err != nil {
		return ticketPriority, err
	}

	return ticketPriority, nil
}

func (c *Client) TicketPriorityCreate(t TicketPriority) (TicketPriority, error) {
	var ticketPriority TicketPriority

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_priorities"), t)
	if err != nil {
		return ticketPriority, err
	}

	if err = c.sendWithAuth(req, &ticketPriority); err != nil {
		return ticketPriority, err
	}

	return ticketPriority, nil
}

func (c *Client) TicketPriorityUpdate(ticketPriorityID int, t TicketPriority) (TicketPriority, error) {
	var ticketPriority TicketPriority

	req, err := c.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_priorities/%d", ticketPriorityID)), t)
	if err != nil {
		return ticketPriority, err
	}

	if err = c.sendWithAuth(req, &ticketPriority); err != nil {
		return ticketPriority, err
	}

	return ticketPriority, nil
}

func (c *Client) TicketPriorityDelete(ticketPriorityID int) error {

	req, err := c.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_priorities/%d", ticketPriorityID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
