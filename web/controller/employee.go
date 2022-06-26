package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
//type EmployeeService struct {
//	repo repository.EmployeeRepo
//}
//
//func NewEmployeeService(repo repository.EmployeeRepo) *EmployeeService {
//	return &EmployeeService{repo: repo}
//}

func (h *Handler) createEmployee(c *gin.Context) {
	//employeeId, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	// throw error response
	//}
	var input entities.Employee
	if err := c.BindJSON(&input); err != nil {
		// throw error response
	}
	id, err := h.services.Employee.Create(input)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllEmployees(c *gin.Context) {
	employees, err := h.services.Employee.GetAll()
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) getEmployeeByName(c *gin.Context) {
	employeeName := c.Param("Name")
	employee, err := h.services.Employee.GetByName(employeeName)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, employee)
}

func (h *Handler) updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var input entities.EmployeeInput
	if err := c.BindJSON(&input); err != nil {
		// throw error response
	}
	if err := h.services.Employee.Update(id, input); err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, "updated")
}

func (h *Handler) deleteEmployee(c gin.Context) {
	employeeId := c.Param("id")
	err := h.services.Employee.Delete(employeeId)
	if err != nil {
		// throw error response
	}

	c.JSON(http.StatusOK, "deleted")
}
