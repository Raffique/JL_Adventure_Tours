package middlewares

import (
	"net/http"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        claims, err := utils.VerifyToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Check if the user's role meets or exceeds the required role
        userRole := claims["role"].(string)

        if !isRoleAllowed(userRole, requiredRole) {
            c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
            c.Abort()
            return
        }

        c.Next()
    }
}

// Helper function to determine if the user's role meets the required role
func isRoleAllowed(userRole, requiredRole string) bool {
    roles := map[string]int{
        "customer":   1,
        "admin":      2,
        "super_admin": 3,
    }

    // Check if user's role level is equal to or higher than required role
    return roles[userRole] >= roles[requiredRole]
}
