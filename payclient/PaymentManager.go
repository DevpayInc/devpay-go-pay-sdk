package payclient

import (
	"encoding/json"
	"errors"
)

type PaymentManager struct {
	RestAPIClient *RestAPIClient
	Config        *Config
}

func (paymentManager *PaymentManager) ConfirmPayment(paymentDetail PaymentDetail) (bool, error) {

	token, err := paymentManager.createPaymentIntent(paymentDetail)
	if err != nil {
		return false, err
	}
	result, err := paymentManager.confirmIntent(token, paymentDetail)
	return result, err
}

func createDevpayRestClient() *RestAPIClient {
	var header = make(map[string]string)
	header["Content-Type"] = "application/json"

	// Select end point
	var endPoint = "https://api.devpay.io"

	// Create Rest API
	restClient := &RestAPIClient{BaseUrl: endPoint, Headers: header}
	return restClient
}

func (manager *PaymentManager) createPaymentIntent(paymentDetail PaymentDetail) (string, error) {

	cardDict := make(map[string]interface{})
	cardDict["Number"] = paymentDetail.Card.CardNum
	cardDict["ExpMonth"] = paymentDetail.Card.CardExpiryMonth
	cardDict["ExpYear"] = paymentDetail.Card.CardExpiryYear
	cardDict["Cvc"] = paymentDetail.Card.Cvv

	chargeDetails := make(map[string]interface{})
	chargeDetails["amount"] = paymentDetail.Amount
	chargeDetails["fee_amount"] = 0
	chargeDetails["description"] = paymentDetail.Description
	chargeDetails["account_id"] = manager.Config.AccountId
	chargeDetails["secreate_key"] = manager.Config.Secret

	if manager.Config.Sandbox {
		chargeDetails["environment"] = "sandbox"
	}

	payload := map[string]interface{}{
		"CardDetails":   cardDict,
		"ChargeDetails": chargeDetails,
	}

	b, _ := json.Marshal(payload)

	var header = make(map[string]string)
	response, err := manager.RestAPIClient.Post("/v1/charge/create_Intend", b, header)
	if err != nil {
		return "", err
	}

	var mappedData map[string]interface{}
	err = json.Unmarshal([]byte(response), &mappedData)
	if err != nil {
		return "", err
	}

	token, ok := mappedData["token"].(string)
	if !ok {
		return "", errors.New("failed to process the data")
	}
	return token, nil
}

func (manager *PaymentManager) confirmIntent(token string, paymentDetail PaymentDetail) (bool, error) {

	cardDict := make(map[string]interface{})
	cardDict["Number"] = paymentDetail.Card.CardNum
	cardDict["ExpMonth"] = paymentDetail.Card.CardExpiryMonth
	cardDict["ExpYear"] = paymentDetail.Card.CardExpiryYear
	cardDict["Cvc"] = paymentDetail.Card.Cvv
	cardDict["token"] = token

	chargeDetails := make(map[string]interface{})
	chargeDetails["amount"] = paymentDetail.Amount
	chargeDetails["fee_amount"] = 0
	chargeDetails["description"] = paymentDetail.Description
	chargeDetails["account_id"] = manager.Config.AccountId
	chargeDetails["secreate_key"] = manager.Config.Secret

	if manager.Config.Sandbox {
		chargeDetails["environment"] = "sandbox"
	}

	payload := map[string]interface{}{
		"CardDetails":   cardDict,
		"ChargeDetails": chargeDetails,
	}

	b, _ := json.Marshal(payload)

	var header = make(map[string]string)
	response, err := manager.RestAPIClient.Post("/v1/charge/confirm_charge", b, header)
	if err != nil {
		return false, err
	}

	var mappedData map[string]interface{}
	err = json.Unmarshal([]byte(response), &mappedData)
	if err != nil {
		return false, err
	}

	result, ok := mappedData["status"]
	if !ok {
		return false, errors.New("failed to process the data")
	}
	return (result.(float64) == 1), nil
}
