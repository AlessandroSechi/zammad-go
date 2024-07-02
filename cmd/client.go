package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/AlessandroSechi/zammad-go"
)

func main() {
	client, _ := zammad.NewClient(&zammad.Client{
		Username: "",
		Password: "",
		Token:    "63v2JOQ3_EbBQ0_DpAiA-n2Fuaa7dymCvT4e29fXiPNyrYFu-ItN1IosVI2nNEr8",
		OAuth:    "",
		Url:      "https://helpdesk.science.ru.nl",
	})

	q := url.QueryEscape("state.name:(new OR open)")
	tickets, err := client.TicketSearch(q, 1000000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", tickets)
}
