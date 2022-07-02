package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSale(c *gin.Context) {
	var input entities.Sale
	if err := c.BindJSON(&input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to parse input data")
		return
	}
	id, err := h.services.Sale.Create(input)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to create sale")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllSales(c *gin.Context) {
	cc, err := h.services.Sale.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		// throw error response
	}
	c.JSON(http.StatusOK, cc)
}

func (h *Handler) getSaleByUpcCheck(c *gin.Context) {
	upc := c.Param("upc")
	checkNumber := c.Param("check_number")
	sale, err := h.services.Sale.GetByUpcCheck(upc, checkNumber)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, sale)
}

func (h *Handler) deleteSale(c *gin.Context) {
	upc := c.Param("upc")
	checkNumber := c.Param("check_number")
	err := h.services.Sale.Delete(upc, checkNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "deleted")
}
