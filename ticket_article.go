package zammad

import "fmt"

func (c *Client) TicketArticleByTicket(ticketID int) (*[]map[string]interface{}, error) {
	var ticketArticles []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_articles/by_ticket/%d", ticketID)), nil)
	if err != nil {
		return &ticketArticles, err
	}

	if err = c.SendWithAuth(req, &ticketArticles); err != nil {
		return &ticketArticles, err
	}

	return &ticketArticles, nil
}

func (c *Client) TicketArticleShow(ticketArticleID int) (*map[string]interface{}, error) {
	var ticketArticle map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/ticket_articles/%d", ticketArticleID)), nil)
	if err != nil {
		return &ticketArticle, err
	}

	if err = c.SendWithAuth(req, &ticketArticle); err != nil {
		return &ticketArticle, err
	}

	return &ticketArticle, nil
}

func (c *Client) TicketArticleCreate(t *map[string]interface{}) (*map[string]interface{}, error) {
	var ticketArticle map[string]interface{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/ticket_articles"), t)
	if err != nil {
		return &ticketArticle, err
	}

	if err = c.SendWithAuth(req, &ticketArticle); err != nil {
		return &ticketArticle, err
	}

	return &ticketArticle, nil
}