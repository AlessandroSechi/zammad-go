package zammad

import (
	"fmt"
	"time"
)

// Organization represent a Zammad organisation. See https://docs.zammad.org/en/latest/api/organization.html
type Organization struct {
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	Shared             bool      `json:"shared"`
	Domain             string    `json:"domain"`
	DomainAssignment   bool      `json:"domain_assignment"`
	Active             bool      `json:"active"`
	Note               string    `json:"note"`
	Vip                bool      `json:"vip"`
	UpdatedByID        int       `json:"updated_by_id"`
	CreatedByID        int       `json:"created_by_id"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	MemberIds          []int     `json:"member_ids"`
	SecondaryMemberIds []int     `json:"secondary_member_ids"`
}

func (c *Client) OrganizationList() ([]Organization, error) {
	var organizations []Organization

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/organizations"), nil)
	if err != nil {
		return organizations, err
	}

	if err = c.sendWithAuth(req, &organizations); err != nil {
		return organizations, err
	}

	return organizations, nil
}

func (c *Client) OrganizationSearch(query string, limit int) ([]Organization, error) {
	var organizations []Organization

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/organizations/search?query=%slimit=%d", query, limit)), nil)
	if err != nil {
		return organizations, err
	}

	if err = c.sendWithAuth(req, &organizations); err != nil {
		return organizations, err
	}

	return organizations, nil
}

func (c *Client) OrganizationShow(organizationID int) (Organization, error) {
	var organization Organization

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/organizations/%d", organizationID)), nil)
	if err != nil {
		return organization, err
	}

	if err = c.sendWithAuth(req, &organization); err != nil {
		return organization, err
	}

	return organization, nil
}

func (c *Client) OrganizationCreate(o Organization) (Organization, error) {
	var organization Organization

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/organizations"), o)
	if err != nil {
		return organization, err
	}

	if err = c.sendWithAuth(req, &organization); err != nil {
		return organization, err
	}

	return organization, nil
}

func (c *Client) OrganizationUpdate(organizationID int, o Organization) (Organization, error) {
	var organization Organization

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/organizations/%d", organizationID)), o)
	if err != nil {
		return organization, err
	}

	if err = c.sendWithAuth(req, &organization); err != nil {
		return organization, err
	}

	return organization, nil
}

func (c *Client) OrganizationDelete(organizationID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/organizations/%d", organizationID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
