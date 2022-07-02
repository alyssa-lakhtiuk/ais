package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createCheck(c *gin.Context) {
	var input []entities.CheckInput
	if err := c.BindJSON(&input); err != nil {
		// throw error response
	}
	id, err := h.services.Check.Create(input)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllChecks(c *gin.Context) {
	cc, err := h.services.Check.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		// throw error response
	}
	c.JSON(http.StatusOK, cc)
}

func (h *Handler) getCheckByNumber(c *gin.Context) {
	checkNum := c.Param("number")
	category, err := h.services.Check.GetByNumber(checkNum)
	if err != nil {
		// throw error response
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) deleteCheck(c *gin.Context) {
	id := c.Param("number")
	err := h.services.Check.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, "deleted")
}
