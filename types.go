package zammad

import (
	"fmt"
	"net/http"
)

type (
	// Client is used to query Zammad. It is safe to use concurrently.
	Client struct {
		Client   *http.Client // Client is the http client used to make the queries.
		Username string       // Username and Password are used when doing basic auth.
		Password string       // Password used when doing basic auth
		Token    string       // Token is used when using an Access Token
		OAuth    string
		Url      string
		// FromFunc is used to set the From HTTP header, if you want to act on behalf of another user.
		// See https://docs.zammad.org/en/latest/api/intro.html#actions-on-behalf-of-other-users. If not nil
		// *and* returning a non empty string, this value will be used in the request.
		FromFunc func() string
	}

	// TODO: not used yet.
	ErrorResponse struct {
		Description      string `json:"error"`
		DescriptionHuman string `json:"error_human"`
	}
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprint(r.Description)
}
