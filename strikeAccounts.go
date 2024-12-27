package gostrikeapis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//
// Accounts : Fetch public account profile info by ID
// https://docs.strike.me/api/fetch-public-account-profile-info-by-id

type StrikeFetchPublicAccountProfileInfoByID struct {
	ApiToken  string
	AccountId string
	Result    StrikeFetchPublicAccountProfileInfoByIdResult
	Error     StrikeError
}

type StrikeFetchPublicAccountProfileInfoByIdResult struct {
	AccountID  string                                                  `json:"id"`
	Handle     string                                                  `json:"handle"`
	CanReceive bool                                                    `json:"canReceive"`
	Currencies []StrikeFetchPublicAccountProfileInfoByIdResultCurrency `json:"currencies"`
}

type StrikeFetchPublicAccountProfileInfoByIdResultCurrency struct {
	Currency          string `json:"currency"`
	IsDefaultCurrency bool   `json:"isDefaultCurrency"`
	IsAvailable       bool   `json:"isAvailable"`
	IsInvoiceable     bool   `json:"isInvoiceable"`
}

func (tStrikeFetchPublicAccountProfileInfoByID *StrikeFetchPublicAccountProfileInfoByID) Get() (err error) {

	if tStrikeFetchPublicAccountProfileInfoByID.AccountId == "" {
		err = fmt.Errorf("tStrikeFetchPublicAccountProfileInfoByID.AccountId is empty")
		return
	}

	url := fmt.Sprintf("https://api.strike.me/v1/accounts/%s/profile", tStrikeFetchPublicAccountProfileInfoByID.AccountId)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		err = fmt.Errorf("@http.NewRequest(method,url,nil): %v", err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", tStrikeFetchPublicAccountProfileInfoByID.ApiToken))

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
		err = json.Unmarshal(body, &tStrikeFetchPublicAccountProfileInfoByID.Result)
	} else {
		err = json.Unmarshal(body, &tStrikeFetchPublicAccountProfileInfoByID.Error)
	}
	if err != nil {
		err = fmt.Errorf("@json.Unmarshal(body,): %v", err)
		return
	}

	return
}
