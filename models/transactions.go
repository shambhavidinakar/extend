package models

import "time"

type Transaction struct {
	Id                         string    `json:"id"`
	CardHolderName             string    `json:"cardholderName"`
	VCDisplayName              string    `json:"vcnDisplayName"`
	VirtualCardId              string    `json:"virtualCardId"`
	Type                       string    `json:"type"`
	Status                     string    `json:"status"`
	ApprovalCode               string    `json:"approvalCode"`
	AuthBillingAmountCents     int       `json:"authBillingAmountCents"`
	AuthBillingCurrency        string    `json:"authBillingCurrency"`
	ClearingBillingAmountCents int       `json:"clearingBillingAmountCents"`
	ClearingBillingCurrency    string    `json:"clearingBillingCurrency"`
	MccGroup                   string    `json:"mccGroup"`
	MerchentId                 string    `json:"merchantId"`
	MerchentName               string    `json:"merchantName"`
	AuthedAt                   time.Time `json:"authedAt"`
	ClearedAt                  time.Time `json:"clearedAt"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}
