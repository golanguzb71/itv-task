package middleware

import (
	"itv/internal/dto"
	"itv/pkg/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(jwtService *auth.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponseDTO{Error: "authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponseDTO{Error: "invalid authorization header format"})
			c.Abort()
			return
		}

		claims, err := m.jwtService.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponseDTO{Error: "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func (m *AuthMiddleware) RoleAuth(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponseDTO{Error: "unauthorized"})
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponseDTO{Error: "server error"})
			c.Abort()
			return
		}

		allowed := false
		for _, role := range roles {
			if roleStr == role {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, dto.ErrorResponseDTO{Error: "forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
