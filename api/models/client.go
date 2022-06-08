package models

type Client struct {
	Id          string `json:"client_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Count       int32  `json:"count" swaggerignore:"true"`
}

type ClientCreate struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	CompanyID   string `json:"company_id"`
}

type ListClients struct {
	Clients []Client `json:"clients"`
	Count   int32    `json:"count"`
}
