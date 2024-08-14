package zammad

import (
	"fmt"
	"time"
)

// Group is a Zammad group.
type Group struct {
	ID                 int       `json:"id,omitempty"`
	Name               string    `json:"name"`
	NameLast           string    `json:"name_last"`
	FollowUpPossible   string    `json:"follow_up_possible"`
	FollowUpAssignment bool      `json:"follow_up_assignment"`
	Active             bool      `json:"active"`
	UpdatedByID        int       `json:"updated_by_id"`
	CreatedByID        int       `json:"created_by_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	UserIds            []int     `json:"user_ids"`
}

func (c *Client) GroupList() ([]Group, error) {
	var groups []Group

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/groups"), nil)
	if err != nil {
		return nil, err
	}

	if err = c.sendWithAuth(req, &groups); err != nil {
		return nil, err
	}

	return groups, nil
}

func (c *Client) GroupShow(groupID int) (Group, error) {
	var group Group

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/groups/%d", groupID)), nil)
	if err != nil {
		return group, err
	}

	if err = c.sendWithAuth(req, &group); err != nil {
		return group, err
	}

	return group, nil
}

func (c *Client) GroupCreate(g Group) (Group, error) {
	var group Group

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/groups"), g)
	if err != nil {
		return group, err
	}

	if err = c.sendWithAuth(req, &group); err != nil {
		return group, err
	}

	return group, nil
}

func (c *Client) GroupUpdate(groupID int, g Group) (Group, error) {
	var group Group

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/groups/%d", groupID)), g)
	if err != nil {
		return group, err
	}

	if err = c.sendWithAuth(req, &group); err != nil {
		return group, err
	}

	return group, nil
}

func (c *Client) GroupDelete(groupID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/groups/%d", groupID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
