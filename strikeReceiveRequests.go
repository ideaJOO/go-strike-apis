package gostrikeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//
// Receive Requests : Create a receive request
// https://docs.strike.me/api/create-a-receive-request

type StrikeCreateReceiveRequest struct {
	ApiToken        string
	Amount          StrikeAmount
	Description     string
	ExpiryInSeconds int
	Result          StrikeCreateReceiveRequestResult
	Error           StrikeError
}

type StrikeCreateReceiveRequestResult struct {
	ReceiveRequestID string                                 `json:"receiveRequestId"`
	Created          string                                 `json:"created"`
	TargetCurrency   string                                 `json:"targetCurrency"`
	Bolt11           StrikeCreateReceiveRequestResultBolt11 `json:"bolt11"`
}

type StrikeCreateReceiveRequestResultBolt11 struct {
	Invoice         string       `json:"invoice"`         // "lnbc10n1pnkel25pp5g3lldp..."
	RequestedAmount StrikeAmount `json:"requestedAmount"` // StrikeAmount
	BtcAmount       string       `json:"btcAmount"`       // "0.00000001"
	Description     string       `json:"description"`     // "For Pizza"
	PaymentHash     string       `json:"paymentHash"`     // "447ff684b9a72e162c89..."
	Expires         string       `json:"expires"`         // "2024-12-26T06:54:08.8567891+00:00"
}

func (tStrikeCreateReceiveRequest *StrikeCreateReceiveRequest) Post() (err error) {

	if tStrikeCreateReceiveRequest.Amount.Amount == "" || tStrikeCreateReceiveRequest.Amount.Currency != "BTC" { // Only BTC
		err = fmt.Errorf("incorrect tStrikeCreateReceiveRequest.Amount")
		return
	}

	if tStrikeCreateReceiveRequest.ExpiryInSeconds == 0 {
		tStrikeCreateReceiveRequest.ExpiryInSeconds = 60 // Default
	}

	url := "https://api.strike.me/v1/receive-requests"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
	"bolt11": {
		"amount": {
			"amount": "%s",
			"currency": "%s"
		},
		"description": "%s",
		"expiryInSeconds": %d
		},
		"targetCurrency": "%s"
  	}`, tStrikeCreateReceiveRequest.Amount.Amount, tStrikeCreateReceiveRequest.Amount.Currency, tStrikeCreateReceiveRequest.Description, tStrikeCreateReceiveRequest.ExpiryInSeconds, tStrikeCreateReceiveRequest.Amount.Currency))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		err = fmt.Errorf("@http.NewRequest(method,url,payload): %v", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tStrikeCreateReceiveRequest.ApiToken))

	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("@client.Do(req): %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {

	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("@io.ReadAll(res.Body): %v", err)
		return
	}

	if res.StatusCode == 201 {
		err = json.Unmarshal(body, &tStrikeCreateReceiveRequest.Result)
	} else {
		err = json.Unmarshal(body, &tStrikeCreateReceiveRequest.Error)
	}
	if err != nil {
		err = fmt.Errorf("@json.Unmarshal(body,): %v", err)
		return
	}

	return
}

// 03.
// Receive Requests : Get receives for receive request
// https://docs.strike.me/api/get-receives-for-receive-request

type StrikeGetReceivesForReceiveRequest struct {
	ApiToken         string
	ReceiveRequestId string
	Result           StrikeGetReceivesForReceiveRequestResult
	Error            StrikeError
}

type StrikeGetReceivesForReceiveRequestResult struct {
	Items []StrikeGetReceivesForReceiveRequestResultItem `json:"items"`
	Count int                                            `json:"count"`
}

type StrikeGetReceivesForReceiveRequestResultItem struct {
	ReceiveId        string                                          `json:"receiveId"`        //  "..."
	ReceiveRequestId string                                          `json:"receiveRequestId"` //  "..."
	Type             string                                          `json:"type"`             // "P2P"
	State            string                                          `json:"state"`            //  "COMPLETED" "PENDING"
	AmountReceived   StrikeAmount                                    `json:"amountReceived"`   // StrikeAmount
	AmountCredited   StrikeAmount                                    `json:"amountCredited"`   // StrikeAmount
	Created          string                                          `json:"created"`          // "2024-12-26T09:45:26.414205+00:00"
	Completed        string                                          `json:"completed"`        // "2024-12-26T09:45:26.414205+00:00"
	P2p              StrikeGetReceivesForReceiveRequestResultItemP2p `json:"p2p"`              // StrikeReceiveRequestCheckResultItemP2p
}

type StrikeGetReceivesForReceiveRequestResultItemP2p struct {
	PayerAccountId string `json:"payerAccountId"` // "..."
}

func (tStrikeGetReceivesForReceiveRequest *StrikeGetReceivesForReceiveRequest) Get() (err error) {

	if tStrikeGetReceivesForReceiveRequest.ReceiveRequestId == "" {
		err = fmt.Errorf("incorrect tStrikeGetReceivesForReceiveRequest.ReceiveRequestId")
		return
	}

	url := "https://api.strike.me/v1/receive-requests/01940259-98a4-7bb5-9acd-1c5a5f0c946f/receives"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		err = fmt.Errorf("@http.NewRequest(method,url,payload): %v", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tStrikeGetReceivesForReceiveRequest.ApiToken))
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
		err = json.Unmarshal(body, &tStrikeGetReceivesForReceiveRequest.Result)
	} else {
		err = json.Unmarshal(body, &tStrikeGetReceivesForReceiveRequest.Error)
	}
	if err != nil {
		err = fmt.Errorf("@json.Unmarshal(body,): %v", err)
		return
	}

	return
}
