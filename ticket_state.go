package zammad

import (
	"fmt"
	"net/http"
	"time"
)

// TicketState is a Zammad ticket state
type TicketState struct {
	ID               int       `json:"id"`
	StateTypeID      int       `json:"state_type_id"`
	Name             string    `json:"name"`
	IgnoreEscalation bool      `json:"ignore_escalation"`
	DefaultCreate    bool      `json:"default_create"`
	DefaultFollowUp  bool      `json:"default_follow_up"`
	Active           bool      `json:"active"`
	UpdatedByID      int       `json:"updated_by_id"`
	CreatedByID      int       `json:"created_by_id"`
	CreatedAt        time.Time `json:"created_at"`
}

func (c *Client) TicketStateListResult(opts ...Option) *Result[TicketState] {
	return &Result[TicketState]{
		res:     nil,
		resFunc: c.TicketStateListWithOptions,
		opts:    NewRequestOptions(opts...),
	}
}

func (c *Client) TicketStateList() ([]TicketState, error) {
	return c.TicketStateListResult().FetchAll()
}

func (c *Client) TicketStateListWithOptions(ro RequestOptions) ([]TicketState, error) {
	var ticketStates []TicketState

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_states"), nil)
	if err != nil {
		return ticketStates, err
	}

	req.URL.RawQuery = ro.URLParams()

	if err = c.sendWithAuth(req, &ticketStates); err != nil {
		return ticketStates, err
	}

	return ticketStates, nil
}

func (c *Client) TicketStateShow(ticketStateID int) (TicketState, error) {
	var ticketState TicketState

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_states/%d", ticketStateID)), nil)
	if err != nil {
		return ticketState, err
	}

	if err = c.sendWithAuth(req, &ticketState); err != nil {
		return ticketState, err
	}

	return ticketState, nil
}

func (c *Client) TicketStateCreate(t TicketState) (TicketState, error) {
	var ticketState TicketState

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_states"), t)
	if err != nil {
		return ticketState, err
	}

	if err = c.sendWithAuth(req, &ticketState); err != nil {
		return ticketState, err
	}

	return ticketState, nil
}

func (c *Client) TicketStateUpdate(ticketStateID int, t TicketState) (TicketState, error) {
	var ticketState TicketState

	req, err := c.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_states/%d", ticketStateID)), t)
	if err != nil {
		return ticketState, err
	}

	if err = c.sendWithAuth(req, &ticketState); err != nil {
		return ticketState, err
	}

	return ticketState, nil
}

func (c *Client) TicketStateDelete(ticketStateID int) error {

	req, err := c.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_states/%d", ticketStateID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
