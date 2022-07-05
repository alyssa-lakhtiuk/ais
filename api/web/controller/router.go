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
	router.GET("/manager_homepage", h.signIn)

	router.GET("/employees", h.getAllEmployees)      //
	router.GET("/create-employee", h.createEmployee) //
	router.POST("/employee", h.employeeCreated)      // не відкривається
	router.GET("/employee/:id", h.getEmployeeByName) // by name  // з Brad не працює
	router.GET("/employees/:id", h.getEmployeeById)  // by id  // помилка
	router.PUT("/employee/:id", h.updateEmployee)
	router.DELETE("/employee/:id", h.deleteEmployee)
	// Category pages
	router.GET("/create-category", h.createCategory)     // html
	router.POST("/category", h.categoryCreated)          // html
	router.GET("/category/:id", h.getCategoryByName)     // by name // працює
	router.GET("/categories/:id", h.getCategoryByNumber) // by number  // кидає помилку
	router.GET("/categories", h.getAllCategories)        // html
	router.DELETE("/category/:name", h.deleteCategory)
	router.PUT("/category/:id", h.updateCategory)
	// Product pages
	router.GET("/create-product", h.createProduct)
	router.POST("/product", h.productCreated)         // випадний список сука
	router.GET("/products", h.getAllProducts)         // html
	router.GET("/product/:name", h.getProductByName)  // by name // працює
	router.GET("/products/:id", h.getProductByNumber) // by id // працює
	router.DELETE("/product/:id", h.deleteProduct)
	router.PUT("/product/:id", h.updateProduct)
	// Customer card pages
	router.GET("/create-customer-card", h.createCustomerCard)   // html
	router.POST("/customer-card", h.customerCardCreated)        //html
	router.GET("/customer-card/:id", h.getCustomerCardByNumber) // by name // працює ++
	router.GET("/customer-cards", h.getAllCustomerCards)        // html
	router.DELETE("/customer-cards/:id", h.deleteCustomerCard)
	router.PUT("/customer-card/:id", h.updateCustomerCard)
	// Store product pages
	router.GET("/create-store-product", h.createStoreProduct) // треба випадний список
	router.POST("/store-product", h.storeProductCreated)
	router.GET("/store-products", h.getAllStoreProducts)       // null
	router.GET("/store-products/:upc", h.getStoreProductByUpc) // by upc // не працює, потестити
	router.DELETE("/store-product/:id", h.deleteStoreProduct)
	router.PUT("/store-product/:id", h.updateStoreProduct)
	// Sale pages
	// Check pages
	router.POST("/check", h.createCheck)
	router.GET("/checks", h.getAllChecks)
	router.DELETE("/check/:id", h.deleteCheck)

	////router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	//auth := router.Group("/auth")
	//{
	//	auth.POST("/sign-up", h.signUp)
	//	auth.POST("/sign-in", h.signIn)
	//}
	//
	//api := router.Group("/api", h.userIdentity)
	//{
	//	lists := api.Group("/lists")
	//	{
	//		lists.POST("/", h.createList)
	//		lists.GET("/", h.getAllLists)
	//		lists.GET("/:id", h.getListById)
	//		lists.PUT("/:id", h.updateList)
	//		lists.DELETE("/:id", h.deleteList)
	//
	//		items := lists.Group(":id/items")
	//		{
	//			items.POST("/", h.createItem)
	//			items.GET("/", h.getAllItems)
	//		}
	//	}
	//
	//	items := api.Group("items")
	//	{
	//		items.GET("/:id", h.getItemById)
	//		items.PUT("/:id", h.updateItem)
	//		items.DELETE("/:id", h.deleteItem)
	//	}
	//}

	return router
}
