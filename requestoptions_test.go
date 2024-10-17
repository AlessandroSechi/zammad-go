package zammad

import (
	"net/url"
	"testing"
)

func TestRequestOptions(t *testing.T) {
	var requestOptionsTests = []struct {
		opts   []Option
		expect string
	}{
		{[]Option{WithPage(2), WithPerPage(50)}, "page=2&per_page=50"},
		{[]Option{WithSortBy("name"), WithOrderBy("asc")}, "order_by=asc&sort_by=name"},
		{[]Option{WithPage(2), WithPerPage(50), WithSortBy("date"), WithOrderBy("desc")}, "order_by=desc&page=2&per_page=50&sort_by=date"},
	}

	for _, tt := range requestOptionsTests {
		t.Run(url.QueryEscape(tt.expect), func(t *testing.T) {
			ro := NewRequestOptions(tt.opts...)
			params := ro.URLParams()
			if params != tt.expect {
				t.Errorf("expected %s, got %s", tt.expect, params)
			}
		})
	}
}
