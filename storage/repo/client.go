package repo

import (
	"github.com/sirojiddin-kx/bron/api/models"
)

type ClientRepoI interface {
	Create(req models.ClientCreate) (err error)
	GetAllClients(companyId string, limit, page int32) (resp []models.Client, err error) 

}