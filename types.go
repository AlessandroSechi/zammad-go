package zammad

import (
	"fmt"
	"net/http"
)

type (
	Client struct {
		Client   *http.Client
		Username string
		Password string
		Token    string
		OAuth    string
		Url      string
	}

	ErrorResponse struct {
		Description      string `json:"error"`
		DescriptionHuman string `json:"error_human"`
	}
)

func (r *ErrorResponse) Error() string {
	return fmt.Sprint(r.Description)
}
