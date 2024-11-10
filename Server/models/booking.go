package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// StringArray is a custom type for handling string arrays
type StringArray []string

// Value implements the driver.Valuer interface
func (a StringArray) Value() (driver.Value, error) {
    if len(a) == 0 {
        return "{}", nil
    }
    // Escape single quotes and backslashes
    escaped := make([]string, len(a))
    for i, v := range a {
        escaped[i] = strings.ReplaceAll(strings.ReplaceAll(v, "\\", "\\\\"), "'", "\\'")
    }
    return "{" + strings.Join(escaped, ",") + "}", nil
}

// Scan implements the sql.Scanner interface for StringArray
func (a *StringArray) Scan(value interface{}) error {
    if value == nil {
        *a = StringArray{}
        return nil
    }

    var str string
    switch v := value.(type) {
    case []byte:
        str = string(v)
    case string:
        str = v
    default:
        return fmt.Errorf("unsupported type: %T", value)
    }

    // Split the string into array elements
    *a = StringArray(strings.Split(strings.Trim(str, "{}"), ","))
    return nil
}


type Booking struct {
    gorm.Model
    UserID   uint           `json:"user_id"`
    Title    string         `json:"title"`      // Title or description of the booking
    Price    float64        `json:"price"`      // Price of the booking item
    Pictures StringArray       `gorm:"type:text[]" json:"pictures"` // List of image URLs
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
    Pictures  StringArray  `gorm:"type:text[]" json:"pictures"`
    Periods   string    `json:"periods"`     // Serialized JSON string of booking periods
    Status    string    `json:"status"`
    Action    string    `json:"action"`      // Action performed: created, updated, deleted
    Timestamp time.Time `json:"timestamp"`   // When the action was performed
}
