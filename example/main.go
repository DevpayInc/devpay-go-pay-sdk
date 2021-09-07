package main

import (
	"fmt"

	"github.com/dev-pay/go-pay-sdk/payclient"
)

func main() {
	client := payclient.New(&payclient.Config{
		Secret:               "SECRET",
		AccountId:            "AccountId",
		Sandbox:              true,
		EnableVerboseLogging: true})

	success, err := client.ConfirmPayment(payclient.PaymentDetail{
		Amount: 103,
		Card: payclient.Card{CardNum: "4242424242424242",
			CardExpiryMonth: "10",
			CardExpiryYear:  "2024",
			Cvv:             "102"},
		BillingAddress: payclient.BillingAddress{Country: "US",
			Zip:    "38138",
			State:  "TN",
			City:   "Memphis",
			Street: "123 ABC Lane"},
		Name: "John",
	})

	if err != nil {
		fmt.Println("Error - " + err.Error())
	} else if success {
		fmt.Printf("Payment successful")
	}

}
