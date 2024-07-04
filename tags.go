package zammad

import "fmt"

// Tag is a Zammad tag.
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (c *Client) TagList(ticketID int) ([]Tag, error) {
	var tags struct {
		Tags []string
	}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tags?object=Ticket&o_id=%d", ticketID)), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, &tags); err != nil {
		return nil, err
	}

	tags1 := make([]Tag, len(tags.Tags))
	for i := range tags.Tags {
		tags1[i] = Tag{Name: tags.Tags[i]}
	}

	return tags1, nil
}

func (c *Client) TagSearch(term string) ([]Tag, error) {
	var tags []Tag

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_search?term=%s", term)), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func (c *Client) TagAdd(t Tag) error {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/tags/add"), t)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagRemove(t Tag) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, "/api/v1/tags/remove"), t)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminList() ([]Tag, error) {
	var tags []Tag

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/tag_list"), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func (c *Client) TagAdminCreate(t Tag) error {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/tag_list"), t)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminRename(tagID int, t Tag) error {

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_list/%d", tagID)), t)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminDelete(tagID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_list/%d", tagID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
