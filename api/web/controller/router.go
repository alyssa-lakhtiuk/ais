package controller

import (
	"ais/web/service"
	"github.com/gin-gonic/gin"
	"html/template"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func (h *Handler) NewRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/sign-in-input", h.signInInput)
	router.POST("/sign-in", h.signIn)
	router.GET("/homepage", h.goToHomePage)

	router.GET("/employees", h.getAllEmployees)      //
	router.GET("/create-employee", h.createEmployee) //
	router.POST("/employee", h.employeeCreated)      //
	router.GET("/employee/:id", h.getEmployeeByName) // by name  // з Brad не працює
	router.GET("/employees/:id", h.getEmployeeById)  // by id  // помилка
	router.POST("/edit-employee", h.updateEmployee)
	router.POST("/request-employee", h.onlyOneEmployeeCategory)
	router.GET("/edit-employee", h.updateEmployeeOpen)
	// Reports
	router.GET("/report-employee", h.createEmployeeReport)
	router.GET("/report-customer", h.createCustomerReport)
	router.GET("/report-product", h.createProductReport)
	router.GET("/report-category", h.createCategoryReport)
	router.GET("/report-check", h.createCheckReport)
	router.GET("/report-store-product", h.createStoreReport)
	//router.DELETE("/employee/:id", h.deleteEmployee)
	router.DELETE("/delete-employee", h.deleteEmployee)
	//router.DELETE("/employee/:id", h.deleteEmployee)
	//router.POST("/delete-employee", h.deleteEmployee)
	//router.GET("/who-am-i", h.WhoAmI)
	// Category pages
	router.GET("/create-category", h.createCategory)
	router.POST("/category", h.categoryCreated)
	router.GET("/category/:id", h.getCategoryByName)
	router.GET("/categories/:id", h.getCategoryByNumber)
	router.GET("/categories", h.getAllCategories)
	router.DELETE("/category/:name", h.deleteCategory)
	//router.PUT("/category/:id", h.updateCategory)
	router.PUT("/edit-category", h.updateCategory)
	router.GET("/edit-category", h.updateCategoryOpen)
	router.DELETE("/delete-category", h.deleteCategory)
	router.GET("/delete-category", h.deleteCategory)
	// Product pages
	router.GET("/create-product", h.createProduct)
	router.POST("/product", h.productCreated)
	router.GET("/products", h.getAllProducts)
	router.GET("/product/:name", h.getProductByName)
	router.GET("/products/:id", h.getProductByNumber)
	router.DELETE("/delete-product", h.deleteProduct)
	//router.PUT("/product/:id", h.updateProduct)
	router.POST("/edit-product", h.updateProduct)
	router.GET("/edit-product", h.updateProductOpen)
	router.POST("/delete-product", h.deleteProduct)
	// Customer card pages
	router.GET("/create-customer-card", h.createCustomerCard)   // html
	router.POST("/customer-card", h.customerCardCreated)        //html
	router.GET("/customer-card/:id", h.getCustomerCardByNumber) // by name // працює ++
	router.GET("/customer-cards", h.getAllCustomerCards)        // html
	router.DELETE("/customer-card", h.deleteCustomerCard)
	//router.PUT("/customer-card/:id", h.updateCustomerCard)
	router.PUT("/edit-customer-card", h.updateCustomerCard)
	router.GET("/edit-customer-card", h.updateCustomerCardOpen)
	router.POST("/delete-customer-card", h.deleteCustomerCard)
	// Store product pages
	router.GET("/create-store-product", h.createStoreProduct) // треба випадний список
	router.POST("/store-product", h.storeProductCreated)
	router.GET("/store-products", h.getAllStoreProducts)       // html
	router.GET("/store-products/:upc", h.getStoreProductByUpc) // by upc
	router.DELETE("/delete-store-product", h.deleteStoreProduct)
	//router.PUT("/store-product/:id", h.updateStoreProduct)
	router.POST("/edit-store-product", h.updateStoreProduct)
	router.GET("/edit-store-product", h.updateStoreProductOpen)
	router.POST("/delete-store-product", h.deleteStoreProduct)
	router.POST("/store-products-upc", h.searchStoreProductsUPC)
	// Sale pages
	// Check pages
	router.POST("/check", h.createCheck)
	router.GET("/checks", h.getAllChecks)
	router.DELETE("/check/:id", h.deleteCheck)
	router.GET("/add-product-to-check", h.addProductToCheckOpen)
	router.POST("/add-product-to-check", h.addProductToCheck)
	router.GET("/who-am-i", h.whoAmI)

	// Zvit
	router.GET("/quantity-category", h.quantitiesByCat)
	router.POST("/checks-category", h.checksByCat)
	///////
	router.GET("/count-cities", h.countCities)
	router.POST("/checks-price", h.checksByPrice)
	return router
}
