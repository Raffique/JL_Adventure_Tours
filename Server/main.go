package main

import (
	"github.com/Raffique/JL_Adventure_Tours/Server/config"
	"github.com/Raffique/JL_Adventure_Tours/Server/controllers"
	"github.com/Raffique/JL_Adventure_Tours/Server/migrations"
	"github.com/Raffique/JL_Adventure_Tours/Server/routes"
	"github.com/Raffique/JL_Adventure_Tours/Server/services"
	"github.com/Raffique/JL_Adventure_Tours/Server/storage"
	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadEnv()
    config.ConnectDatabase()
	storage.InitStorage()
    services.InitStripe()

    db := config.DB
    migrations.Migrate(db)

    controllers.InitBookingController(db)

    router := gin.Default()
    routes.SetupRoutes(router)

    router.Run(":8080")
}
