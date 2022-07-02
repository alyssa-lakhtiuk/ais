package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// це тільки контролер, який колає відповідні сервіси!
// мейбі навіть зробити тут один файлик

// ввести відомості про новий товар
// оновити відомості про товар
// вилучити відомості про товар

func (h *Handler) createProduct(c *gin.Context) {
	var input entities.Product
	if err := c.BindJSON(&input); err != nil {
		// throw error response
	}
	id, err := h.services.Product.Create(input)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"productId": input.Id,
		"id":        id,
	})
}

func (h *Handler) getAllProducts(c *gin.Context) {
	products, err := h.services.Product.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		// throw error response
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) getProductByName(c *gin.Context) {
	productName := c.Param("name")
	category, err := h.services.Product.GetByName(productName)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		// err
	}
	err = h.services.Product.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "deleted")
}

func (h *Handler) updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		// err
	}
	var input entities.Product
	if err := c.BindJSON(&input); err != nil {
		// throw error response
	}
	if err := h.services.Product.Update(id, input); err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, "updated")
}
