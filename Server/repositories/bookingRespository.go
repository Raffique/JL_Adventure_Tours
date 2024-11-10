package repositories

import (
	"server/models"

	"gorm.io/gorm"
)

type BookingRepository struct {
    DB *gorm.DB
}

func (r *BookingRepository) CreateBooking(booking *models.Booking) error {
    return r.DB.Create(booking).Error
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

    // Update the fields of the booking based on the updatedBooking data
    booking.BookingType = updatedBooking.BookingType
    booking.Details = updatedBooking.Details
    booking.Status = updatedBooking.Status

    return r.DB.Save(&booking).Error
}

// DeleteBooking deletes a booking by its ID
func (r *BookingRepository) DeleteBooking(id uint) error {
    return r.DB.Delete(&models.Booking{}, id).Error
}
