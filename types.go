package zammad

import (
	"fmt"
	"net/http"
)

type (
	// Client is used to query Zammad. It is safe to use concurrently. If you (inadvertly) added
	// multiple authencation options that will be applied in the order, basic auth, token based, and
	// then oauth. Where the last one set, wins.
	Client struct {
		Client   Doer
		Username string
		Password string
		Token    string
		OAuth    string
		Url      string
		// FromFunc is used to set the From HTTP header, if you want to act on behalf of another user.
		// See https://docs.zammad.org/en/latest/api/intro.html#actions-on-behalf-of-other-users. If not nil
		// *and* returning a non empty string, this value will be used in the request.
		FromFunc func() string
	}

	// ErrorResponse is the response returned by Zammad when an error occured.
	ErrorResponse struct {
		Description      string `json:"error"`
		DescriptionHuman string `json:"error_human"`
	}

	// Doer is an interface that allows mimicking a *http.Client.
	Doer interface {
		Do(*http.Request) (*http.Response, error)
	}
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprint(r.Description)
}
