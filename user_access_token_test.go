package zammad

import (
	"fmt"
	"os"
	"path"
	"testing"
)

var userAccessTokenTests = []struct {
	File       string // json file to use
	Func       string // function to call
	Expect     int    // expected amount of things
	ExpectPerm int    // expected amount of permissions that List returns (0 = not used)
}{
	{"user_access_tokenlist.json", "UserAccessTokenList", 1, 68},
}

func TestUserAccessToken(t *testing.T) {
	z := &Client{}
	for i, tt := range userAccessTokenTests {
		data, _ := os.ReadFile(path.Join("testdata", tt.File))
		z.Client = testClient{body: data}
		t.Run(fmt.Sprintf("%0d-%s", i, tt.Func), func(t *testing.T) {
			var outerr error
			switch tt.Func {
			case "UserAccessTokenList":
				ts, err := z.UserAccessTokenList()
				if len(ts) != tt.Expect {
					t.Errorf("expected %d tokens, got %d", tt.Expect, len(ts))
				}
				if tt.ExpectPerm > 0 {
					if x := len(ts[0].Permissions); x != tt.ExpectPerm {
						t.Errorf("expected %d permissions, got %d", tt.ExpectPerm, x)
					}
				}
				outerr = err
			}

			if outerr != nil {
				t.Errorf("failed to get %s: %s", tt.Func, outerr)
			}
		})
	}
}
