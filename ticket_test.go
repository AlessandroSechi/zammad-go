package zammad

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"testing"
)

// testZammad returns a client that is configured via the enviroment vars: ZAMMAD_INSTANCE and ZAMMAD_TOKEN.
// If one of them is not set, nil is returned.
func testZammad() *Client {
	inst := os.Getenv("ZAMMAD_INSTANCE")
	tok := os.Getenv("ZAMMAD_TOKEN")

	if inst == "" || tok == "" {
		return nil
	}
	client := New(inst)
	client.Token = tok
	return client
}

type testClient struct {
	pages int
	body  []byte
}

// Do implements the Doer interface.
func (t testClient) Do(r *http.Request) (*http.Response, error) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page == 0 {
		page = 1
	}

	if page > t.pages {
		return &http.Response{
			Body:       io.NopCloser(bytes.NewBufferString("[]")),
			StatusCode: 200,
		}, nil
	}

	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(t.body)),
		StatusCode: 200,
	}, nil
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
		z.Client = testClient{body: data, pages: 1}
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

func TestTicketCreate(t *testing.T) {
	client := testZammad()
	if client == nil {
		t.SkipNow()
	}

	// these values might not map to your zammad instance.
	ticket := Ticket{
		Title:    "test: test",
		Group:    "Sysadmin",
		Customer: "bram",
		Article: TicketArticle{
			Subject:  "Hello",
			Body:     "Hello test",
			Type:     "note",
			Internal: true,
		},
	}

	ticket1, err := client.TicketCreate(ticket)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", ticket1)
}
