package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Ustawienie połączenia z bazą danych
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migracje modeli
	err = db.AutoMigrate(&Product{}, &Category{}, &Cart{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Rejestracja routerów
	RegisterProductRoutes(e, db)
	RegisterCartRoutes(e, db)
	RegisterPaymentRoutes(e)

	// Start serwera
	e.Logger.Fatal(e.Start(":8080"))
}
