package main

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"time"
)

type PaymentRequest struct {
	CardNumber string  `json:"cardNumber"`
	ExpiryDate string  `json:"expiryDate"` // format MM/YY
	CVV        string  `json:"cvv"`
	Amount     float64 `json:"amount"`
}

func RegisterPaymentRoutes(e *echo.Echo) {
	e.POST("/payments", handlePayment)
}

func handlePayment(c echo.Context) error {
	var paymentRequest PaymentRequest
	if err := c.Bind(&paymentRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid payment data")
	}

	if len(paymentRequest.CardNumber) != 16 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid card number")
	}

	time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)

	if rand.Intn(2) == 0 {
		return c.JSON(http.StatusOK, map[string]string{"status": "success", "message": "Payment successful!"})
	} else {
		return echo.NewHTTPError(http.StatusPaymentRequired, "Payment failed")
	}
}
