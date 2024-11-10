package main

import (
	"server/config"
	"server/controllers"
	"server/migrations"
	"server/routes"
	"server/services"

	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    config.ConnectDatabase()
    services.InitStripe()

    db := config.DB
    migrations.Migrate(db)

    controllers.InitBookingController(db)

    router := gin.Default()
    routes.SetupRoutes(router)

    router.Run(":8080")
}
