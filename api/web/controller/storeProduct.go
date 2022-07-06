package controller

import (
	"ais/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
)

func (h *Handler) storeProductCreated(c *gin.Context) {
	var input entities.StoreProduct
	var err error
	//if err := c.BindJSON(&input); err != nil {
	//	// throw error response
	//	respondWithError(c, http.StatusBadRequest, "cant process data, check it")
	//	return
	//}
	input.UPC = c.Request.FormValue("upc")
	input.IDProduct, err = strconv.Atoi(c.Request.FormValue("product"))
	input.SellingPrice, err = strconv.ParseFloat(c.Request.FormValue("selling_price"), 64)
	input.ProductsNumber, err = strconv.Atoi(c.Request.FormValue("quantit"))
	input.PromotionalProduct, err = strconv.ParseBool(c.Request.FormValue("promotion"))
	//input.UPCProm = c.Request.FormValue("promotion")
	/// !!!!!!!!!!!!!! upcprom !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!1
	///!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	_, err = h.services.StoreProduct.Create(input)
	if err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "fail to create store product")
		return
	}
	h.getAllStoreProducts(c)
	//Tpl.ExecuteTemplate(c.Writer, "")
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"productId": input.UPC,
	//	"id":        id,
	//})
}

func (h *Handler) createStoreProduct(c *gin.Context) {
	var products []entities.Product
	var err error
	products, err = h.services.Product.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	Tpl.ExecuteTemplate(c.Writer, "add_stock_product.html", products)
}

func (h *Handler) getAllStoreProducts(c *gin.Context) {
	authHeader, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "authorization first", nil)
	}
	currentEmplId := authHeader.Value
	roleDromDB, err := h.services.Role.GetByIdEmployee(currentEmplId)

	products, err := h.services.StoreProduct.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	sort.SliceStable(products, func(i, j int) bool {
		return products[i].ProductsNumber < products[j].ProductsNumber
	})
	if roleDromDB.Role == "manager" {
		Tpl.ExecuteTemplate(c.Writer, "manager_stock_product.html", products)
	} else {
		Tpl.ExecuteTemplate(c.Writer, "cashier_stock_product.html", products)
	}

	//c.JSON(http.StatusOK, products)
}

func (h *Handler) getStoreProductByUpc(c *gin.Context) {
	upc := c.Param("upc")
	category, err := h.services.StoreProduct.GetByName(upc)
	if err != nil {
		// throw error response
		//respondWithError(c, http.StatusBadRequest, "unable to get store product")
		//return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Handler) deleteStoreProduct(c *gin.Context) {
	//id := c.Param("id")
	id := c.Request.FormValue("upc_id")
	err := h.services.StoreProduct.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	h.getAllStoreProducts(c)
	//c.JSON(http.StatusOK, "deleted")
}

func (h *Handler) updateStoreProductOpen(c *gin.Context) {
	storeProductUPCForUPD := c.Request.FormValue("upc_id")
	stProdToEdit, _ := h.services.StoreProduct.GetByName(storeProductUPCForUPD)

	products, err := h.services.Product.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
		// throw error response
	}
	Stproducts, err := h.services.StoreProduct.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	var combo entities.ProductStoreProductUpc
	combo.Stp = stProdToEdit
	combo.Pr = products
	combo.UpC = Stproducts

	Tpl.ExecuteTemplate(c.Writer, "edit_stock_product.html", combo)
}

func (h *Handler) updateStoreProduct(c *gin.Context) {
	//upc := c.Param("id")
	var input entities.StoreProduct
	upc := c.Request.FormValue("upc")
	input.IDProduct, _ = strconv.Atoi(c.Request.FormValue("product"))
	input.SellingPrice, _ = strconv.ParseFloat(c.Request.FormValue("selling_price"), 64)
	input.ProductsNumber, _ = strconv.Atoi(c.Request.FormValue("number"))
	input.PromotionalProduct, _ = strconv.ParseBool(c.Request.FormValue("promotion"))

	//if err := c.BindJSON(&input); err != nil {
	//	// throw error response
	//	respondWithError(c, http.StatusBadRequest, "unable to parse input data")
	//	return
	//}
	if err := h.services.StoreProduct.Update(upc, input); err != nil {
		// throw error response
		respondWithError(c, http.StatusBadRequest, "unable to update store product")
		return
	}
	//c.JSON(http.StatusOK, "updated")
	h.getAllStoreProducts(c)
	//Tpl.ExecuteTemplate(c.Writer, "manager_stock_product.html", entities.Message{Mess: "employee updated"})
}

func (h *Handler) searchStoreProductsUPC(c *gin.Context) {
	authHeader, err := c.Request.Cookie("Authorization")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "authorization first", nil)
	}
	upc := c.Request.FormValue("upc")
	currentEmplId := authHeader.Value
	roleDromDB, err := h.services.Role.GetByIdEmployee(currentEmplId)

	products, err := h.services.StoreProduct.SearchUPC(upc)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	sort.SliceStable(products, func(i, j int) bool {
		return products[i].ProductsNumber < products[j].ProductsNumber
	})
	if roleDromDB.Role == "manager" {
		Tpl.ExecuteTemplate(c.Writer, "manager_stock_product.html", products)
	} else {
		Tpl.ExecuteTemplate(c.Writer, "cashier_stock_product.html", products)
	}

	//c.JSON(http.StatusOK, products)
}

func (h *Handler) createStoreReport(c *gin.Context) {
	employees, _ := h.services.StoreProduct.GetAll()
	Tpl.ExecuteTemplate(c.Writer, "report_stock_product.html", employees)
}
