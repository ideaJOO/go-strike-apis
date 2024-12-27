package gostrikeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//
// Payments : Create LNURL/LN address payment quote
// https://docs.strike.me/api/create-lnurl-ln-address-payment-quote

type StrikeCreateLnurlLnAddressPaymentQuote struct {
	ApiToken       string
	LnAddressOrUrl string
	Amount         StrikeAmount
	Description    string
	Result         StrikeCreateLnurlLnAddressPaymentQuoteResult
	Error          StrikeError
}

type StrikeCreateLnurlLnAddressPaymentQuoteResult struct {
	PaymentQuoteID string       `json:"paymentQuoteId"`
	Description    string       `json:"description"`
	Amount         StrikeAmount `json:"amount"`
	TotalAmount    StrikeAmount `json:"totalAmount"`
}

func (tStrikeCreateLnurlLnAddressPaymentQuote *StrikeCreateLnurlLnAddressPaymentQuote) Post() (err error) {

	if tStrikeCreateLnurlLnAddressPaymentQuote.LnAddressOrUrl == "" {
		err = fmt.Errorf("tStrikeCreateLnurlLnAddressPaymentQuote.LnAddressOrUrl is empty")
		return
	}
	if tStrikeCreateLnurlLnAddressPaymentQuote.Amount.Amount == "" || tStrikeCreateLnurlLnAddressPaymentQuote.Amount.Currency != "BTC" {
		err = fmt.Errorf("incorrect tStrikeCreateLnurlLnAddressPaymentQuote.Amount")
		return
	}
	if tStrikeCreateLnurlLnAddressPaymentQuote.Description == "" {
		err = fmt.Errorf("tStrikeCreateLnurlLnAddressPaymentQuote.Description is empty")
		return
	}

	url := "https://api.strike.me/v1/payment-quotes/lightning/lnurl"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
		"lnAddressOrUrl": "%s",
		"sourceCurrency": "%s",
		"amount": {
			"amount": "%s",
			"currency": "%s"
			},
			"description": "%s"
			}`, tStrikeCreateLnurlLnAddressPaymentQuote.LnAddressOrUrl, tStrikeCreateLnurlLnAddressPaymentQuote.Amount.Currency, tStrikeCreateLnurlLnAddressPaymentQuote.Amount.Amount, tStrikeCreateLnurlLnAddressPaymentQuote.Amount.Currency, tStrikeCreateLnurlLnAddressPaymentQuote.Description))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		err = fmt.Errorf("@http.NewRequest(method,url,payload): %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tStrikeCreateLnurlLnAddressPaymentQuote.ApiToken))

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("@client.Do(req): %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("@io.ReadAll(res.Body): %v", err)
		return
	}

	if res.StatusCode == 200 {
		err = json.Unmarshal(body, &tStrikeCreateLnurlLnAddressPaymentQuote.Result)
	} else {
		err = json.Unmarshal(body, &tStrikeCreateLnurlLnAddressPaymentQuote.Error)
	}
	if err != nil {
		err = fmt.Errorf("@json.Unmarshal(body,): %v", err)
		return
	}

	return
}

//
// Payments : Execute payment quote
// https://docs.strike.me/api/execute-payment-quote

type StrikeExecutePaymentQuote struct {
	ApiToken       string
	PaymentQuoteID string
	Result         StrikeExecutePaymentQuoteResult
	Error          StrikeError
}
type StrikeExecutePaymentQuoteResult struct {
	PaymentId   string       `json:"paymentId"`   // paymentId
	State       string       `json:"state"`       // PENDING, COMPLETED, FAILE
	Result      string       `json:"result"`      // PENDING, SUCCESS, FAILURE
	Completed   string       `json:"completed"`   // "2024-12-27T04:50:46.2870744+00:00"
	Delivered   string       `json:"delivered"`   // "2024-12-27T04:50:46.2870744+00:00"
	Amount      StrikeAmount `json:"amount"`      // amount
	TotalAmount StrikeAmount `json:"totalAmount"` // amount
}

func (tStrikeExecutePaymentQuote *StrikeExecutePaymentQuote) Fetch() (err error) {

	if tStrikeExecutePaymentQuote.PaymentQuoteID == "" {
		err = fmt.Errorf("tStrikeExecutePaymentQuote.PaymentQuoteID is empty")
		return
	}

	url := fmt.Sprintf("https://api.strike.me/v1/payment-quotes/%s/execute", tStrikeExecutePaymentQuote.PaymentQuoteID)
	method := "PATCH"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		err = fmt.Errorf("@http.NewRequest(method,url,nil): %v", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tStrikeExecutePaymentQuote.ApiToken))

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("@client.Do(req): %v", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("@io.ReadAll(res.Body): %v", err)
		return
	}

	if res.StatusCode == 200 || res.StatusCode == 202 {
		err = json.Unmarshal(body, &tStrikeExecutePaymentQuote.Result)
	} else {
		err = json.Unmarshal(body, &tStrikeExecutePaymentQuote.Error)
	}
	if err != nil {
		err = fmt.Errorf("@json.Unmarshal(body,): %v", err)
		return
	}

	return
}
