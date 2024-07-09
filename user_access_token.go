package zammad

import "fmt"

// UserAccessToken is a Zammad User access token. See https://docs.zammad.org/en/latest/api/user-access-token.html.
// TODO(miek): make is usable for create and list
type UserAccessToken map[string]interface{}

func (c *Client) UserAccessTokenList() ([]UserAccessToken, error) {
	var userAccessTokens []UserAccessToken

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/user_access_token"), nil)
	if err != nil {
		return userAccessTokens, err
	}

	if err = c.SendWithAuth(req, userAccessTokens); err != nil {
		return userAccessTokens, err
	}

	return userAccessTokens, nil
}

func (c *Client) UserAccessTokenCreate(t UserAccessToken) (UserAccessToken, error) {
	var userAccessToken map[string]interface{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/user_access_token"), t)
	if err != nil {
		return userAccessToken, err
	}

	if err = c.SendWithAuth(req, userAccessToken); err != nil {
		return userAccessToken, err
	}

	return userAccessToken, nil
}

func (c *Client) UserAccessTokenDelete(tokenID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/user_access_token/%d", tokenID)), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
