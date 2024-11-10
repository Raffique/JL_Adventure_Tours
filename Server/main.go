package main

import (
	"log"

	"github.com/Raffique/JL_Adventure_Tours/Server/config"
	"github.com/Raffique/JL_Adventure_Tours/Server/controllers"
	"github.com/Raffique/JL_Adventure_Tours/Server/migrations"
	"github.com/Raffique/JL_Adventure_Tours/Server/models"
	"github.com/Raffique/JL_Adventure_Tours/Server/routes"
	"github.com/Raffique/JL_Adventure_Tours/Server/services"
	"github.com/Raffique/JL_Adventure_Tours/Server/storage"
	"github.com/Raffique/JL_Adventure_Tours/Server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
    config.LoadEnv()
    config.ConnectDatabase()
	storage.InitStorage()
    services.InitStripe()

    db := config.DB
    migrations.Migrate(db)

    initializeSuperAdmin(db)

    controllers.InitBookingController(db)

    router := gin.Default()
    routes.SetupRoutes(router)

    router.Run(":8080")
}

func initializeSuperAdmin(db *gorm.DB) {
    var user models.User
    if err := db.Where("role = ?", "super_admin").First(&user).Error; err == gorm.ErrRecordNotFound {
        // Super admin doesn't exist, create a new one
        password := "password1234"
        hashedPassword, _ := utils.HashPassword(password)

        superAdmin := models.User{
            Username: "superadmin",
            Password: hashedPassword,
            Role:     "super_admin",
        }

        if err := db.Create(&superAdmin).Error; err != nil {
            log.Fatalf("Failed to create super admin: %v", err)
        }

        log.Println("Super admin user created successfully.")
    } else if err != nil {
        log.Fatalf("Failed to check super admin existence: %v", err)
    } else {
        log.Println("Super admin user already exists.")
    }
}