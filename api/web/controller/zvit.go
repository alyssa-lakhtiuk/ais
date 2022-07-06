package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) quantitiesByCat(c *gin.Context) {
	var res []entities.QuantityByCat
	res, _ = h.services.Zvit.GetQuantitiesByCategory()
	Tpl.ExecuteTemplate(c.Writer, "query_price_category.html", res)
}

func (h *Handler) checksByCat(c *gin.Context) {
	category := c.Request.FormValue("number2")
	var res []entities.CheckByCat
	res, _ = h.services.Zvit.GetChecksByCategory(category)
	for i := 0; i < len(res); i++ {
		res[i].Cat = category
	}
	//c.JSON(http.StatusOK, res)
	Tpl.ExecuteTemplate(c.Writer, "query_product_sell_check.html", res)
}

/////
func (h *Handler) countCities(c *gin.Context) {
	//category := c.Request.FormValue("")
	//category := c.Param("id")
	var res []entities.CountCustomersCities
	res, _ = h.services.Zvit.GetCountByCities()
	//c.JSON(http.StatusOK, res) // тут темплейт, замість цього рядка
	Tpl.ExecuteTemplate(c.Writer, "query_count_city.html", res)

}

func (h *Handler) checksByPrice(c *gin.Context) {
	//price, _ := strconv.Atoi(c.Request.FormValue("нейм з темплейта")) //
	price, _ := strconv.Atoi(c.Param("id")) // це з лінка бере параметр
	var res []entities.SecondStruct
	res, _ = h.services.Zvit.GetChecksByPrice(price)

	c.JSON(http.StatusOK, res) // тут темплейт

}
