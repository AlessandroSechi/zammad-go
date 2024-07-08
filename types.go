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
		Url      string // Url is the URL of Zammad.
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
