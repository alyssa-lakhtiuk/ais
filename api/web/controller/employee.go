package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createEmployee(c *gin.Context) {
	var input entities.Employee
	if err := c.BindJSON(&input); err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to parse input data, check if it's correct")
		// throw error response
		return
	}
	_, err := h.services.Employee.Create(input)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "fail to create employee, check if your data is correct")
		c.JSON(http.StatusOK, err)
		//return
	}
	c.JSON(http.StatusOK, input)
}

func (h *Handler) getAllEmployees(c *gin.Context) {
	employees, err := h.services.Employee.GetAll()
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "can't get all employees, maybe there is no even one")
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) getEmployeeByName(c *gin.Context) {
	employeeName := c.Param("name")
	employee, err := h.services.Employee.GetByName(employeeName)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "employee with such name doesn't exist")
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *Handler) updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var input entities.EmployeeInput
	if err := c.BindJSON(&input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to read input data, check is it correct")
		return
	}
	if err := h.services.Employee.Update(id, input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to update")
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) deleteEmployee(c *gin.Context) {
	employeeId := c.Param("id")
	err := h.services.Employee.Delete(employeeId)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to delete employee")
		// throw error response
		return
	}
	c.JSON(http.StatusOK, employeeId)
}
