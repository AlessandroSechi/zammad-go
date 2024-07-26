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
	ExpiresAt  Date      `json:"expires_at,omitempty"` // ExpiresAt must be set when using in UserAccessTokenCreate.
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	// Permissions holds all available Zammad User access token permissions. It's only used UserAccessTokenList.
	// When returns UserAccessTokens from UserAccessTokenList only the first access token will have a filled out
	// permissions struct. For the remaining token this will be empty.
	Permissions []Permission `json:"permissions,omitempty"`
}

// Permission is a Zammad permission.
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

// Date is a timestamp that is only specified as YYYY-MM-DD, without time zone or a clock time.
type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(`"2006-01-02"`, string(data))
	if err != nil {
		return err
	}
	(*d).Time = t
	return nil
}

func (c *Client) UserAccessTokenList() ([]UserAccessToken, error) {
	type TockList struct {
		Tokens      []UserAccessToken `json:"tokens"`
		Permissions []Permission      `json:"permissions"`
	}

	var tockList TockList
	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.Url, "/api/v1/user_access_token"), nil)
	if err != nil {
		return nil, err
	}

	if err = c.sendWithAuth(req, &tockList); err != nil {
		return nil, err
	}

	userAccessTokens := make([]UserAccessToken, len(tockList.Tokens))
	copy(userAccessTokens, tockList.Tokens)
	if len(tockList.Tokens) > 0 {
		userAccessTokens[0].Permissions = tockList.Permissions
	}

	return userAccessTokens, nil
}

func (c *Client) UserAccessTokenCreate(t UserAccessToken) (UserAccessToken, error) {
	var userAccessToken UserAccessToken

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.Url, "/api/v1/user_access_token"), t)
	if err != nil {
		return userAccessToken, err
	}

	if err = c.sendWithAuth(req, &userAccessToken); err != nil {
		return userAccessToken, err
	}

	return userAccessToken, nil
}

func (c *Client) UserAccessTokenDelete(tokenID int) error {

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s", c.Url, fmt.Sprintf("/api/v1/user_access_token/%d", tokenID)), nil)
	if err != nil {
		return err
	}

	if err = c.sendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}
