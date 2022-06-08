package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pborman/uuid"
	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/sirojiddin-kx/bron/storage/repo"
)

type orderService struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) repo.OrderRepoI {
	return orderService{db: db}
}

func (c orderService) Create(req models.CreateOrder) (err error) {
	var (
		id string
	)

	query := `INSERT INTO orders(
		guid,
		employee_id,
		client_id,
		services,
		start_time,
		order_date,
		company_id,
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
		current_timestamp
	) RETURNING guid`

	orderGuid := uuid.New()
	err = c.db.QueryRow(query, orderGuid, req.EmployeeId, req.ClientId, req.Services, req.StartTime, req.Date, req.CompanyId).Scan(&id)
	if err != nil {
		return err
	}

	fmt.Println("Guid", id)

	updatePriceAndEndTime := `
	with data as (
		select sum(duration) total_duration, sum(price) price from company_service where guid=any(select unnest(services) from orders where guid=$1)
	)
	update orders set end_time = (start_time + (select total_duration from  data) * interval '1 minute'), total_price= (select price from data) where guid=$1;
	`
	_, err = c.db.Exec(updatePriceAndEndTime, id)
	if err != nil {
		return err
	}

	return
}

func (c orderService) GetOrders(companyId string, limit, page  int32) (resp []models.Order, err error) {
	var (
		offset = (page - 1) * limit
	)

	query := `
		SELECT 
			count(1) over(),
			o.guid,
			concat(e.first_name, ' ', e.last_name),
			concat(c.first_name, ' ', c.last_name),
			c.phone_number,
			o.order_date,
			o.start_time,
			o.end_time,
			o.total_price
		FROM orders o
		LEFT JOIN employee e on o.employee_id = e.guid
		LEFT JOIN client c on o.client_id = c.guid 
		WHERE o.company_id = $1
		limit $2 offset $3
	`

	rows, err := c.db.Query(query, companyId, limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var order models.Order
		err = rows.Scan(
			&order.Count,
			&order.Id,
			&order.EmployeeFullname,
			&order.ClientFullname,
			&order.ClientContact,
			&order.OrderDate,
			&order.StartTime,
			&order.EndTime,
			&order.TotalPrice,
		)

		if err != nil {
			return
		}
		resp = append(resp, order)
	}
	fmt.Println("RESP === > ", resp)
	return
}
