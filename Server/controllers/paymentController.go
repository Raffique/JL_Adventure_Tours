package controllers

import (
	"net/http"
	"server/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
    amountStr := c.Query("amount")
    amount, err := strconv.ParseInt(amountStr, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
        return
    }

    paymentIntent, err := services.CreatePaymentIntent(amount, "usd")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment intent"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"client_secret": paymentIntent.ClientSecret})
}
