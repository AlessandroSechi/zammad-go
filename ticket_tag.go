package zammad

import (
	"fmt"
	"net/http"
)

func (c *Client) TicketTagByTicket(ticketID int) ([]Tag, error) {
	var tags struct {
		Tags []string
	}

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tags?object=Ticket&o_id=%d", ticketID)), nil)
	if err != nil {
		return nil, err
	}

	if err = c.sendWithAuth(req, &tags); err != nil {
		return nil, err
	}

	tags1 := make([]Tag, len(tags.Tags))
	for i := range tags.Tags {
		tags1[i] = Tag{Name: tags.Tags[i]}
	}

	return tags1, nil
}
