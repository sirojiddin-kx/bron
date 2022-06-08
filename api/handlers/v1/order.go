package v1

import (
	"fmt"
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/sirojiddin-kx/bron/pkg/logger"
)

//@Router /v1/order [post]
//@Summary Create Order
//@Description API for creating orders
//@Tags order
//@Accept json
//@Produce json
//@Param order body models.CreateOrder true "order"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) CreateOrder(c *gin.Context) {
	var order models.CreateOrder

	err := c.ShouldBindJSON(&order)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	err = h.storage.OrderRepo().Create(order)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "order created succesfully", "")
}

//@Router /v1/orders [get]
//@Summary List Orders
//@Description API for get list orders
//@Tags order
//@Accept json
//@Produce json
//@Param company_id query string true "company_id"
//@Param limit query string false "limit"
//@Param page query string false "page"
//@Success 200 {object} models.ListOrders
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) ListOrders(c *gin.Context) {
	var (
		count int32
	)

	companyID := c.DefaultQuery("company_id", "")
	limit, err := ParseQueryParam(c, "limit", "10")
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "can't parse limit", err)
		return
	}

	page, err := ParseQueryParam(c, "page", "10")
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "can't parse page", err)
	}

	resp, err := h.storage.OrderRepo().GetOrders(companyID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    models.ErrorCodeInternal,
			Message: err.Error(),
		})
		h.log.Error("database error", logger.Error(err))
		return
	}
	fmt.Println("Resp========+> ", resp)
	if len(resp) > 0 {
		count = resp[0].Count
	}

	c.JSON(http.StatusOK, gin.H{
		"employees": resp,
		"count":     count,
	})
}
