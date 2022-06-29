package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCategory(c *gin.Context) {
	//employeeId, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	// throw error response
	//}
	var input entities.Category
	if err := c.BindJSON(&input); err != nil {
		// throw error response
	}
	id, err := h.services.Category.Create(input)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllCategories(c *gin.Context) {
	employees, err := h.services.Category.GetAll()
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, employees)
}

func (h *Handler) getCategoryByName(c *gin.Context) {
	categoryName := c.Param("name")
	category, err := h.services.Category.GetByName(categoryName)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) updateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// throw error response
	}
	var input entities.CategoryInput
	if err := c.BindJSON(&input); err != nil {
		// throw error response
	}
	if err := h.services.Category.Update(id, input); err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, "updated")
}

func (h *Handler) deleteCategory(c *gin.Context) {
	categoryName := c.Param("name")
	err := h.services.Category.Delete(categoryName)
	if err != nil {
		// throw error response
	}

	c.JSON(http.StatusOK, "deleted")
}
