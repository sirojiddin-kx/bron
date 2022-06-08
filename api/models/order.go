package models

type CreateOrder struct {
	EmployeeId string   `json:"employee_id"`
	ClientId   string   `json:"client_id"`
	Services   []string `json:"services"`
	StartTime  string   `json:"start_time"`
	CompanyId  string   `json:"company_id"`
	Date       string   `json:"order_date"`
}

type Order struct {
	Id               string   `json:"order_id"`
	EmployeeFullname string   `json:"employee_fullname"`
	ClientFullname   string   `json:"client_fullname"`
	ClientContact    string   `json:"client_contact"`
	Services         []string `json:"services"`
	StartTime        string   `json:"start_time"`
	EndTime          string   `json:"end_time"`
	OrderDate        string   `json:"order_date"`
	TotalPrice       string   `json:"total_price"`
	Count            int32   `json:"count" swaggerignore:"true"`
}

type ListOrders struct {
	Orders []Order `json:"orders"`
	Count  int32   `json:"count"`
}
