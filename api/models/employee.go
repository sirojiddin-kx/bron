package models

type EmployeeCreate struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Description    string `json:"description"`
	GraphStartTime string `json:graph_start_time`
	GraphEndTime   string `json:"graph_end_time"`
	Position       string `json:"position"`
	CompanyId      string `json:"company_id"`
}

type Employee struct {
	ID             string `json:"guid"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Description    string `json:"description"`
	GraphStartTime string `json:graph_start_time`
	GraphEndTime   string `json:"graph_end_time"`
	Position       string `json:"position"`
	CompanyId      string `json:"company_id"`
	Count          int32  `json:"count" swaggerignore:"true"`
}

type EmployeeList struct {
	Employees []Employee `json:"employees"`
	Count     int32      `json:"count"`
}
