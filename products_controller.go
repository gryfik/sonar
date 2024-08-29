package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

// Definicja stałych dla ścieżek
const (
	productBasePath = "/products"
	productIDPath   = "/products/:id"
)

func RegisterProductRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET(productBasePath, func(c echo.Context) error { return getProducts(c, db) })
	e.POST(productBasePath, func(c echo.Context) error { return createProduct(c, db) })
	e.GET(productIDPath, func(c echo.Context) error { return getProduct(c, db) })
	e.PUT(productIDPath, func(c echo.Context) error { return updateProduct(c, db) })
	e.DELETE(productIDPath, func(c echo.Context) error { return deleteProduct(c, db) })
	e.DELETE(productBasePath, func(c echo.Context) error { return deleteProducts(c, db) })
}

func getProducts(c echo.Context, db *gorm.DB) error {
	var products []Product
	db.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func createProduct(c echo.Context, db *gorm.DB) error {
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request data"})
	}
	if product.Name == "" || len(product.Name) < 2 || len(product.Name) > 100 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name must be at least 2 and no more than 100 characters long"})
	}
	if product.Description == "" || len(product.Description) < 10 || len(product.Description) > 1000 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Description must be at least 10 and no more than 1000 characters long"})
	}
	if product.Price <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Price must be greater than 0"})
	}
	result := db.Create(product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}
	return c.JSON(http.StatusCreated, product)
}

func getProduct(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	var product Product
	result := db.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	var product Product
	if db.First(&product, id).Error != nil {
		return c.JSON(http.StatusNotFound, "Invalid product ID")
	}
	if err := c.Bind(&product); err != nil {
		return err
	}
	if product.Name == "" || len(product.Name) < 2 || len(product.Name) > 100 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name must be at least 2 and no more than 100 characters long"})
	}
	if product.Description == "" || len(product.Description) < 10 || len(product.Description) > 1000 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Description must be at least 10 and no more than 1000 characters long"})
	}
	if product.Price <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Price must be greater than 0"})
	}
	db.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context, db *gorm.DB) error {
	id := c.Param("id")
	var product Product
	result := db.First(&product, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	result = db.Delete(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	return c.JSON(http.StatusOK, "Product deleted")
}

func deleteProducts(c echo.Context, db *gorm.DB) error {
	result := db.Where("1 = 1").Delete(&Product{})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete products")
	}
	return c.JSON(http.StatusOK, "All products deleted")
}
