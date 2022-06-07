package repo

import (
	"github.com/sirojiddin-kx/bron/api/models"
)

type EmployeeRepoI interface {
	Create(req models.EmployeeCreate) (err error)
	GetEmployeeList(companyID string, limit, page int32) (resp []models.Employee, err error)
	//GetEmployeeById(companyId, employeeID string) (resp models.Employee, err error)
}
