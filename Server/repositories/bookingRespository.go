package repositories

import (
	"encoding/json"
	"time"

	"github.com/Raffique/JL_Adventure_Tours/Server/models"
	"gorm.io/gorm"
)

type BookingRepository struct {
    DB *gorm.DB
}

// Helper function to log booking history
func (r *BookingRepository) logHistory(booking *models.Booking, action string) error {
    periodsJSON, _ := json.Marshal(booking.Periods)

    history := models.BookingHistory{
        BookingID: booking.ID,
        UserID:    booking.UserID,
        Title:     booking.Title,
        Price:     booking.Price,
        Pictures:  booking.Pictures,
        Periods:   string(periodsJSON),
        Status:    booking.Status,
        Action:    action,
        Timestamp: time.Now(),
    }
    return r.DB.Create(&history).Error
}

func (r *BookingRepository) CreateBooking(booking *models.Booking) error {
    err := r.DB.Create(booking).Error
    if err == nil {
        r.logHistory(booking, "created")
    }
    return err
}

func (r *BookingRepository) GetBookings() ([]models.Booking, error) {
    var bookings []models.Booking
    err := r.DB.Find(&bookings).Error
    return bookings, err
}

// GetBookingByID retrieves a booking by its ID
func (r *BookingRepository) GetBookingByID(id uint) (*models.Booking, error) {
    var booking models.Booking
    err := r.DB.First(&booking, id).Error
    if err != nil {
        return nil, err
    }
    return &booking, nil
}

func (r *BookingRepository) GetBookingsByUserID(userID uint) ([]models.Booking, error) {
    var bookings []models.Booking
    err := r.DB.Where("user_id = ?", userID).Find(&bookings).Error
    return bookings, err
}


// UpdateBooking updates an existing booking by its ID
func (r *BookingRepository) UpdateBooking(id uint, updatedBooking *models.Booking) error {
    var booking models.Booking
    if err := r.DB.First(&booking, id).Error; err != nil {
        return err
    }

    booking.Title = updatedBooking.Title
    booking.Price = updatedBooking.Price
    booking.Pictures = updatedBooking.Pictures
    booking.Status = updatedBooking.Status
    booking.Periods = updatedBooking.Periods

    err := r.DB.Save(&booking).Error
    if err == nil {
        r.logHistory(&booking, "updated")
    }
    return err
}

// DeleteBooking deletes a booking by its ID
func (r *BookingRepository) DeleteBooking(id uint) error {
    var booking models.Booking
    if err := r.DB.First(&booking, id).Error; err != nil {
        return err
    }

    err := r.DB.Delete(&booking).Error
    if err == nil {
        r.logHistory(&booking, "deleted")
    }
    return err
}
