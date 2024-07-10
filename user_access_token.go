package zammad

import (
	"fmt"
	"time"
)

// UserAccessToken is a Zammad User access token. See https://docs.zammad.org/en/latest/api/user-access-token.html.
type UserAccessToken struct {
	ID          int      `json:"id"`
	Token       string   `json:"token,omitempty"` // Token is only set in the returned UserAccessToken in UserAccessTokenCreate.
	Name        string   `json:"name,omitempty"`  // Name is must be set and is only used in UserAccessTokenCreate.
	UserID      int      `json:"user_id"`
	Action      string   `json:"action"`
	Label       string   `json:"label"`
	Permission  []string `json:"permission,omitempty"` // Permission must be set and is only used in UserAccessTokenCreate.
	Preferences struct {
		Permission []string `json:"permission"`
	} `json:"preferences"`
	LastUsedAt time.Time `json:"last_used_at"`
	ExpiresAt  time.Time `json:"expires_at,omitempty"` // ExpiresAt must be set when using in UserAccessTokenCreate.
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Permissions holds all available Zammad User access token permissions. It's only used UserAccessTokenList.
	// When returns UserAccessTokens from UserAccessTokenList only the first access token will have a filled out
	// permissions struct. For the remaining token this will be empty.
	Permissions []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Note        string `json:"note"`
		Preferences struct {
			Translations []string `json:"translations"`
			Disabled     bool     `json:"disabled"`
		} `json:"preferences,omitempty"`
		Active      bool      `json:"active"`
		AllowSignup bool      `json:"allow_signup"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"permissions,omitempty"`
}

// UserAccessTokenPermission is a Zammad User access token permissions. See
// https://docs.zammad.org/en/latest/api/user-access-token.html.
type UserAccessTokenPermission struct{}

func (c *Client) UserAccessTokenList() ([]UserAccessToken, error) {
	// define Permission here too because we don't want to leak the type as it's not much of use.
	type Permission struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Note        string `json:"note"`
		Preferences struct {
			Translations []string `json:"translations"`
			Disabled     bool     `json:"disabled"`
		} `json:"preferences,omitempty"`
		Active      bool      `json:"active"`
		AllowSignup bool      `json:"allow_signup"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	type TockList struct {
		Tokens      []UserAccessToken `json:"tokens"`
		Permissions []Permission      `json:"permissions"`
	}

	var tockList TockList
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/user_access_token"), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, &tockList); err != nil {
		return nil, err
	}

	userAccessTokens := make([]UserAccessToken, len(tockList.Tokens))
	for i, t := range tockList.Tokens {
		userAccessTokens[i] = t
		if i == 0 {
			userAccessTokens[i].Permissions = t.Permissions
		}
	}

	return userAccessTokens, nil
}

func (c *Client) UserAccessTokenCreate(t UserAccessToken) (UserAccessToken, error) {
	var userAccessToken UserAccessToken

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
