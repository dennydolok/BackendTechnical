package rest

import (
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controllerOrder struct {
	service domains.OrderService
}

func (co *controllerOrder) GetOrderById(c echo.Context) error {
	order_id, _ := strconv.Atoi(c.Param("id"))
	order := co.service.GetOrderById(order_id)
	if (models.Order{}) == order {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"kode":    http.StatusNotFound,
			"message": "order not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"kode":  http.StatusOK,
		"order": order,
	})
}

func (co *controllerOrder) CreateOrder(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)
	err := co.service.CreateOrder(order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Error",
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "success",
	})
}
