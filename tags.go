package zammad

import (
	"fmt"
	"net/http"
	"net/url"
)

// Tag is a Zammad tag. See https://docs.zammad.org/en/latest/api/ticket/tags.html.
type Tag struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func (c *Client) TagSearch(term string) ([]Tag, error) {
	var tags []Tag

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_search?term=%s", url.QueryEscape(term))), nil)
	if err != nil {
		return nil, err
	}

	if err = c.sendWithAuth(req, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func (c *Client) TagAdd(t Tag) error {

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/tags/add"), t)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagRemove(t Tag) error {

	req, err := c.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.Url, "/api/v1/tags/remove"), t)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminList() ([]Tag, error) {
	var tags []Tag

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, "/api/v1/tag_list"), nil)
	if err != nil {
		return nil, err
	}

	if err = c.sendWithAuth(req, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}

func (c *Client) TagAdminCreate(t Tag) error {

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/tag_list"), t)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminRename(tagID int, t Tag) error {

	req, err := c.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_list/%d", tagID)), t)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) TagAdminDelete(tagID int) error {

	req, err := c.NewRequest(http.MethodDelete, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/tag_list/%d", tagID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
