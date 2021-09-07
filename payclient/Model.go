package payclient

type PaymentDetail struct {
	Amount         int
	Card           Card
	BillingAddress BillingAddress
	MetaData       map[string]string
	Name           string
	Description    string
}

type Card struct {
	CardNum         string
	CardExpiryMonth string
	CardExpiryYear  string
	Cvv             string
}

type BillingAddress struct {
	Country string `json:"country"`
	Zip     string `json:"zip"`
	State   string `json:"state"`
	Street  string `json:"street"`
	City    string `json:"city"`
}
