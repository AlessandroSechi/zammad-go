package zammad

import (
	"net/url"
	"strconv"
)

// RequestOptions holds the parameters for making a request.
type RequestOptions struct {
	page    int    // Page number for pagination
	perPage int    // Number of items per page
	sortBy  string // Field to sort by
	orderBy string // Order of sorting (e.g., asc, desc)
}

// NewRequestOptions creates a new RequestOptions instance with the provided options.
// It applies each Option function to the RequestOptions instance.
func NewRequestOptions(opts ...Option) RequestOptions {
	// set default for page/perPage as enforced by the API
	ro := RequestOptions{}

	for _, opt := range opts {
		opt(&ro)
	}

	return ro
}

// URLParams returns the URL-encoded parameters for the request.
func (ro RequestOptions) URLParams() string {
	params := url.Values{}

	if ro.page > 0 {
		params.Set("page", strconv.Itoa(ro.page))
	}

	if ro.perPage > 0 {
		params.Set("per_page", strconv.Itoa(ro.perPage))
	}

	if ro.sortBy != "" {
		params.Set("sort_by", ro.sortBy)
	}

	if ro.orderBy != "" {
		params.Set("order_by", ro.orderBy)
	}

	return params.Encode()
}

// Option is a function that modifies a RequestOptions instance.
type Option func(ro *RequestOptions)

// WithPage sets the page number in RequestOptions.
func WithPage(page int) Option {
	return func(ro *RequestOptions) {
		ro.page = page
	}
}

// WithPerPage sets the number of items per page in RequestOptions.
func WithPerPage(perPage int) Option {
	return func(ro *RequestOptions) {
		ro.perPage = perPage
	}
}

// WithOrderBy sets the order of sorting in RequestOptions.
func WithOrderBy(orderBy string) Option {
	return func(ro *RequestOptions) {
		ro.orderBy = orderBy
	}
}

// WithSortBy sets the field to sort by in RequestOptions.
func WithSortBy(sortBy string) Option {
	return func(ro *RequestOptions) {
		ro.sortBy = sortBy
	}
}
