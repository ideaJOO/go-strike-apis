package gostrikeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//
// Balances : Get account balance details
// https://docs.strike.me/api/get-account-balance-details

type StrikeGetAccountBalanceDetails struct {
	ApiToken string
	Results  []StrikeGetAccountBalanceDetailsResult
	Error    StrikeError
}

type StrikeGetAccountBalanceDetailsResult struct {
	Currency  string `json:"currency"`  // "BTC"
	Current   string `json:"current"`   // "0.00001"
	Pending   string `json:"pending"`   // "0"
	Outgoing  string `json:"outgoing"`  // "0"
	Reserved  string `json:"reserved"`  // "0"
	Available string `json:"available"` // "0.00001"
	Total     string `json:"total"`     // "0.00001"
}

func (tStrikeGetAccountBalanceDetails *StrikeGetAccountBalanceDetails) Get() (err error) {

	if tStrikeGetAccountBalanceDetails.ApiToken == "" {
		err = fmt.Errorf("tStrikeGetAccountBalanceDetails.ApiToken is empty")
		return
	}

	tStrikeGetAccountBalanceDetails.Results = make([]StrikeGetAccountBalanceDetailsResult, 0)

	url := "https://api.strike.me/v1/balances"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		err = fmt.Errorf("@http.NewRequest(method,url,nil): %v", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tStrikeGetAccountBalanceDetails.ApiToken))

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
		err = json.Unmarshal(body, &tStrikeGetAccountBalanceDetails.Results)
	} else {
		err = json.Unmarshal(body, &tStrikeGetAccountBalanceDetails.Error)
		fmt.Printf("\n\n%s\n", string(body))
	}
	if err != nil {
		err = fmt.Errorf("@json.Unmarshal(body,): %v", err)
		return
	}

	return
}
