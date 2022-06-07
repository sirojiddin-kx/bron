package models

type CompanyServiceCreate struct {
	Title     string `json:"title"`
	Duration  int32  `json:"duration"`
	Price     int32  `json:"price"`
	CompanyID string `json:"company_id"`
}

type CompanyService struct {
	ID        string `json:"guid"`
	Title     string `json:"title"`
	Duration  int32  `json:"duration"`
	Price     int32  `json:"price"`
	CompanyID string `json:"company_id"`
	Count     int32 `json:"count" swaggerignore:"true"`
}

type CompanyServices struct {
	Services []CompanyService `json:"services"`
	Count    int64            `json:"count" db:"count"`
}
