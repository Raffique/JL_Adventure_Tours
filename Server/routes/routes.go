package routes

import (
	"server/controllers"
	"server/middlewares"

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
        superAdmin.POST("/users", controllers.CreateUser) // Example for user management
        superAdmin.DELETE("/users/:id", controllers.DeleteUser)
        // More super admin-specific routes can go here
    }

    payment := router.Group("/payment")
    payment.Use(middlewares.AuthMiddleware())
    {
        payment.POST("/create", controllers.CreatePayment)
    }
}
