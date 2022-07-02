package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createStoreProduct(c *gin.Context) {
	var input entities.StoreProduct
	if err := c.BindJSON(&input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "cant process data, check it")
		return
	}
	id, err := h.services.StoreProduct.Create(input)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "fail to create store product")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"productId": input.UPC,
		"id":        id,
	})
}

func (h *Handler) getAllStoreProducts(c *gin.Context) {
	products, err := h.services.StoreProduct.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
		// throw error response
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) getStoreProductByUpc(c *gin.Context) {
	upc := c.Param("upc")
	category, err := h.services.StoreProduct.GetByName(upc)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to get store product")
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) deleteStoreProduct(c *gin.Context) {
	id := c.Param("upc")
	err := h.services.StoreProduct.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, "deleted")
}

func (h *Handler) updateStoreProduct(c *gin.Context) {
	upc := c.Param("upc")
	var input entities.StoreProduct
	if err := c.BindJSON(&input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to parse input data")
		return
	}
	if err := h.services.StoreProduct.Update(upc, input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to update store product")
		return
	}
	c.JSON(http.StatusOK, "updated")
}
