# Go client for Zammad REST API

This is an Go implementation to lets you use the Zammad's REST API, see https://zammad.com/en/product/features/rest-api
for an introduction.

### Example

```go
package main

import (
    "log"

    "github.com/AlessandroSechi/zammad-go"
)

func main() {
    client := zammad.New("https://my-zammad-instance.com")
    client.Token = "my-accces-token"
    // or basic auth see godoc

    users, err := client.UserList() // Get all users
    if err != nil {
        log.Fatal(err)
    }

    user, err := client.UserShow(1)  // Get User with ID 1
    if err != nil {
        log.Fatal(err)
    }
}
```
