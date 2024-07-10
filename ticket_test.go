package zammad

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"testing"
)

type testClient struct {
	body []byte
}

// Do implements the Doer interface.
func (t testClient) Do(*http.Request) (*http.Response, error) {
	r := &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(t.body)),
		StatusCode: 200,
	}
	return r, nil
}

var ticketTests = []struct {
	File   string // json file to use
	Func   string // function to call
	Expect int    // expected amount of things
}{
	{"ticketlist.json", "TicketList", 2},
	{"ticketsearch.json", "TicketSearch", 3},
}

func TestTicket(t *testing.T) {
	z := &Client{}
	for i, tt := range ticketTests {
		data, _ := os.ReadFile(path.Join("testdata", tt.File))
		z.Client = testClient{body: data}
		t.Run(fmt.Sprintf("%0d-%s", i, tt.Func), func(t *testing.T) {
			var outerr error
			switch tt.Func {
			case "TicketList":
				ts, err := z.TicketList()
				if len(ts) != tt.Expect {
					t.Errorf("expected %d tickets, got %d", tt.Expect, len(ts))
				}
				outerr = err
			case "TicketSearch":
				ts, err := z.TicketSearch("does-not-matter", 0)
				if len(ts) != tt.Expect {
					t.Errorf("expected %d tickets, got %d", tt.Expect, len(ts))
				}
				outerr = err
			}

			if outerr != nil {
				t.Errorf("failed to get %s: %s", tt.Func, outerr)
			}
		})
	}
}
