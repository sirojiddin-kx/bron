package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"
	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/sirojiddin-kx/bron/storage/repo"
)

type employeeService struct {
	db *sqlx.DB
}

func NewEmployeeRepo(db *sqlx.DB) repo.EmployeeRepoI {
	return employeeService{db: db}
}

func (c employeeService) Create(req models.EmployeeCreate) (err error) {

	query := `INSERT INTO employee(
		guid,
		first_name,
		last_name,
		description,
		start_time,
		end_time,
		company_id,
		position,
		create_at
	)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		current_timestamp
	)`

	employeeId := uuid.New()

	_, err = c.db.Exec(
		query,
		employeeId,
		req.FirstName,
		req.LastName,
		req.Description,
		req.GraphStartTime,
		req.GraphEndTime,
		req.CompanyId,
		req.Position,
	)

	if err != nil {
		return err
	}

	return
}

func (c employeeService) GetEmployeeList(companyID string, limit, page int32) (resp []models.Employee, err error) {
	var (
		offset = (page - 1) * limit
	)

	query := `
		SELECT 
			count(1) over(),
			guid,
			first_name,
			last_name,
			description,
			start_time,
			end_time,
			company_id,
			position
		FROM employee
		WHERE company_id = $1
		limit $2 offset $3
	`

	rows, err := c.db.Query(query, companyID, limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		err = rows.Scan(
			&employee.Count,
			&employee.ID,
			&employee.FirstName,
			&employee.LastName,
			&employee.Description,
			&employee.GraphStartTime,
			&employee.GraphEndTime,
			&employee.CompanyId,
			&employee.Position,
		)

		if err != nil {
			return
		}

		resp = append(resp, employee)
	}

	return
}

func (c employeeService) GetEmployeeById(companyId, employeeId string) (resp models.Employee, err error) {
	query := `
		SELECT
			guid,
			first_name,
			last_name,
			description,
			start_time,
			end_time,
			position,
			company_id
		FROM employee
		WHERE guid=$1 and company_id=$2
	`

	err = c.db.QueryRow(query, employeeId, companyId).Scan(
		&resp.ID,
		&resp.FirstName,
		&resp.LastName,
		&resp.Description,
		&resp.GraphStartTime,
		&resp.GraphEndTime,
		&resp.Position,
		&resp.CompanyId,
	)

	if err != nil {
		return models.Employee{}, err
	}

	return
}

