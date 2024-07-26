package zammad

import (
	"fmt"
	"time"
)

// OnlineNotification represent a Zammad online notification. See https://docs.zammad.org/en/latest/api/notification.html.
type OnlineNotification struct {
	ID             int       `json:"id"`
	OID            int       `json:"o_id"`
	ObjectLookupID int       `json:"object_lookup_id"`
	TypeLookupID   int       `json:"type_lookup_id"`
	UserID         int       `json:"user_id"`
	Seen           bool      `json:"seen"`
	UpdatedByID    int       `json:"updated_by_id"`
	CreatedByID    int       `json:"created_by_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (c *Client) OnlineNotificationList() ([]OnlineNotification, error) {
	var notifications []OnlineNotification

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/online_notifications"), nil)
	if err != nil {
		return notifications, err
	}

	if err = c.sendWithAuth(req, &notifications); err != nil {
		return notifications, err
	}

	return notifications, nil
}

func (c *Client) OnlineNotificationShow(notificationID int) (OnlineNotification, error) {
	var notification OnlineNotification

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/online_notifications/%d", notificationID)), nil)
	if err != nil {
		return notification, err
	}

	if err = c.sendWithAuth(req, &notification); err != nil {
		return notification, err
	}

	return notification, nil
}

func (c *Client) OnlineNotificationUpdate(notificationID int, n OnlineNotification) (OnlineNotification, error) {
	var notification OnlineNotification

	req, err := c.NewRequest("PUT", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/online_notifications/%d", notificationID)), n)
	if err != nil {
		return notification, err
	}

	if err = c.sendWithAuth(req, notification); err != nil {
		return notification, err
	}

	return notification, nil
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
