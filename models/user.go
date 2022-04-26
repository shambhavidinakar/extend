package models

type CardHolder struct {
	Id               string `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	OrganizationID   string `json:"organizationId"`
	OrganizationRole string `json:"organizationRole"`
}
