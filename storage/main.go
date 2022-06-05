package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirojiddin-kx/bron/storage/postgres"
	"github.com/sirojiddin-kx/bron/storage/repo"
)

type StorageI interface {
	CompanyServiceRepo() repo.CompanyServiceRepoI
}

type storagePG struct {
	db             *sqlx.DB
	companyService repo.CompanyServiceRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return storagePG{
		db:             db,
		companyService: postgres.NewCompanyServcieRepo(db),
	}
}

func (s storagePG) CompanyServiceRepo() repo.CompanyServiceRepoI {
	return s.companyService
}
