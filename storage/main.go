package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirojiddin-kx/bron/storage/postgres"
	"github.com/sirojiddin-kx/bron/storage/repo"
)

type StorageI interface {
	CompanyServiceRepo() repo.CompanyServiceRepoI
	EmployeeRepo() repo.EmployeeRepoI
	ClientRepo() repo.ClientRepoI
	OrderRepo() repo.OrderRepoI
}

type storagePG struct {
	db             *sqlx.DB
	companyService repo.CompanyServiceRepoI
	employee       repo.EmployeeRepoI
	client         repo.ClientRepoI
	order          repo.OrderRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return storagePG{
		db:             db,
		companyService: postgres.NewCompanyServcieRepo(db),
		employee:       postgres.NewEmployeeRepo(db),
		client:         postgres.NewClientRepo(db),
		order:          postgres.NewOrderRepo(db),
	}
}

func (s storagePG) CompanyServiceRepo() repo.CompanyServiceRepoI {
	return s.companyService
}

func (s storagePG) EmployeeRepo() repo.EmployeeRepoI {
	return s.employee
}

func (s storagePG) ClientRepo() repo.ClientRepoI {
	return s.client
}

func (s storagePG) OrderRepo() repo.OrderRepoI {
	return s.order
}