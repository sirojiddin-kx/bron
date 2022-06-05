package v1

import (
	"net/http"
	"strconv"

	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/gin-gonic/gin"
)

//@Router /v1/company-service [post]
//@Summary Create Company Service
//@Description API for create company-service
//@Tags company-service
//@Accept json
//@Produce json
//@Param employee body models.CompanyServiceCreate true "bus-stop"
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

	res, err := h.storage.CompanyServiceRepo().Create(companyService)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "company service created successfully", res)
}