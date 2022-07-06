package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
)

func (h *Handler) employeeCreated(c *gin.Context) {
	if c.Request.Method != "POST" {
		http.Redirect(c.Writer, c.Request, "/create-employee", http.StatusSeeOther)
		return
	}
	var err error
	var input entities.Employee
	input.ID = c.Request.FormValue("lastname")
	input.SurName = c.Request.FormValue("lastname")
	input.FirstName = c.Request.FormValue("firstname")
	input.Patronymic.String = c.Request.FormValue("patronymic")
	input.Role = c.Request.FormValue("emp_role")
	input.Salary, err = strconv.ParseFloat(c.Request.FormValue("salary"), 64)
	input.DateOfStart = c.Request.FormValue("day_of_start_job")
	input.DateOfBirth = c.Request.FormValue("day_of_birth")
	input.PhoneNumber = c.Request.FormValue("telephone")
	input.City = c.Request.FormValue("city_name")
	input.Street = c.Request.FormValue("street")
	input.ZipCode = c.Request.FormValue("index")
	_, err = h.services.Employee.Create(input)
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "fail to create employee, check if your data is correct")
		//c.JSON(http.StatusOK, err)
		//return
	}
	//c.JSON(http.StatusOK, input)
	err = Tpl.ExecuteTemplate(c.Writer, "done_employee.html", input)
}

func (h *Handler) createEmployee(c *gin.Context) {
	Tpl.ExecuteTemplate(c.Writer, "add_employee.html", nil)
}

func (h *Handler) getAllEmployees(c *gin.Context) {
	employees, err := h.services.Employee.GetAll()
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "can't get all employees, maybe there is no even one")
		//return
	}
	//c.JSON(http.StatusOK, employees)
	sort.SliceStable(employees, func(i, j int) bool {
		return employees[i].SurName < employees[j].SurName
	})
	Tpl.ExecuteTemplate(c.Writer, "manager_employee.html", employees)
}

func (h *Handler) getEmployeeByName(c *gin.Context) {
	employeeName := c.Param("id")
	employee, err := h.services.Employee.GetByName(employeeName)
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "employee with such name doesn't exist")
		//return
	}
	c.JSON(http.StatusOK, employee)
}

func (h *Handler) getEmployeeById(c *gin.Context) {
	employeeID := c.Param("id")
	employee, err := h.services.Employee.GetById(employeeID)
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "employee with such id doesn't exist")
		//return
	}
	c.JSON(http.StatusOK, employee)
}

var inputtedId string

func (h *Handler) updateEmployeeOpen(c *gin.Context) {
	inputtedId = c.Request.FormValue("id")
	employeeToEdit, _ := h.services.Employee.GetById(inputtedId)
	Tpl.ExecuteTemplate(c.Writer, "edit_employee.html", employeeToEdit)
}

func (h *Handler) updateEmployee(c *gin.Context) {
	var input entities.EmployeeInput
	id := inputtedId
	//input.ID = c.Request.FormValue("lastname")
	input.SurName = c.Request.FormValue("lastname")
	input.FirstName = c.Request.FormValue("firstname")
	input.Patronymic.String = c.Request.FormValue("patronymic")
	input.Role = c.Request.FormValue("emp_role")
	input.Salary, _ = strconv.ParseFloat(c.Request.FormValue("salary"), 64)
	input.DateOfStart = c.Request.FormValue("day_of_start_job")
	input.DateOfBirth = c.Request.FormValue("day_of_birth")
	input.PhoneNumber = c.Request.FormValue("telephone")
	input.City = c.Request.FormValue("city_name")
	input.Street = c.Request.FormValue("street")
	input.ZipCode = c.Request.FormValue("ZipCode")
	//if err := c.BindJSON(&input); err != nil {
	//	// throw error response
	//	respondWithError(c, http.StatusBadRequest, "unable to read input data, check is it correct")
	//	return
	//}
	if err := h.services.Employee.Update(id, input); err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "unable to update")
		//return
	}
	h.getAllEmployees(c)
	//Tpl.ExecuteTemplate(c.Writer, "manager_employee.html", )
}

func (h *Handler) deleteEmployee(c *gin.Context) {
	employeeId := c.Request.FormValue("id")
	err := h.services.Employee.Delete(employeeId)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to delete employee")
		// throw error response
		return
	}
	h.getAllEmployees(c)
	//c.JSON(http.StatusOK, employeeId)
}

func (h *Handler) onlyOneEmployeeCategory(c *gin.Context) {
	role := c.Request.FormValue("sort_role")
	if role != "manager" && role != "cashier" {
		return
	}
	employees, err := h.services.Employee.GetAllByCategory(role)
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "can't get all employees, maybe there is no even one")
		//return
	}
	//c.JSON(http.StatusOK, employees)
	sort.SliceStable(employees, func(i, j int) bool {
		return employees[i].SurName < employees[j].SurName
	})
	Tpl.ExecuteTemplate(c.Writer, "manager_employee.html", employees)
}

func (h *Handler) searchEmployee(c *gin.Context) {
	role := c.Request.FormValue("sort_role")
	if role != "manager" && role != "cashier" {
		return
	}
	employees, err := h.services.Employee.GetAllByCategory(role)
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "can't get all employees, maybe there is no even one")
		//return
	}
	//c.JSON(http.StatusOK, employees)
	sort.SliceStable(employees, func(i, j int) bool {
		return employees[i].SurName < employees[j].SurName
	})
	Tpl.ExecuteTemplate(c.Writer, "manager_employee.html", employees)
}
