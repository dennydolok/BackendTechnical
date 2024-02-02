package rest

import (
	"dennydolok/experimentals-1/config"
	"dennydolok/experimentals-1/databases"
	"dennydolok/experimentals-1/repositories"
	"dennydolok/experimentals-1/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterAPI(e *echo.Echo, config config.Config) {
	database := databases.InitMysql(config)

	// Customer API
	repositoryCustomer := repositories.NewUserRepository(database)
	serviceCustomer := services.NewserviceCustomers(config, repositoryCustomer)
	controllerCustomer := controllerCustomer{
		service: serviceCustomer,
	}
	customerAPI := e.Group("/customer")
	customerAPI.GET("", controllerCustomer.GetCustomerController, middleware.Logger())
	customerAPI.POST("", controllerCustomer.CreateCustomerController, middleware.Logger())
	customerAPI.DELETE("/:id", controllerCustomer.DeleteCustomerController, middleware.Logger())
	customerAPI.PUT("/:id", controllerCustomer.UpdateCustomer)

	// Product API
	repositoryProduct := repositories.NewProductRepository(database)
	serviceProduct := services.NewserviceProduct(config, repositoryProduct)
	controllerProduct := controllerProduct{
		service: serviceProduct,
	}
	productAPI := e.Group("/product")
	productAPI.GET("", controllerProduct.GetProductController, middleware.Logger())
	productAPI.POST("", controllerProduct.CreateProductController, middleware.Logger())
	productAPI.DELETE("/:id", controllerProduct.DeleteProductController, middleware.Logger())
	productAPI.PUT("/:id", controllerProduct.UpdateProductController, middleware.Logger())

	// Order API
	repositoryOrder := repositories.NewOrderRepository(database)
	serviceOrder := services.NewserviceOrder(config, repositoryOrder)
	controllerOrder := controllerOrder{
		service: serviceOrder,
	}
	orderAPI := e.Group("/order")
	orderAPI.GET("/:id", controllerOrder.GetOrderById, middleware.Logger())
	orderAPI.POST("", controllerOrder.CreateOrder, middleware.Logger())
}
