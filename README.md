### Go client for Zammad REST API


### Example

```go
import "github.com/AlessandroSechi/zammad-go"

// Create a client instance
client, err := zammad.NewClient(&zammad.Client{
		Username: "",
		Password: "",
		Token:    "",
		OAuth:    "",
		Url:      "http://my-zammad-instance.com",
	})

users, err := client.UserList() //Get Users

user, err := client.UserShow(1) //Get User
```