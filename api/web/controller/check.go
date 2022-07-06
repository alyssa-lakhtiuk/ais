package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) addProductToCheck(c *gin.Context) {
	var chIn []entities.CheckInput
	var chF entities.CheckInput
	chF.UPC = c.Request.FormValue("upc1")
	chF.ProductNumber, _ = strconv.Atoi(c.Request.FormValue("number1"))
	chF.CustomerNumber = c.Request.FormValue("description")
	authHeader, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "authorization first", nil)
	}
	chF.IDEmployee = authHeader.Value
	chIn = append(chIn, chF)
	var chF2 entities.CheckInput
	chF2.UPC = c.Request.FormValue("upc2")
	chF2.ProductNumber, _ = strconv.Atoi(c.Request.FormValue("number2"))
	chF2.IDEmployee = chF.IDEmployee
	chF2.CustomerNumber = chF.CustomerNumber
	var chF3 entities.CheckInput
	chF3.UPC = c.Request.FormValue("upc3")
	chF3.ProductNumber, _ = strconv.Atoi(c.Request.FormValue("number3"))
	chF3.IDEmployee = chF.IDEmployee
	chF3.CustomerNumber = chF.CustomerNumber
	var chF4 entities.CheckInput
	chF4.UPC = c.Request.FormValue("upc4")
	chF4.ProductNumber, _ = strconv.Atoi(c.Request.FormValue("number4"))
	chF4.IDEmployee = chF.IDEmployee
	chF4.CustomerNumber = chF.CustomerNumber
	h.services.Check.Create(chIn)
	h.getAllChecks(c)
}

func (h *Handler) createCheck(c *gin.Context) {
	var input []entities.CheckInput
	id, err := h.services.Check.Create(input)
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "unable to create check")
		//return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllChecks(c *gin.Context) {
	cc, err := h.services.Check.GetAll()
	if err != nil {
		//c.JSON(http.StatusBadRequest, err)
		// throw error response
		//return
	}
	Tpl.ExecuteTemplate(c.Writer, "cashier_check.html", cc)
	//c.JSON(http.StatusOK, cc)
}

func (h *Handler) getCheckByNumber(c *gin.Context) {
	checkNum := c.Param("number")
	category, err := h.services.Check.GetByNumber(checkNum)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to get check")
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) deleteCheck(c *gin.Context) {
	id := c.Param("number")
	err := h.services.Check.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//c.JSON(http.StatusOK, id)
}
