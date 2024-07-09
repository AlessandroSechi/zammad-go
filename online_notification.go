package zammad

import "fmt"

func (c *Client) OnlineNotificationList() (*[]map[string]interface{}, error) {
	var notifications []map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/online_notifications"), nil)
	if err != nil {
		return &notifications, err
	}

	if err = c.sendWithAuth(req, &notifications); err != nil {
		return &notifications, err
	}

	return &notifications, nil
}

func (c *Client) OnlineNotificationShow(notificationID int) (*map[string]interface{}, error) {
	var notification map[string]interface{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/online_notifications/%d", notificationID)), nil)
	if err != nil {
		return &notification, err
	}

	if err = c.sendWithAuth(req, &notification); err != nil {
		return &notification, err
	}

	return &notification, nil
}

func (c *Client) OnlineNotificationUpdate(notificationID int, n *map[string]interface{}) (*map[string]interface{}, error) {
	var notification map[string]interface{}

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/online_notifications/%d", notificationID)), n)
	if err != nil {
		return &notification, err
	}

	if err = c.sendWithAuth(req, &notification); err != nil {
		return &notification, err
	}

	return &notification, nil
}

func (c *Client) OnlineNotificationDelete(notificationID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/online_notifications/%d", notificationID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) OnlineNotificationMarkAllAsRead() error {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/online_notifications/mark_all_as_read"), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
