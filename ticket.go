package zammad

import (
	"fmt"
	"time"
)

// Ticket is a zammad ticket.
type Ticket struct {
	ID                    int       `json:"id"`
	GroupID               int       `json:"group_id"`
	PriorityID            int       `json:"priority_id"`
	StateID               int       `json:"state_id"`
	OrganizationID        int       `json:"organization_id"`
	Number                string    `json:"number"`
	Title                 string    `json:"title"`
	OwnerID               int       `json:"owner_id"`
	CustomerID            int       `json:"customer_id"`
	LastContactAt         time.Time `json:"last_contact_at"`
	LastContactAgentAt    time.Time `json:"last_contact_agent_at"`
	LastContactCustomerAt time.Time `json:"last_contact_customer_at"`
	CreateArticleTypeID   int       `json:"create_article_type_id"`
	CreateArticleSenderID int       `json:"create_article_sender_id"`
	ArticleCount          int       `json:"article_count"`
	UpdatedByID           int       `json:"updated_by_id"`
	CreatedByID           int       `json:"created_by_id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func (c *Client) TicketList() ([]Ticket, error) {
	var tickets []Ticket

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/tickets"), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, &tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (c *Client) TicketSearch(query string, limit int) ([]Ticket, error) {
	type Assets struct {
		AssetTicket map[int]Ticket `json:"ticket"`
	}

	type TickSearch struct {
		Tickets []int `json:"tickets"`
		Count   int   `json:"tickets_count"`
		Assets  `json:"assets"`
	}

	var ticksearch TickSearch
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/search?query=%s&limit=%d", query, limit)), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, &ticksearch); err != nil {
		return nil, err
	}

	tickets := make([]Ticket, ticksearch.Count)
	i := 0
	for _, t1 := range ticksearch.Assets.AssetTicket {
		tickets[i] = t1
		i++
	}
	return tickets, nil
}

func (c *Client) TicketShow(ticketID int) (Ticket, error) {
	var ticket Ticket

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/%d", ticketID)), nil)
	if err != nil {
		return ticket, err
	}

	if err = c.SendWithAuth(req, &ticket); err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (c *Client) TicketCreate(t Ticket) (Ticket, error) {
	var ticket Ticket

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/tickets"), t)
	if err != nil {
		return ticket, err
	}

	if err = c.SendWithAuth(req, &ticket); err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (c *Client) TicketUpdate(ticketID int, t Ticket) (Ticket, error) {
	var ticket Ticket

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/%d", ticketID)), t)
	if err != nil {
		return ticket, err
	}

	if err = c.SendWithAuth(req, &ticket); err != nil {
		return ticket, err
	}

	return ticket, nil
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
