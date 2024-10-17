package zammad

import (
	"fmt"
	"net/http"
)

// Object represent a Zammad object. See https://docs.zammad.org/en/latest/api/object.html.
// Also note the warning there:
//
//	Adjusting objects via API can cause serious issues with your instance. Proceed with absolute caution and ensure
//	to adjust any of Zammads default fields.
type Object *map[string]interface{}

func (c *Client) ObjectListResult(opts ...Option) *Result[Object] {
	return &Result[Object]{
		res:     nil,
		resFunc: c.ObjectListWithOptions,
		opts:    NewRequestOptions(opts...),
	}
}

func (c *Client) ObjectList() ([]Object, error) {
	return c.ObjectListResult().FetchAll()
}

func (c *Client) ObjectListWithOptions(ro RequestOptions) ([]Object, error) {
	var objects []Object

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, "/api/v1/object_manager_attributes"), nil)
	if err != nil {
		return objects, err
	}

	req.URL.RawQuery = ro.URLParams()

	if err = c.sendWithAuth(req, objects); err != nil {
		return objects, err
	}

	return objects, nil
}

func (c *Client) ObjectShow(objectID int) (Object, error) {
	var object Object

	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/object_manager_attributes/%d", objectID)), nil)
	if err != nil {
		return object, err
	}

	if err = c.sendWithAuth(req, object); err != nil {
		return object, err
	}

	return object, nil
}

func (c *Client) ObjectCreate(o Object) (Object, error) {
	var object Object

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/object_manager_attributes"), o)
	if err != nil {
		return object, err
	}

	if err = c.sendWithAuth(req, object); err != nil {
		return object, err
	}

	return object, nil
}

func (c *Client) ObjectUpdate(objectID int, o Object) (Object, error) {
	var object Object

	req, err := c.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/object_manager_attributes/%d", objectID)), o)
	if err != nil {
		return object, err
	}

	if err = c.sendWithAuth(req, object); err != nil {
		return object, err
	}

	return object, nil
}

func (c *Client) ObjectExecuteDatabaseMigration() error {

	req, err := c.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", c.Url, "/api/v1/object_manager_attributes_execute_migrations"), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
