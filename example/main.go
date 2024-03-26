package main

import (
	"context"
	"log"

	"github.com/fairytale5571/go-whitepay"
)

func main() {
	wp := whitepay.New("123", "slug")
	ctx := context.Background()
	var response, err = wp.CreateNewOrder(ctx, &whitepay.CreateNewOrderRequest{
		Amount:          "10000",
		Currency:        "USD",
		ExternalOrderID: "123456",
		SuccessfulLink:  "https://google.com",
		FailureLink:     "https://bing.com",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Response: %+v", response)
}
