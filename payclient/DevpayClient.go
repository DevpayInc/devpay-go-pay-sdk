package payclient

import (
	"log"
	"os"
)

type DevpayClient struct {
	Config *Config
}

type Config struct {
	Secret               string
	AccountId            string
	Sandbox              bool
	EnableVerboseLogging bool
}

var (

	// API end-points
	DevpayAPI = "https://api.devpay.io"

	// Info logger, used when EnableVerboseLogging set to true
	InfoLogger *log.Logger
)

func New(config *Config) *DevpayClient {
	if config == nil {
		config = &Config{}
		config.Sandbox = false
		config.EnableVerboseLogging = false
	}

	client := &DevpayClient{
		Config: config,
	}
	return client
}

func (devpayClient *DevpayClient) ConfirmPayment(paymntDetail PaymentDetail) (success bool, err error) {

	var InfoLogger *log.Logger = nil
	// Create InfoLogger if verbose logging is enabled
	if devpayClient.Config.EnableVerboseLogging {
		InfoLogger = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	restClient := devpayClient.createRestClient()
	restClient.InfoLogger = InfoLogger

	// Create payment manager
	paymentManager := &PaymentManager{RestAPIClient: restClient,
		Config: devpayClient.Config}

	// Confirm payment
	return paymentManager.ConfirmPayment(paymntDetail)
}

func (devpayClient *DevpayClient) createRestClient() *RestAPIClient {
	var header = make(map[string]string)
	header["Content-Type"] = "application/json"

	// Create Rest API
	restClient := &RestAPIClient{BaseUrl: DevpayAPI, Headers: header}
	return restClient
}
