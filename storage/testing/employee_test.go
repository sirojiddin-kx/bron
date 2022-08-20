package testing

import (
	"testing"

	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	err := TestStorage.EmployeeRepo().Create(models.EmployeeCreate{
		FirstName:      "Normuradov",
		LastName:       "Javlon",
		Description:    "Experienced Barber with 3 year",
		GraphStartTime: "09:00:00",
		GraphEndTime:   "20:00:00",
		CompanyId:      "b7ab5845-70e9-4632-84eb-80b59319cf8d",
		Position:       "barber",
	})
	assert.NoError(t, err)
}

func TestGetEmployeeList(t *testing.T) {
	var (
		companyId = "b7ab5845-70e9-4632-84eb-80b59319cf8d"
		limit, page int32 = 10, 1
	)
	_, err := TestStorage.EmployeeRepo().GetEmployeeList(companyId, limit, page)
	assert.NoError(t, err)
}

func TestGetEmployeeById(t *testing.T) {
	var (
		companyId = "b7ab5845-70e9-4632-84eb-80b59319cf8d"
		employeeId = "c1066438-2351-4fcd-b2ff-7348c146369b"
	)
	_, err := TestStorage.EmployeeRepo().GetEmployeeById(companyId, employeeId)
	assert.NoError(t, err)
}