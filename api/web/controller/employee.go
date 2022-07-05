package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
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
	input.Patronymic = ""
	input.Patronymic = c.Request.FormValue("patronymic")
	input.Role = c.Request.FormValue("emp_role")
	input.Salary, err = strconv.ParseFloat(c.Request.FormValue("salary"), 64)
	//std := c.Request.FormValue("start_day")
	//c.JSON(http.StatusOK, std)
	//c.JSON(http.StatusOK, "|||")
	//c.JSON(http.StatusOK, c.Request.FormValue("day_of_birth"))
	//c.JSON(http.StatusOK, "//")
	//layout := "2006-01-02"
	input.DateOfStart = c.Request.FormValue("day_of_start_job")
	//c.JSON(http.StatusOK, inputStDate)
	//c.JSON(http.StatusOK, "**")
	//helpArr := strings.Split(inputStDate, "-")
	//inputStDate = fmt.Sprintf("%s-%s-%s", helpArr[2], helpArr[1], helpArr[0])
	//k := strings.Split(inputStDate, ".")
	//strStartDate := fmt.Sprintf("%s-%s-%s", k[0], k[1], k[2])
	//startTime, _ := time.Parse(layout, inputStDate)
	//input.DateOfStart = startTime
	//c.JSON(http.StatusOK, input.DateOfStart)
	//c.JSON(http.StatusOK, "**")
	input.DateOfBirth = c.Request.FormValue("day_of_birth")
	//helpArr2 := strings.Split(inputBDate, ".")
	//BDate := fmt.Sprintf("%s-%s-%s", helpArr2[0], helpArr2[1], helpArr2[2])
	//BTime, _ := time.Parse(layout, inputBDate)
	//input.DateOfBirth = BTime
	//c.JSON(http.StatusOK, input.DateOfBirth)
	input.PhoneNumber = c.Request.FormValue("telephone")
	input.City = c.Request.FormValue("city_name")
	input.Street = c.Request.FormValue("street")
	input.ZipCode = c.Request.FormValue("index")
	//c.JSON(http.StatusOK, input.DateOfBirth)
	//c.JSON(http.StatusOK, "|||")
	//c.JSON(http.StatusOK, input.DateOfStart)
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
