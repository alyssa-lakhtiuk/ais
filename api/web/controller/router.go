package controller

import (
	"ais/web/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) NewRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("employees", h.getAllEmployees)
	router.POST("employee", h.createEmployee)
	router.PUT("employee/:id", h.updateEmployee)
	router.DELETE("employee/:id", h.deleteEmployee)
	// Category pages
	router.POST("category/", h.createCategory)
	router.GET("categories/", h.getAllCategories)
	router.DELETE("category/:name", h.deleteCategory)
	router.PUT("category/:id", h.updateCategory)
	// Product pages
	router.POST("product", h.createProduct)
	router.GET("products", h.getAllProducts)
	router.DELETE("product/:id", h.deleteProduct)
	router.PUT("product/:id", h.updateProduct)
	// Customer card pages
	router.POST("customer-card", h.createCustomerCard)
	router.GET("customer-cards", h.getAllCustomerCards)
	router.DELETE("customer-cards/:id", h.deleteCustomerCard)
	router.PUT("customer-card/:id", h.updateCustomerCard)
	// Sale pages
	// Check pages
	router.POST("check", h.createCheck)
	router.GET("checks", h.getAllChecks)
	router.DELETE("check/:id", h.deleteCheck)
	// Store product pages
	router.POST("store-product", h.createStoreProduct)
	router.GET("store-products", h.getAllStoreProducts)
	router.DELETE("store-product/:id", h.deleteStoreProduct)
	router.PUT("store-product/:id", h.updateStoreProduct)

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
