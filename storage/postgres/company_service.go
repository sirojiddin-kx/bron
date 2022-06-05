package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirojiddin-kx/bron/api/storage/repo"
	"github.com/sirojiddin-kx/bron/storage/repo"
)

type companyService struct {
	db *sqlx.DB
}

func NewCompanyServcieRepo(db *sqlx.DB) repo.CompanyServiceRepoI {
	return companyService{db: db}
}
