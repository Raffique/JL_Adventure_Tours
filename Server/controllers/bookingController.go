package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Raffique/JL_Adventure_Tours/Server/models"
	"github.com/Raffique/JL_Adventure_Tours/Server/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var bookingRepo *repositories.BookingRepository

func InitBookingController(db *gorm.DB) {
    bookingRepo = &repositories.BookingRepository{DB: db}
}

func formatPostgresArray(arr []string) string {
    return "{" + strings.Join(arr, ",") + "}"
}

func CreateBooking(c *gin.Context) {
    var booking models.Booking
    if err := c.ShouldBindJSON(&booking); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := bookingRepo.CreateBooking(&booking); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
        return
    }

    c.JSON(http.StatusOK, booking)
}

func GetBookings(c *gin.Context) {
    userRole := c.MustGet("role").(string)  // Retrieve the role from context
    userID := c.MustGet("user_id").(uint)   // Retrieve the user ID from context (if needed)

    var bookings []models.Booking
    var err error

    if userRole == "customer" {
        // Customers only see their own bookings
        bookings, err = bookingRepo.GetBookingsByUserID(userID)
    } else {
        // Admins and Super Admins see all bookings
        bookings, err = bookingRepo.GetBookings()
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookings"})
        return
    }

    c.JSON(http.StatusOK, bookings)
}

func GetBookingByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

    booking, err := bookingRepo.GetBookingByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
        return
    }

    c.JSON(http.StatusOK, booking)
}

func UpdateBooking(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

    var updatedBooking models.Booking
    if err := c.ShouldBindJSON(&updatedBooking); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := bookingRepo.UpdateBooking(uint(id), &updatedBooking); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Booking updated successfully"})
}

func DeleteBooking(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

    if err := bookingRepo.DeleteBooking(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}