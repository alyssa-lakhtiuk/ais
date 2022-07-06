package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
)

// це тільки контролер, який колає відповідні сервіси!
// мейбі навіть зробити тут один файлик

// ввести відомості про новий товар
// оновити відомості про товар
// вилучити відомості про товар

func (h *Handler) productCreated(c *gin.Context) {
	var input entities.Product
	input.Name = c.Request.FormValue("name_prod")
	input.CategoryNum, _ = strconv.Atoi(c.Request.FormValue("category"))
	input.Characteristics = c.Request.FormValue("description")
	//if err := c.BindJSON(&input); err != nil {
	//	// throw error response
	//	respondWithError(c, http.StatusBadRequest, "unable to get input data")
	//}
	_, err := h.services.Product.Create(input)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to create product, check is your data correct")
	}
	Tpl.ExecuteTemplate(c.Writer, "done_product.html", input)
}

func (h *Handler) createProduct(c *gin.Context) {
	var categories []entities.Category

	var err error
	categories, err = h.services.Category.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	Tpl.ExecuteTemplate(c.Writer, "add_product.html", categories)
}

func (h *Handler) getAllProducts(c *gin.Context) {
	authHeader, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "authorization first", nil)
	}
	currentEmplId := authHeader.Value
	roleDromDB, err := h.services.Role.GetByIdEmployee(currentEmplId)

	products, err := h.services.Product.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
		// throw error response
	}
	categories, err := h.services.Category.GetAll()
	var cpr entities.CategoriesProducts
	cpr.C = categories
	cpr.P = products
	//c.JSON(http.StatusOK, products)
	sort.SliceStable(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})
	if roleDromDB.Role == "manager" {
		Tpl.ExecuteTemplate(c.Writer, "manager_product.html", cpr)
	} else {
		Tpl.ExecuteTemplate(c.Writer, "cashier_product.html", products)
	}

}

func (h *Handler) getProductByName(c *gin.Context) {
	productName := c.Param("name")
	//productName := "Parmesan"
	product, err := h.services.Product.GetByName(productName)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to get this product")
		//return
		// throw error response
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) getProductByNumber(c *gin.Context) {
	productName, err := strconv.Atoi(c.Param("id"))
	//productName := 976
	product, err := h.services.Product.GetByNumber(productName)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to get this product")
		//return
		// throw error response
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) deleteProduct(c *gin.Context) {
	//id, err := strconv.Atoi(c.Param("id"))
	id, err := strconv.Atoi(c.Request.FormValue("id"))
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to get id of product")
		return
	}
	err = h.services.Product.Delete(id)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "unable to delete product, maybe it doesn't exist") //c.JSON(http.StatusBadRequest, err)
		return
	}
	h.getAllProducts(c)
	//c.JSON(http.StatusOK, id)
}

var updateProductNum int

func (h *Handler) updateProductOpen(c *gin.Context) {
	updateProductNum, _ = strconv.Atoi(c.Request.FormValue("id"))
	productToEdit, _ := h.services.Product.GetByNumber(updateProductNum)

	var categories []entities.Category
	var err error
	categories, err = h.services.Category.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	var combo entities.CategoryProduct
	combo.C = categories
	combo.P = productToEdit
	Tpl.ExecuteTemplate(c.Writer, "edit_product.html", combo)
}
func (h *Handler) updateProduct(c *gin.Context) {
	var input entities.Product
	input.Name = c.Request.FormValue("name_prod")
	input.CategoryNum, _ = strconv.Atoi(c.Request.FormValue("category"))
	input.Characteristics = c.Request.FormValue("Characteristics")
	if err := h.services.Product.Update(updateProductNum, input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to update, check is your data correct")
		return
	}
	//c.JSON(http.StatusOK, "updated")
	h.getAllProducts(c)
	//Tpl.ExecuteTemplate(c.Writer, "edit_product.html", entities.Message{Mess: "employee updated"})
}

func (h *Handler) onlyOneProductCategory(c *gin.Context) {
	role := c.Request.FormValue("sort_cat")
	//if role != "manager" && role != "cashier" {
	//	h.getAllEmployees(c)
	//	//Tpl.ExecuteTemplate(c.Writer, "manager_employee.html", allEmployees)
	//}
	products, err := h.services.Product.GetAllByCategory(role)
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "can't get all employees, maybe there is no even one")
		//return
	}
	//c.JSON(http.StatusOK, employees)
	sort.SliceStable(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})
	categories, err := h.services.Category.GetAll()
	var cpr entities.CategoriesProducts
	cpr.C = categories
	cpr.P = products
	Tpl.ExecuteTemplate(c.Writer, "manager_employee.html", cpr)
}
