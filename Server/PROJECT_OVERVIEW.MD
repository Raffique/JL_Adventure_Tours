Project Overview

The backend is designed for a booking website where users (customers) can book rooms, cars, and trips. It allows customers to sign up, make bookings, and process payments. Admins have permissions to manage all bookings, and super admins have additional privileges, like managing users.
Key Features

    User Authentication & JWT:
        Users can sign up and log in, with their information stored securely in the PostgreSQL database.
        Authentication is handled using JWT, with tokens containing user ID and role data for access control.

    Role-Based Access Control (RBAC):
        Users are divided into roles: customer, admin, and super_admin.
        Customers can create and view their bookings.
        Admins can view and manage all bookings.
        Super Admins can manage both bookings and user accounts.
        Middleware checks the user’s role to control access to specific routes based on their permissions.

    Database Configuration & Migrations:
        Configured PostgreSQL connection in config/config.go.
        Used gormigrate for database migrations to create tables for User and Booking.
        Migrations are version-controlled, allowing for smooth updates and rollbacks of the database schema.

    Models:
        User Model: Defines user attributes, including username, password, and role.
        Booking Model: Defines booking attributes, including user_id, booking_type, details, and status.
        Relationships: Users have a one-to-many relationship with bookings.

    Repositories:
        The BookingRepository contains CRUD operations for bookings:
            CreateBooking, GetBookings, GetBookingByID, UpdateBooking, and DeleteBooking.
        An additional GetBookingsByUserID method restricts customers to only see their own bookings.

    Controllers:
        AuthController: Handles user signup and login, generating JWT tokens on successful login.
        BookingController: Handles CRUD operations for bookings, applying access restrictions based on user roles.
        PaymentController: Handles payment processing by creating a Stripe PaymentIntent.

    Middleware:
        AuthMiddleware: Verifies JWT tokens to ensure only authenticated users can access certain routes.
        RoleMiddleware: Restricts access based on the user’s role hierarchy (customer, admin, super admin).

    Stripe Integration for Payments:
        Integrated Stripe for payment processing, enabling customers to initiate payments when making a booking.
        The Stripe PaymentIntent API is used to create payment intents, and the client receives a client_secret for securely processing payments on the frontend.

    Routes:
        Configured routes in routes/routes.go with RBAC:
            /auth: Sign up and login routes.
            /bookings: Routes for creating and viewing bookings, restricted to authenticated customers.
            /admin: Admin-only routes for managing all bookings.
            /super-admin: Super admin-only routes for managing users and other admin operations.

    Environment Variables:
        .env file stores environment variables for database connection, JWT secret, and Stripe API key, keeping sensitive information secure.

    Docker and Docker Compose:
        Dockerfile and docker-compose.yml are set up to run the Gin backend and PostgreSQL in separate containers, making it easier to deploy and manage the application.


Project Directory Structure

Here’s a summary of the project structure:

```bash
booking_website/
├── config/                # Database configuration
│   └── config.go
├── controllers/           # Controllers for handling requests
│   ├── authController.go
│   ├── bookingController.go
│   └── paymentController.go
├── middlewares/           # Middleware for authentication and role-based access control
│   ├── authMiddleware.go
│   └── roleMiddleware.go
├── migrations/            # Database migrations
│   └── migrations.go
├── models/                # Database models
│   ├── user.go
│   └── booking.go
├── repositories/          # Database CRUD operations
│   └── bookingRepository.go
├── routes/                # API route definitions
│   └── routes.go
├── services/              # Business logic and payment integration
│   ├── bookingService.go
│   └── paymentService.go
├── utils/                 # Utility functions for JWT and hashing
│   ├── jwt.go
│   └── hash.go
├── .env                   # Environment variables
├── Dockerfile             # Dockerfile for the Gin backend
├── docker-compose.yml     # Docker Compose configuration
├── go.mod                 # Go module dependencies
├── go.sum                 # Dependency lock file
└── main.go                # Entry point of the application
```

Summary of Key Code Files

    config/config.go: Configures the database connection using environment variables.
    models/user.go & models/booking.go: Define the schema for users and bookings.
    controllers/authController.go: Manages user signup and login.
    controllers/bookingController.go: Handles booking CRUD operations with RBAC.
    controllers/paymentController.go: Integrates with Stripe to handle payments.
    middlewares/authMiddleware.go & roleMiddleware.go: Ensure that routes are accessible based on user authentication and roles.
    services/paymentService.go: Creates Stripe PaymentIntents for payments.
    routes/routes.go: Sets up routes with appropriate middleware for role-based access control.

Summary of Major Functionality

    Authentication and Authorization: JWT-based token authentication with role-based access control.
    Database CRUD Operations: GORM-based repository pattern for modular CRUD functionality.
    Role-Based Access Control: Fine-grained access control using role-based middleware.
    Stripe Payment Integration: Secure payment handling with Stripe PaymentIntents.
    Environment Configuration and Dockerization: Environment variables for sensitive data and Docker setup for seamless deployment.