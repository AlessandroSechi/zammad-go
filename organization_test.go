package zammad

import (
	"fmt"
	"os"
	"path"
	"testing"
)

var organizationTests = []struct {
	File   string // json file to use
	Func   string // function to call
	Expect int    // expected amount of things
}{
	{"organizationlist.json", "OrganizationList", 1},
}

func TestOrganization(t *testing.T) {
	z := &Client{}
	for i, tt := range organizationTests {
		data, _ := os.ReadFile(path.Join("testdata", tt.File))
		z.Client = testClient{body: data}
		t.Run(fmt.Sprintf("%0d-%s", i, tt.Func), func(t *testing.T) {
			var outerr error
			switch tt.Func {
			case "OrganizationList":
				ts, err := z.OrganizationList()
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
