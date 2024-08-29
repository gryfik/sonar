package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func RegisterCartRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/cart", func(c echo.Context) error { return getCart(c, db) })
	e.PUT("/cart", func(c echo.Context) error { return updateCart(c, db) })
	e.DELETE("/cart", func(c echo.Context) error { return deleteCart(c, db) })
}

func getCart(c echo.Context, db *gorm.DB) error {
	var cart Cart
	db.Preload("Products").First(&cart)
	if cart.ID == 0 {
		return c.JSON(http.StatusNotFound, "Cart not found")
	}
	return c.JSON(http.StatusOK, cart)
}

func updateCart(c echo.Context, db *gorm.DB) error {
	cart := new(Cart)
	if err := c.Bind(cart); err != nil {
		return err
	}
	db.Save(&cart)
	return c.JSON(http.StatusOK, cart)
}

func deleteCart(c echo.Context, db *gorm.DB) error {
	result := db.Where("1 = 1").Delete(&Cart{})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete carts")
	}
	return c.JSON(http.StatusOK, "Cart deleted")
}
