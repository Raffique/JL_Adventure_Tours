package routes

import (
	"github.com/Raffique/JL_Adventure_Tours/Server/controllers"
	"github.com/Raffique/JL_Adventure_Tours/Server/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    auth := router.Group("/auth")
    {
        auth.POST("/signup", controllers.SignUp)
        auth.POST("/login", controllers.Login)
    }

    bookings := router.Group("/bookings")
    bookings.Use(middlewares.AuthMiddleware())
    {
		bookings.GET("/", controllers.GetBookings)
        bookings.POST("/", controllers.CreateBooking)
        bookings.GET("/:id", controllers.GetBookingByID)
        bookings.PUT("/:id", controllers.UpdateBooking)
        bookings.DELETE("/:id", controllers.DeleteBooking)
    }

    admin := router.Group("/admin")
    admin.Use(middlewares.AuthMiddleware(), middlewares.RoleMiddleware("admin"))
    {
        admin.GET("/bookings", controllers.GetBookings)
        admin.POST("/bookings", controllers.CreateBooking)
		admin.PUT("/bookings/:id", controllers.UpdateBooking)
        admin.DELETE("/bookings/:id", controllers.DeleteBooking)
        // Add other admin routes here for managing users, bookings, etc.
    }

	superAdmin := router.Group("/super-admin")
    superAdmin.Use(middlewares.AuthMiddleware(), middlewares.RoleMiddleware("super_admin"))
    {
        // Only accessible to super admins
        superAdmin.POST("/users", controllers.CreateUser)      // Create a new user
        superAdmin.GET("/users", controllers.GetUsers)         // Retrieve all users
        superAdmin.GET("/users/:id", controllers.GetUser)      // Retrieve a single user
        superAdmin.PUT("/users/:id", controllers.UpdateUser)   // Update a user's information
        superAdmin.DELETE("/users/:id", controllers.DeleteUser) // Delete a user
    }

    payment := router.Group("/payment")
    payment.Use(middlewares.AuthMiddleware())
    {
        payment.POST("/create", controllers.CreatePayment)
    }
}
