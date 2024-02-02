package rest

import (
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controllerCustomer struct {
	service domains.CustomerService
}

func (cc *controllerCustomer) GetCustomerController(c echo.Context) error {
	customers := cc.service.GetCustomer()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     http.StatusOK,
		"customer": customers,
	})
}

func (cc *controllerCustomer) CreateCustomerController(c echo.Context) error {
	customer := models.Customer{}
	c.Bind(&customer)
	err := cc.service.CreateCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success",
	})
}

func (cc *controllerCustomer) DeleteCustomerController(c echo.Context) error {
	customer_id, _ := strconv.Atoi(c.Param("id"))
	err := cc.service.DeleteCustomer(int(customer_id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":  http.StatusOK,
		"pesan": "sukses",
	})
}

func (cc *controllerCustomer) UpdateCustomer(c echo.Context) error {
	customer_id, _ := strconv.Atoi(c.Param("id"))
	if customer_id == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": "Id 0",
		})
	}
	customer := models.Customer{}
	c.Bind(&customer)
	err := cc.service.UpdateCustomer(int(customer_id), customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":  http.StatusOK,
		"pesan": "sukses",
	})
}
