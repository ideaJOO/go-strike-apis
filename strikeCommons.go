package gostrikeapis

//
// Common

type StrikeError struct {
	TraceId string          `json:"traceId"`
	Data    StrikeErrorData `json:"data"`
}

type StrikeErrorData struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type StrikeAmount struct {
	Amount   string `json:"amount"`   // "0.00001"
	Currency string `json:"currency"` // "BTC"
}
