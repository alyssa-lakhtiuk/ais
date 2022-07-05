package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) customerCardCreated(c *gin.Context) {
	var input entities.CustomerCard
	var err error
	input.Number = c.Request.FormValue("card_number")
	input.CustomerSurname = c.Request.FormValue("lastname")
	input.CustomerName = c.Request.FormValue("firstname")
	input.CustomerPatronymic.String = c.Request.FormValue("patronymic")
	input.PhoneNumber = c.Request.FormValue("telephone")
	input.City.String = c.Request.FormValue("city_name")
	input.Street.String = c.Request.FormValue("street")
	input.ZipCode.String = c.Request.FormValue("index")
	input.Percent, err = strconv.Atoi(c.Request.FormValue("percents"))
	//if err := c.BindJSON(&input); err != nil {
	//	// throw error response
	//}
	_, err = h.services.CustomerCard.Create(input)
	if err != nil {
		// throw error response
	}
	Tpl.ExecuteTemplate(c.Writer, "done_client.html", input)
}
func (h *Handler) createCustomerCard(c *gin.Context) {
	Tpl.ExecuteTemplate(c.Writer, "add_client.html", nil)
}

func (h *Handler) getAllCustomerCards(c *gin.Context) {
	authHeader, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "authorization first", nil)
	}
	currentEmplId := authHeader.Value
	roleDromDB, err := h.services.Role.GetByIdEmployee(currentEmplId)

	cc, err := h.services.CustomerCard.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
		// throw error response
	}
	//c.JSON(http.StatusOK, cc)
	if roleDromDB.Role == "manager" {
		Tpl.ExecuteTemplate(c.Writer, "manager_client.html", cc)
	} else {
		Tpl.ExecuteTemplate(c.Writer, "cashier_client.html", cc)
	}

}

func (h *Handler) getCustomerCardByNumber(c *gin.Context) {
	categoryName := c.Param("id")
	category, err := h.services.CustomerCard.GetByNumber(categoryName)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to get customer card")
		return
		// throw error response
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) updateCustomerCardOpen(c *gin.Context) {
	cardNumber := c.Request.FormValue("number_card")
	ccToEdit, _ := h.services.CustomerCard.GetByNumber(cardNumber)
	Tpl.ExecuteTemplate(c.Writer, "edit_client.html", ccToEdit)
}

func (h *Handler) updateCustomerCard(c *gin.Context) {
	//id := c.Param("id")
	var input entities.CustomerCard
	cardNumber := c.Request.FormValue("card_number")
	//input.Number = c.Request.FormValue("card_number")
	input.CustomerSurname = c.Request.FormValue("surname")
	input.CustomerName = c.Request.FormValue("firstname")
	input.CustomerPatronymic.String = c.Request.FormValue("patronymic")
	input.PhoneNumber = c.Request.FormValue("telephone")
	input.City.String = c.Request.FormValue("city_name")
	input.Street.String = c.Request.FormValue("street")
	input.ZipCode.String = c.Request.FormValue("index")
	input.Percent, _ = strconv.Atoi(c.Request.FormValue("percents"))

	//if err := c.BindJSON(&input); err != nil {
	//	// throw error response
	//	respondWithError(c, http.StatusBadRequest, "unable to parse input data")
	//	return
	//}
	if err := h.services.CustomerCard.Update(cardNumber, input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to update")
		return
	}
	h.getAllCustomerCards(c)
	//Tpl.ExecuteTemplate(c.Writer, "edit_employee.html", entities.Message{Mess: "customer card updated"})
	//c.JSON(http.StatusOK, cardNumber)
}

func (h *Handler) deleteCustomerCard(c *gin.Context) {
	id := c.Request.FormValue("number_card")
	err := h.services.CustomerCard.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//c.JSON(http.StatusOK, id)
	h.getAllCustomerCards(c)
}
