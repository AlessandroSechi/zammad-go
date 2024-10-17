package zammad

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// Ticket is a zammad ticket.
type Ticket struct {
	Title                 string        `json:"title"`
	Group                 string        `json:"group"`
	OwnerID               int           `json:"owner_id,omitempty"`
	ID                    int           `json:"id,omitempty"`
	Article               TicketArticle `json:"article,omitempty"`
	GroupID               int           `json:"group_id,omitempty"`
	PriorityID            int           `json:"priority_id,omitempty"`
	StateID               int           `json:"state_id,omitempty"`
	State                 string        `json:"state,omitempty"`
	OrganizationID        int           `json:"organization_id"`
	Number                string        `json:"number,omitempty"`
	Customer              string        `json:"customer,omitempty"`
	CustomerID            int           `json:"customer_id,omitempty"`
	LastContactAt         time.Time     `json:"last_contact_at,omitempty"`
	LastContactAgentAt    time.Time     `json:"last_contact_agent_at,omitempty"`
	LastContactCustomerAt time.Time     `json:"last_contact_customer_at,omitempty"`
	CreateArticleTypeID   int           `json:"create_article_type_id,omitempty"`
	CreateArticleSenderID int           `json:"create_article_sender_id,omitempty"`
	ArticleCount          int           `json:"article_count,omitempty"`
	UpdatedByID           int           `json:"updated_by_id,omitempty"`
	CreatedByID           int           `json:"created_by_id,omitempty"`
	CreatedAt             time.Time     `json:"created_at,omitempty"`
	UpdatedAt             time.Time     `json:"updated_at,omitempty"`
}

func (c *Client) TicketListResult(opts ...Option) *Result[Ticket] {
	return &Result[Ticket]{
		res:     nil,
		resFunc: c.TicketListWithOptions,
		opts:    NewRequestOptions(opts...),
	}
}

func (c *Client) TicketList() ([]Ticket, error) {
	return c.TicketListResult().FetchAll()
}

func (c *Client) TicketListWithOptions(ro RequestOptions) ([]Ticket, error) {
	var tickets []Ticket

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, "/api/v1/tickets"), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = ro.URLParams()

	if err = c.sendWithAuth(req, &tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

// TicketSearch searches for tickets. See https://docs.zammad.org/en/latest/api/ticket/index.html#search.
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
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/search?query=%s&limit=%d", url.QueryEscape(query), limit)), nil)
	if err != nil {
		return nil, err
	}

	if err = c.sendWithAuth(req, &ticksearch); err != nil {
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

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/%d", ticketID)), nil)
	if err != nil {
		return ticket, err
	}

	if err = c.sendWithAuth(req, &ticket); err != nil {
		return ticket, err
	}

	return ticket, nil
}

// TicketCreate is used to create a ticket. For this you need to assemble a bare-bones Ticket:
//
//	ticket := Ticket{
//		Title:      "your subject",
//		Group:      "your group",
//		CustomerID: 10, // your customer ID
//		Article: TicketArticle{
//			Subject: "subject of comment",
//			Body: "body of comment",
//		},
//	}
func (c *Client) TicketCreate(t Ticket) (Ticket, error) {
	var ticket Ticket

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/tickets"), t)
	if err != nil {
		return ticket, err
	}

	if err = c.sendWithAuth(req, &ticket); err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (c *Client) TicketUpdate(ticketID int, t Ticket) (Ticket, error) {
	var ticket Ticket

	req, err := c.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/%d", ticketID)), t)
	if err != nil {
		return ticket, err
	}

	if err = c.sendWithAuth(req, &ticket); err != nil {
		return ticket, err
	}

	return ticket, nil
}

func (c *Client) TicketDelete(ticketID int) error {

	req, err := c.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tickets/%d", ticketID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
