package gostrikeapis

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	godotenvlight "github.com/ideajoo/go-dotenv-light"
)

var StrikeApiToken string

func initEnv() {
	godotenvlight.Export(false, "./.env")
	StrikeApiToken = os.Getenv("STRIKE_API_TOKEN") // or YourToken
}

func TestStrikeGetAccountBalanceDetails(t *testing.T) {
	initEnv()
	tStrikeGetAccountBalanceDetails := StrikeGetAccountBalanceDetails{}
	tStrikeGetAccountBalanceDetails.ApiToken = StrikeApiToken
	err := tStrikeGetAccountBalanceDetails.Get()
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}
	tByte, _ := json.MarshalIndent(tStrikeGetAccountBalanceDetails, "", "  ")
	fmt.Printf("\n%s\n", tByte)
}

func TestStrikeReceiveRequest(t *testing.T) {
	tStrikeCreateReceiveRequest := StrikeCreateReceiveRequest{}
	tStrikeCreateReceiveRequest.ApiToken = StrikeApiToken
	tStrikeCreateReceiveRequest.Amount.Amount = "0.00000001"
	tStrikeCreateReceiveRequest.Amount.Currency = "BTC"
	tStrikeCreateReceiveRequest.Description = "For Pizza"
	tStrikeCreateReceiveRequest.ExpiryInSeconds = 1200

	err := tStrikeCreateReceiveRequest.Post()
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}
	tByte, _ := json.MarshalIndent(tStrikeCreateReceiveRequest, "", "  ")
	fmt.Printf("\n%s\n", tByte)
}

func TestStrikeGetReceivesForReceiveRequest(t *testing.T) {
	tStrikeGetReceivesForReceiveRequest := StrikeGetReceivesForReceiveRequest{}
	tStrikeGetReceivesForReceiveRequest.ApiToken = StrikeApiToken
	tStrikeGetReceivesForReceiveRequest.ReceiveRequestId = "fb85f340-0e33-4e91-9600-8fbc43b4123e" // "satoshibento"

	err := tStrikeGetReceivesForReceiveRequest.Get()
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}
	tByte, _ := json.MarshalIndent(tStrikeGetReceivesForReceiveRequest, "", "  ")
	fmt.Printf("\n%s\n", tByte)
}

func TestStrikeFetchPublicAccountProfileInfoByID(t *testing.T) {
	tStrikeFetchPublicAccountProfileInfoByID := StrikeFetchPublicAccountProfileInfoByID{}
	tStrikeFetchPublicAccountProfileInfoByID.ApiToken = StrikeApiToken
	tStrikeFetchPublicAccountProfileInfoByID.AccountId = "fb85f340-0e33-4e91-9600-8fbc43b4123e" // "satoshibento"
	err := tStrikeFetchPublicAccountProfileInfoByID.Get()
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}
	tByte, _ := json.MarshalIndent(tStrikeFetchPublicAccountProfileInfoByID, "", "  ")
	fmt.Printf("\n%s\n", tByte)
}

func TestStrikeCreateLnurlLnAddressPaymentQuote(t *testing.T) {
	tStrikeCreateLnurlLnAddressPaymentQuote := StrikeCreateLnurlLnAddressPaymentQuote{}
	tStrikeCreateLnurlLnAddressPaymentQuote.ApiToken = StrikeApiToken
	tStrikeCreateLnurlLnAddressPaymentQuote.LnAddressOrUrl = "satoshibento@strike.me" // "satoshibento"
	tStrikeCreateLnurlLnAddressPaymentQuote.Amount.Amount = "0.00000002"
	tStrikeCreateLnurlLnAddressPaymentQuote.Amount.Currency = "BTC"
	tStrikeCreateLnurlLnAddressPaymentQuote.Description = "For Bagel TypeB"
	err := tStrikeCreateLnurlLnAddressPaymentQuote.Post()
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}
	tByte, _ := json.MarshalIndent(tStrikeCreateLnurlLnAddressPaymentQuote, "", "  ")
	fmt.Printf("\n%s\n", tByte)
}

func TestStrikeExecutePaymentQuote(t *testing.T) {

	tStrikeExecutePaymentQuote := StrikeExecutePaymentQuote{}
	tStrikeExecutePaymentQuote.PaymentQuoteID = "35d647a3-c61b-4f3d-a876-46111b84fff0"
	tStrikeExecutePaymentQuote.ApiToken = StrikeApiToken
	err := tStrikeExecutePaymentQuote.Fetch()
	if err != nil {
		fmt.Printf("\n%v\n", err)
		return
	}
	tByte, _ := json.MarshalIndent(tStrikeExecutePaymentQuote, "", "  ")
	fmt.Printf("\n%s\n", tByte)
}
