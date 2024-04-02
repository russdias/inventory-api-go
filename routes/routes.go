package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/russelldias98/go-api/controllers"
)

func SetupRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/suppliers/{supplierId}/products/{productId}/stock", controllers.CheckStockForProduct)
	v1.GET("/products/out-of-stock", controllers.ListOutOfStockProducts)
	v1.GET("/products/low-stock", controllers.ListLowStockProducts)

	supplierRoutes := v1.Group("/suppliers")
	supplierRoutes.GET("/", controllers.GetSuppliers)
	supplierRoutes.GET("/products/:supplierId", controllers.GetProductsBySupplierID)
	supplierRoutes.GET("/:id", controllers.GetSupplierByID)
	supplierRoutes.POST("/", controllers.CreateSupplier)
	supplierRoutes.DELETE("/:id", controllers.DeleteSupplierByID)
	supplierRoutes.PUT("/:id", controllers.UpdateSupplierByID)

	categoryRoutes := v1.Group("/categories")
	categoryRoutes.GET("/", controllers.GetCategories)
	categoryRoutes.GET("/:id", controllers.GetCategoryByID)
	categoryRoutes.POST("/", controllers.CreateCategory)
	categoryRoutes.DELETE("/:id", controllers.DeleteCategoryByID)
	categoryRoutes.PUT("/:id", controllers.UpdateCategoryByID)

	productRoutes := v1.Group("/products")
	productRoutes.GET("/", controllers.GetProducts)
	productRoutes.GET("/:id", controllers.GetProductByID)
	productRoutes.POST("/", controllers.CreateProduct)
	productRoutes.DELETE("/:id", controllers.DeleteProductByID)
	productRoutes.PUT("/:id", controllers.UpdateProductByID)

	err := router.Run(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
