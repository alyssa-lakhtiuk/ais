package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

/////
func (h *Handler) countCities(c *gin.Context) {
	var res []entities.CountCustomersCities
	res, _ = h.services.Zvit.GetCountByCities()
	Tpl.ExecuteTemplate(c.Writer, "query_count_city.html", res)

}

func (h *Handler) checksByPrice(c *gin.Context) {
	price, _ := strconv.Atoi(c.Request.FormValue("price")) //
	//price, _ := strconv.Atoi(c.Param("price")) // це з лінка бере параметр
	var res []entities.SecondStruct
	res, _ = h.services.Zvit.GetChecksByPrice(price)
	Tpl.ExecuteTemplate(c.Writer, "query_check_category.html", res)
	//c.JSON(http.StatusOK, res) // тут темплейт

}
