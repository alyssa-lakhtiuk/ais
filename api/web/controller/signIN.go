package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) signIn(c *gin.Context) {
	var data entities.SignIn
	data.Phone = c.Request.FormValue("phone_number")
	inputtedPassword := c.Request.FormValue("password")
	hashedPassword := generatePasswordHash(inputtedPassword)
	dataFromDB, err := h.services.Role.GetByPhone(data.Phone)
	if err != nil {

	}
	if hashedPassword != dataFromDB.Password {
		c.HTML(http.StatusUnauthorized, "invalid password or number", nil)
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "Authorization",
		Value:   dataFromDB.IdEmployee,
		Expires: time.Now().Add(48 * time.Hour)})
	if dataFromDB.Role == "manager" {
		Tpl.ExecuteTemplate(c.Writer, "manager_homepage.html", nil)
	} else if dataFromDB.Role == "cashier" {
		Tpl.ExecuteTemplate(c.Writer, "cashier_homepage.html", nil)
	}
}

func (h *Handler) goToHomePage(c *gin.Context) {
	authHeader, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "authorization first", nil)
	}
	currentEmplId := authHeader.Value
	roleDromDB, err := h.services.Role.GetByIdEmployee(currentEmplId)

	if roleDromDB.Role == "manager" {
		Tpl.ExecuteTemplate(c.Writer, "manager_homepage.html", nil)
	} else {
		Tpl.ExecuteTemplate(c.Writer, "cashier_homepage.html", nil)
	}
}

func (h *Handler) signInInput(c *gin.Context) {
	Tpl.ExecuteTemplate(c.Writer, "login.html", nil)
}
