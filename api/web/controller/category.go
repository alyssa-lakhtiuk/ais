package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) categoryCreated(c *gin.Context) {
	if c.Request.Method != "POST" {
		http.Redirect(c.Writer, c.Request, "/create-category", http.StatusSeeOther)
		return
	}
	var input entities.Category
	//if err := c.BindJSON(&input); err != nil {
	//	// throw error response
	//	respondWithError(c, http.StatusBadRequest, "unable to parse input data")
	//	return
	//}
	input.Name = c.Request.FormValue("name_category")
	_, err := h.services.Category.Create(input)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to create category")
		return
	}
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"categoryNumber": input.Number,
	//	"id":             id,
	//})
	err = Tpl.ExecuteTemplate(c.Writer, "done_category.html", input)
	if err != nil {
		return
	}
}

func (h *Handler) createCategory(c *gin.Context) {
	Tpl.ExecuteTemplate(c.Writer, "add_category.html", nil)
}

func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.services.Category.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
		// throw error response
	}
	Tpl.ExecuteTemplate(c.Writer, "manager_category.html", categories)
	//c.JSON(http.StatusOK, categories)
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"categoryNumber": categories[0].Number,
	//	"id":             categories[0].Name,
	//})
	//respondWithJSON(c, http.StatusOK, categories)
}

func (h *Handler) getCategoryByName(c *gin.Context) {
	categoryName := c.Param("id") // localhost:8080/category/soft-cheese
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

func (h *Handler) getCategoryByNumber(c *gin.Context) {
	categoryName := c.Param("id") // localhost:8080/category/soft-cheese
	category, err := h.services.Category.GetByName(categoryName)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to get category")
		//return
	}
	c.JSON(http.StatusOK, category)
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
