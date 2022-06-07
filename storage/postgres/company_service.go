package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"
	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/sirojiddin-kx/bron/storage/repo"
)

type companyService struct {
	db *sqlx.DB
}

func NewCompanyServcieRepo(db *sqlx.DB) repo.CompanyServiceRepoI {
	return companyService{db: db}
}

func (c companyService) Create(req models.CompanyServiceCreate) (err error) {

	query := `INSERT INTO company_service(
		guid,
		title,
		duration,
		price,
		company_id,
		created_at
	)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		current_timestamp
	)
	`
	companyServiceGuid := uuid.New()

	_, err = c.db.Exec(query, companyServiceGuid, req.Title, req.Duration, req.Price, req.CompanyID)
	if err != nil {
		return err
	}

	return
}

func (c companyService) GetCompanyServices(companyId string, limit, page int32) (resp []models.CompanyService, err error) {
	var (
		offset = (page - 1) * limit
	)
	query := `
		SELECT 
			count(1) over(),
			guid,
			title,
			duration,
			price,
			company_id
		FROM company_service
		WHERE company_id = $1
		limit $2 offset $3
	`

	rows, err := c.db.Query(query, companyId, limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var companyService models.CompanyService
		err = rows.Scan(
			&companyService.Count,
			&companyService.ID,
			&companyService.Title,
			&companyService.Duration,
			&companyService.Price,
			&companyService.CompanyID,
		)

		if err != nil {
			return
		}
		resp = append(resp, companyService)
	}

	return
}
