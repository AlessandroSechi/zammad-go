package zammad

import (
	"fmt"
	"time"
)

// TIcketArticle represents a Zammad ticket article.
type TicketArticle struct {
	ID       int    `json:"id"`
	TicketID int    `json:"ticket_id"`
	TypeID   int    `json:"type_id"`
	SenderID int    `json:"sender_id"`
	From     string `json:"from"`
	To       string `json:"to"`
	// Don't know if I want to keep the 'any' types here?
	Cc          any       `json:"cc"`
	Subject     any       `json:"subject"`
	ReplyTo     any       `json:"reply_to"`
	ContentType string    `json:"content_type"`
	Body        string    `json:"body"`
	Internal    bool      `json:"internal"`
	UpdatedByID int       `json:"updated_by_id"`
	CreatedByID int       `json:"created_by_id"`
	OriginByID  any       `json:"origin_by_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Type        string    `json:"type"`
	Sender      string    `json:"sender"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}

func (c *Client) TicketArticleByTicket(ticketID int) ([]TicketArticle, error) {
	var ticketArticles []TicketArticle

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_articles/by_ticket/%d", ticketID)), nil)
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

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_articles/%d", ticketArticleID)), nil)
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

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_articles"), t)
	if err != nil {
		return ticketArticle, err
	}

	if err = c.sendWithAuth(req, &ticketArticle); err != nil {
		return ticketArticle, err
	}

	return ticketArticle, nil
}
