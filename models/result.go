package models

type ResultLite struct {
	Card         VirtualCard   `json:"virtualCards"`
	Transactions []Transaction `json:"transactions"`
}

type Result struct {
	Card         interface{}   `json:"virtualCards"`
	Transactions []interface{} `json:"transactions"`
}
