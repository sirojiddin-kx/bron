package v1

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/sirojiddin-kx/bron/pkg/logger"
)

//@Router /v1/company-service [post]
//@Summary Create Company Service
//@Description API for create company-service
//@Tags company-service
//@Accept json
//@Produce json
//@Param employee body models.CompanyServiceCreate true "company-service"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) CompanyServcieCreate(c *gin.Context) {
	var companyService models.CompanyServiceCreate

	err := c.ShouldBind(&companyService)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	err = h.storage.CompanyServiceRepo().Create(companyService)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "company service created successfully", "")
}


//@Router /v1/company-service [get]
//@Summary List Company Services
//@Description API for get list company_services
//@Tags company-service
//@Accept json
//@Produce json
//@Param company_id query string true "company_id"
//@Param limit query string false "limit"
//@Param page query string false "page"
//@Success 200 {object} models.CompanyServices
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) ListCompanyService(c *gin.Context) {
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

	resp, err := h.storage.CompanyServiceRepo().GetCompanyServices(companyID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code: models.ErrorCodeInternal,
			Message: err.Error(),
		})
		h.log.Error("database err", logger.Error(err))
		return
	}

	if len(resp) > 0 {
		count = resp[0].Count
	}

	c.JSON(http.StatusOK, gin.H{
		"services": resp,
		"count": count,
	})

	
}

