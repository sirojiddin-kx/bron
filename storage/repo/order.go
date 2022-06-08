package repo

import (
	"github.com/sirojiddin-kx/bron/api/models"
)

type OrderRepoI interface {
	Create(req models.CreateOrder) (err error)
	GetOrders(companyId string, limit, page int32) (res []models.Order, err error)

}