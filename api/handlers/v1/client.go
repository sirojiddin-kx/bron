package v1

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/sirojiddin-kx/bron/pkg/logger"
	//"github.com/sirojiddin-kx/bron/pkg/logger"
)

//@Router /v1/client [post]
//@Summary Create Company Client
//@Description API for creating clients
//@Tags client
//@Accept json
//@Produce json
//@Param employee body models.ClientCreate true "company-service"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) ClientCreate(c *gin.Context) {
	var client models.ClientCreate

	err := c.ShouldBindJSON(&client)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	err = h.storage.ClientRepo().Create(client)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "client successfully created", "")
}

//@Router /v1/clients [get]
//@Summary List Clients
//@Description API for get list of clients
//@Tags client
//@Accept json
//@Produce json
//@Param company_id query string true "company_id"
//@Param limit query string false "limit"
//@Param page query string false "page"
//@Success 200 {object} models.CompanyServices
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) ListClients(c *gin.Context) {
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

	resp, err := h.storage.ClientRepo().GetAllClients(companyID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    models.ErrorCodeInternal,
			Message: err.Error(),
		})
		h.log.Error("database err", logger.Error(err))
		return
	}

	if len(resp) > 0 {
		count = resp[0].Count
	}

	c.JSON(http.StatusOK, gin.H{
		"clients": resp,
		"count":   count,
	})
}
