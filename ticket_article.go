package zammad

import (
	"fmt"
	"net/http"
	"time"
)

// TicketArticle represents a Zammad ticket article.
type TicketArticle struct {
	ID       int    `json:"id,omitempty"`
	TicketID int    `json:"ticket_id,omitempty"`
	TypeID   int    `json:"type_id,omitempty"`
	SenderID int    `json:"sender_id,omitempty"`
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	// Don't know if I want to keep the 'any' types here?
	Subject     string    `json:"subject"`
	Body        string    `json:"body"`
	Cc          any       `json:"cc"`
	ReplyTo     any       `json:"reply_to"`
	ContentType string    `json:"content_type"`
	Internal    bool      `json:"internal"`
	UpdatedByID int       `json:"updated_by_id,omitempty"`
	CreatedByID int       `json:"created_by_id,omitempty"`
	OriginByID  any       `json:"origin_by_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Type        string    `json:"type,omitempty"`
	Sender      string    `json:"sender,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	UpdatedBy   string    `json:"updated_by,omitempty"`
}

func (c *Client) TicketArticleByTicket(ticketID int) ([]TicketArticle, error) {
	var ticketArticles []TicketArticle

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_articles/by_ticket/%d", ticketID)), nil)
	if err != nil {
		return ticketArticles, err
	}

	if err = c.sendWithAuth(req, &ticketArticles); err != nil {
		return ticketArticles, err
	}

	return ticketArticles, nil
}

func (c *Client) TicketArticleShow(ticketArticleID int) (TicketArticle, error) {
	var ticketArticle TicketArticle

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_articles/%d", ticketArticleID)), nil)
	if err != nil {
		return ticketArticle, err
	}

	if err = c.sendWithAuth(req, &ticketArticle); err != nil {
		return ticketArticle, err
	}

	return ticketArticle, nil
}

func (c *Client) TicketArticleCreate(t TicketArticle) (TicketArticle, error) {
	var ticketArticle TicketArticle

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_articles"), t)
	if err != nil {
		return ticketArticle, err
	}

	if err = c.sendWithAuth(req, &ticketArticle); err != nil {
		return ticketArticle, err
	}

	return ticketArticle, nil
}
