package rest

import (
	"dennydolok/experimentals-1/domains"
	"dennydolok/experimentals-1/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type controllerProduct struct {
	service domains.ProductService
}

func (cp *controllerProduct) GetProductController(c echo.Context) error {
	products := cp.service.GetProduct()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     http.StatusOK,
		"customer": products,
	})
}

func (cp *controllerProduct) CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	err := cp.service.CreateProduct(product)
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

func (cp *controllerProduct) DeleteProductController(c echo.Context) error {
	product_id, _ := strconv.Atoi(c.Param("id"))
	err := cp.service.DeleteProduct(product_id)
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

func (cp *controllerProduct) UpdateProductController(c echo.Context) error {
	product_id, _ := strconv.Atoi(c.Param("id"))
	if product_id == 0 {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"kode":  http.StatusInternalServerError,
			"pesan": "Id 0",
		})
	}
	product := models.Product{}
	c.Bind(&product)
	err := cp.service.UpdateProduct(product_id, product)
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
