package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCategory(c *gin.Context) {

	var input entities.Category
	if err := c.BindJSON(&input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to parse input data")
		return
	}
	id, err := h.services.Category.Create(input)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to create category")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"categoryNumber": input.Number,
		"id":             id,
	})
}

func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.services.Category.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
		// throw error response
	}
	c.JSON(http.StatusOK, categories)
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"categoryNumber": categories[0].Number,
	//	"id":             categories[0].Name,
	//})
	//respondWithJSON(c, http.StatusOK, categories)
}

func (h *Handler) getCategoryByName(c *gin.Context) {
	categoryName := c.Param("name")
	category, err := h.services.Category.GetByName(categoryName)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to get category")
		return
	}
	c.JSON(http.StatusOK, category)
	//err = respondWithJSON(h, http.StatusOK, category)
	//if err != nil {
	//	respondWithErrorLog(t.log, w, http.StatusInternalServerError, err)
	//}
}

func (h *Handler) updateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to parse input url")
		c.JSON(http.StatusOK, id)
		return
		// throw error response
	}
	var input entities.CategoryInput
	if err := c.BindJSON(&input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to parse input data")
		return
	}
	if err := h.services.Category.Update(id, input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to update")
		return
	}
	c.JSON(http.StatusOK, "updated")
}

func (h *Handler) deleteCategory(c *gin.Context) {
	categoryName := c.Param("name")
	err := h.services.Category.Delete(categoryName)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, categoryName)
}
