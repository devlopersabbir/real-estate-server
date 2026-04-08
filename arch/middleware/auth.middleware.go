package middleware

import (
	"github.com/devlopersabbir/juan_don82-server/arch/networks"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/config"
	"github.com/devlopersabbir/juan_don82-server/internal/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := networks.Send(c)
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res.NotFoundError("Authorization header is required", nil)
			return
		}
		// fmt.Println("authHeader", authHeader)

		// parts := strings.Split(authHeader, " ")
		// if len(parts) != 2 || parts[0] != "Bearer" {
		// 	res.UnauthorizedError("Invalid token format", nil)
		// 	return
		// }

		tokenString := authHeader
		env, _ := config.LoadEnv()
		secret := env.JWTConfig.Secret

		claims, err := utils.VerifyToken(tokenString, secret)
		if err != nil {
			res.UnauthorizedError("Invalid or expired token", err)
			c.Abort()
			return
		}

		// Set claims in context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}
