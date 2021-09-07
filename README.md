# go-pay-sdk
A Golang SDK for Devpay Payment Gateway  Get your API Keys at https://devpay.io

# Install
```go
go mod download https://github.com/DevpayInc/devpay-go-pay-sdk
```

# Usage 

```golang

import "github.com/DevpayInc/devpay-go-pay-sdk/payclient"

client := payclient.New(&payclient.Config{
		Secret:               "SECRTE",
		AccountId:            "AccountId"
})

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

```
> For verbose logs, pass EnableVerboseLogging: true in payclient.Config

# Demo
Please refer example code [here](https://github.com/dev-pay/go-pay-sdk/tree/main/example), follow below steps to run the example code
1. Download the code
2. cd to `go-pay-sdk/example`
3. Run `go mod download github.com/dev-pay/go-pay-sdk/payclient`
4. Update inputs in maing.go 
5. Run the file `go run main.go`

# License
Refer [LICENSE](https://github.com/dev-pay/go-pay-sdk/blob/main/LICENSE) file
