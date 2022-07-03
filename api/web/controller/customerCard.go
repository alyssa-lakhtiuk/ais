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
	input.CustomerPatronymic = c.Request.FormValue("patronymic")
	input.PhoneNumber = c.Request.FormValue("telephone")
	input.City = c.Request.FormValue("city_name")
	input.Street = c.Request.FormValue("street")
	input.ZipCode = c.Request.FormValue("index")
	input.Percent, err = strconv.Atoi(c.Request.FormValue("percents"))
	//if err := c.BindJSON(&input); err != nil {
	//	// throw error response
	//}
	_, err = h.services.CustomerCard.Create(input)
	if err != nil {
		// throw error response
	}
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"id": id,
	//})
	Tpl.ExecuteTemplate(c.Writer, "done_client.html", input)
}
func (h *Handler) createCustomerCard(c *gin.Context) {
	Tpl.ExecuteTemplate(c.Writer, "add_client.html", nil)
}

func (h *Handler) getAllCustomerCards(c *gin.Context) {
	cc, err := h.services.CustomerCard.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
		// throw error response
	}
	//c.JSON(http.StatusOK, cc)
	Tpl.ExecuteTemplate(c.Writer, "manager_client.html", cc)
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

func (h *Handler) updateCustomerCard(c *gin.Context) {
	id := c.Param("id")
	var input entities.CustomerCard
	if err := c.BindJSON(&input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to parse input data")
		return
	}
	if err := h.services.CustomerCard.Update(id, input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to update")
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h *Handler) deleteCustomerCard(c *gin.Context) {
	id := c.Param("id")
	err := h.services.CustomerCard.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, id)
}
