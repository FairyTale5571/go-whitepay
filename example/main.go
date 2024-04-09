package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/fairytale5571/go-whitepay"
)

const (
	apiKey                 = "your_api"
	slugName               = "your_slug_name"
	slugTokenSignature     = "your_slug_token"
	merchantTokenSignature = "your_merchant_token"
)

func listener(wp *whitepay.WhitePay) {
	app := gin.Default()
	app.POST("/webhook", gin.WrapH(wp.WebhookHandler()))
	app.POST("/", gin.WrapH(wp.WebhookHandler()))
	_ = app.Run(":8080")
}

func main() {
	wp := whitepay.New(
		apiKey,
		slugName,
		whitepay.WithChannelSize(1024),
		whitepay.WithSlugSignatureToken(slugTokenSignature),
		whitepay.WithMerchantSignatureToken(merchantTokenSignature),
	)

	go listener(wp)

	ctx := context.Background()
	response, err := wp.CreateNewOrder(ctx, &whitepay.CreateNewOrderRequest{
		Amount:          "100",
		Currency:        "USDT",
		ExternalOrderID: "123456",
		SuccessfulLink:  "https://site.com/success",
		FailureLink:     "https://site.com/failure",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Response: %+v", response)

	for {
		select {
		case event := <-wp.Updates():
			if event.EventType == whitepay.EventOrderStatusChanged {

			}
		}
	}
}
