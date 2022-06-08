package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"
	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/sirojiddin-kx/bron/storage/repo"
)

type clientService struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) repo.ClientRepoI {
	return clientService{db: db}
}

func (c clientService) Create(req models.ClientCreate) (err error) {

	query := `INSERT INTO client(
		guid,
		first_name,
		last_name,
		phone_number,
		company_id,
		created_at
	)
	VALUES(
		$1, 
		$2,
		$3,
		$4,
		$5,
		current_timestamp
	)`

	clientGuid := uuid.New()

	_, err = c.db.Exec(query, clientGuid, req.FirstName, req.LastName, req.PhoneNumber, req.CompanyID)
	if err != nil {
		return err
	}

	return
}

func (c clientService) GetAllClients(companyId string, limit, page int32) (resp []models.Client, err error) {
	var (
		offset = (page - 1) * limit
	)

	query := `
		SELECT
			count(1) over(),
			guid,
			first_name,
			last_name,
			phone_number
		FROM client
		WHERE company_id=$1
		ORDER BY created_at
		LIMIT $2 OFFSET $3
	`

	rows, err := c.db.Query(query, companyId, limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var scanObj models.Client
		err = rows.Scan(
			&scanObj.Count,
			&scanObj.Id,
			&scanObj.FirstName,
			&scanObj.LastName,
			&scanObj.PhoneNumber,
		)

		if err != nil {
			return
		}
		resp = append(resp, scanObj)
	}

	return
}

