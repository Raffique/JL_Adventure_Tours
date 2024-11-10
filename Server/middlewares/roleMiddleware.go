package middlewares

import (
	"net/http"
	"strings"

	"github.com/Raffique/JL_Adventure_Tours/Server/utils"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        token = strings.Split(token, "Bearer ")[1]
        claims, err := utils.VerifyToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Access role directly from the claims struct
        userRole := claims.Role

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
