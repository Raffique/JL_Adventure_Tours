package models

import "gorm.io/gorm"

type Booking struct {
    gorm.Model
    UserID    uint   `json:"user_id"`
    BookingType string `json:"booking_type"` // room, car, or trip
    Details    string `json:"details"`       // JSON with booking-specific details
    Status     string `json:"status"`        // booked, canceled, completed
}
