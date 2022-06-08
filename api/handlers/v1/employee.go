package v1

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirojiddin-kx/bron/api/models"
	"github.com/sirojiddin-kx/bron/pkg/logger"
	//"github.com/sirojiddin-kx/bron/pkg/logger"
)

//@Router /v1/employee [post]
//@Summary Create Employee
//@Description API for creating client
//@Tags employee
//@Accept json
//@Produce json
//@Param employee body models.EmployeeCreate true "employee"
//@Success 200 {object} models.SuccessModel
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) EmployeeCreate(c *gin.Context) {
	var employee models.EmployeeCreate

	err := c.ShouldBindJSON(&employee)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusBadRequest, "bad request", err)
		return
	}

	err = h.storage.EmployeeRepo().Create(employee)
	if err != nil {
		h.HandleErrorResponse(c, http.StatusInternalServerError, "database error", err)
		return
	}

	h.HandleSuccessResponse(c, 201, "employee created successfully", "")
}

//@Router /v1/employee [get]
//@Summary List Employee
//@Description API for get list employeeslsls
//@Tags employee
//@Accept json
//@Produce json
//@Param company_id query string true "company_id"
//@Param limit query string false "limit"
//@Param page query string false "page"
//@Success 200 {object} models.EmployeeList
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) ListEmployee(c *gin.Context) {
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

	resp, err := h.storage.EmployeeRepo().GetEmployeeList(companyID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code: models.ErrorCodeInternal,
			Message: err.Error(),
		})
		h.log.Error("database error", logger.Error(err))
		return
	}

	if len(resp) > 0 {
		count = resp[0].Count
	}

	c.JSON(http.StatusOK, gin.H{
		"employees": resp,
		"count": count,
	})
}


//@Router /v1/employee/{employee_id} [get]
//@Summary Get Employee By ID
//@Description API for getting single Employee
//@Tags employee
//@Accept json
//@Produce json
//@Param company_id query string true "company_id"
//@Param employee_id path string true "guid"
//@Success 200 {object} models.EmployeeList
//@Failure 400 {object} models.ResponseError
//@Failure 500 {object} models.ResponseError
func (h handlerV1) GetEmployeeByID(c *gin.Context) {
	var (
		employeeId = c.Param("employee_id")
		companyID = c.DefaultQuery("company_id", "")
	)

	resp, err := h.storage.EmployeeRepo().GetEmployeeById(companyID, employeeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code: models.ErrorCodeInternal,
			Message: err.Error(),
		})
		h.log.Error("database err", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}



