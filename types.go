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
		Client   *http.Client // Client is the http client used to make the queries.
		Username string       // Username and Password are used when doing basic auth.
		Password string       // Password used when doing basic auth.
		Token    string       // Token is used when using an Access Token.
		OAuth    string       // Oauth is used when using Oauth authentication.
		Url      string       // Url is the URL of Zammad.
	}

	ErrorResponse struct {
		Description      string `json:"error"`
		DescriptionHuman string `json:"error_human"`
	}
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprint(r.Description)
}
