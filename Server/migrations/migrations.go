package migrations

import (
	"log"
	"server/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        {
            ID: "20231110_create_users",
            Migrate: func(tx *gorm.DB) error {
                return tx.AutoMigrate(&models.User{})
            },
            Rollback: func(tx *gorm.DB) error {
                return tx.Migrator().DropTable("users")
            },
        },
        {
            ID: "20231110_create_bookings",
            Migrate: func(tx *gorm.DB) error {
                return tx.AutoMigrate(&models.Booking{})
            },
            Rollback: func(tx *gorm.DB) error {
                return tx.Migrator().DropTable("bookings")
            },
        },
    })

    if err := m.Migrate(); err != nil {
        log.Fatalf("Could not migrate: %v", err)
    }
    log.Println("Database migrated successfully")
}
