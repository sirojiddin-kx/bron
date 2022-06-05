package repo

import (
	"github.com/sirojiddin-kx/bron/api/models"
)

type CompanyServiceRepoI interface {
	Create(req models.CompanyServiceCreate) (err error)
	GetAll(companyId string, limit, page int32) (res models.CompanyServices, err error)
	Delete(companyId string) (err error)
	Update(req models.CompanyService) (err error)
}
