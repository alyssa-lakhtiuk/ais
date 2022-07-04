package controller

import (
	"ais/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *Handler) employeeCreated(c *gin.Context) {
	if c.Request.Method != "POST" {
		http.Redirect(c.Writer, c.Request, "/create-employee", http.StatusSeeOther)
		return
	}
	var err error
	var input entities.Employee
	//if err := c.BindJSON(&input); err != nil {
	//	respondWithError(c, http.StatusBadRequest, "unable to parse input data, check if it's correct")
	//	// throw error response
	//	return
	//}
	input.ID = c.Request.FormValue("lastname")
	input.SurName = c.Request.FormValue("lastname")
	input.FirstName = c.Request.FormValue("firstname")
	input.Patronymic = c.Request.FormValue("patronymic")
	input.Role = c.Request.FormValue("emp_role")
	input.Salary, err = strconv.ParseFloat(c.Request.FormValue("salary"), 64)
	layout := "2006-01-02"
	inputStDate := c.Request.FormValue("start_day")
	k := strings.Split(inputStDate, ".")
	strStartDate := fmt.Sprintf("%s-%s-%s", k[0], k[1], k[2])
	startTime, _ := time.Parse(layout, strStartDate)
	input.DateOfStart = startTime
	inputBDate := c.Request.FormValue("start_day")
	helpArr2 := strings.Split(inputBDate, ".")
	BDate := fmt.Sprintf("%s-%s-%s", helpArr2[0], helpArr2[1], helpArr2[2])
	BTime, _ := time.Parse(layout, BDate)
	input.DateOfBirth = BTime
	input.PhoneNumber = c.Request.FormValue("telephone")
	input.City = c.Request.FormValue("city_name")
	input.Street = c.Request.FormValue("street")
	input.ZipCode = c.Request.FormValue("index")
	_, err = h.services.Employee.Create(input)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "fail to create employee, check if your data is correct")
		c.JSON(http.StatusOK, err)
		//return
	}
	//c.JSON(http.StatusOK, input)
	err = Tpl.ExecuteTemplate(c.Writer, "done_employee.html", input)
	if err != nil {
		return
	}
}

func (h *Handler) createEmployee(c *gin.Context) {
	Tpl.ExecuteTemplate(c.Writer, "add_employee.html", nil)
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
	employeeName := c.Param("id")
	employee, err := h.services.Employee.GetByName(employeeName)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "employee with such name doesn't exist")
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *Handler) getEmployeeById(c *gin.Context) {
	employeeID := c.Param("id")
	employee, err := h.services.Employee.GetById(employeeID)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "employee with such id doesn't exist")
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
