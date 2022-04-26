package models

type Pagination struct {
	Page          int `json:"page"`
	PageItemCount int `json:"pageItemCount"`
	TotalItems    int `json:"totalItems"`
	NoOfPages     int `json:"numberOfPages"`
}

type VirtualCard struct {
	Id                 string     `json:"id"`
	Status             string     `json:"status"`
	CardHolder         CardHolder `json:"cardholder"`
	Name               string     `json:"displayName"`
	Expires            string     `json:"expires"`
	Currency           string     `json:"currency"`
	LimitCents         int        `json:"limitCents"`
	BalanceCents       int        `json:"balanceCents"`
	SpendCents         int        `json:"spentCents"`
	LifeTimeSpentCents int        `json:"lifetimeSpentCents"`
	Network            string     `json:"network"`
	CompanyName        string     `json:"companyName"`
}

type VirtualCards struct {
	Pagination   Pagination    `json:"pagination"`
	VirtualCards []VirtualCard `json:"virtualCards"`
}
