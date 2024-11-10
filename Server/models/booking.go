package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
    gorm.Model
    UserID   uint           `json:"user_id"`
    Title    string         `json:"title"`      // Title or description of the booking
    Price    float64        `json:"price"`      // Price of the booking item
    Pictures []string       `gorm:"type:text[]" json:"pictures"` // List of image URLs
    Periods  []BookingPeriod `gorm:"foreignKey:BookingID" json:"periods"` // List of booking periods
    Status   string         `json:"status"`     // Booked, available, etc.
    Details  string         `json:"details"`    // Additional details
}

type BookingPeriod struct {
    gorm.Model
    BookingID uint      `json:"booking_id"`  // Reference to the Booking
    Start     time.Time `json:"start"`       // Start of booking period
    End       time.Time `json:"end"`         // End of booking period
}

type BookingHistory struct {
    gorm.Model
    BookingID uint      `json:"booking_id"`  // Reference to the Booking
    UserID    uint      `json:"user_id"`     // Reference to the User who made the booking
    Title     string    `json:"title"`
    Price     float64   `json:"price"`
    Pictures  []string  `gorm:"type:text[]" json:"pictures"`
    Periods   string    `json:"periods"`     // Serialized JSON string of booking periods
    Status    string    `json:"status"`
    Action    string    `json:"action"`      // Action performed: created, updated, deleted
    Timestamp time.Time `json:"timestamp"`   // When the action was performed
}
