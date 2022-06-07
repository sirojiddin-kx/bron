package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirojiddin-kx/bron/storage/postgres"
	"github.com/sirojiddin-kx/bron/storage/repo"
)

type StorageI interface {
	CompanyServiceRepo() repo.CompanyServiceRepoI
	EmployeeRepo() repo.EmployeeRepoI
}

type storagePG struct {
	db             *sqlx.DB
	companyService repo.CompanyServiceRepoI
	employee       repo.EmployeeRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return storagePG{
		db:             db,
		companyService: postgres.NewCompanyServcieRepo(db),
		employee:       postgres.NewEmployeeRepo(db),
	}
}

func (s storagePG) CompanyServiceRepo() repo.CompanyServiceRepoI {
	return s.companyService
}

func (s storagePG) EmployeeRepo() repo.EmployeeRepoI {

	return s.employee
}
