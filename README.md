# Go client for Zammad REST API

This is an Go implementation to lets you use the Zammad's REST API, see https://zammad.com/en/product/features/rest-api
for an introduction. The documentation in the code has further links to details of the Zammad API.

## Example

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

	user, err := client.UserShow(1) // Get User with ID 1
	if err != nil {
		log.Fatal(err)
	}

	// to iterate over pages
	ticketRes := client.TicketListResult(zammad.WithPerPage(50))
	for ticketRes.Next() {
		tickets, err := ticketRes.Fetch()
		if err != nil {
			log.Fatalf("failed: %s", err)
		}

		log.Printf("number of tickets: %d\n", len(tickets))
	}
}
```

## Development

If you want to test the code against a live zammad instance, you'll need to create an access token
and set the following environment variables before testing:

* `ZAMMAD_INSTANCE`: this is the URL of you zammad instance
* `ZAMMAD_TOKEN`: this the access token you can use
