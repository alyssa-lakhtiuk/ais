package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) pricesByCat(c *gin.Context) {
	var res []entities.PriceByCat
	res, _ = h.services.Zvit.GetPricesByCategory()
	c.JSON(http.StatusOK, res)
}

func (h *Handler) checksByCat(c *gin.Context) {
	//category := c.Request.FormValue("")
	category := c.Param("id")
	var res []entities.CheckByCat
	res, _ = h.services.Zvit.GetChecksByCategory(category)
	for i := 0; i < len(res); i++ {
		res[i].Cat = category
	}
	c.JSON(http.StatusOK, res)
}
