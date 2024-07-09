package zammad

import (
	"fmt"
	"net/http"
)

type (
	// Client is used to query Zammad. It is safe to use concurrently.
	Client struct {
		Client   Doer
		Username string
		Password string
		Token    string
		OAuth    string
		Url      string // Url is the URL of Zammad.
	}

	// TODO: not used yet.
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
