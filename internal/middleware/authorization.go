package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/iki-rumondor/project3-grup9/internal/adapter/http/response"
	"github.com/iki-rumondor/project3-grup9/internal/utils"
)


func IsValidJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerToken = c.Request.Header.Get("Authorization")
		var bearer = strings.HasPrefix(headerToken, "Bearer")

		if !bearer {
			c.AbortWithStatusJSON(http.StatusBadRequest, response.Message{
				Message: "Bearer token is not valid",
			})
			return
		}

		jwt := strings.Split(headerToken, " ")[1]
		
		mapClaims, err := utils.VerifyToken(jwt)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Message{
				Message: err.Error(),
			})
			return
		}

		c.Set("map_claims", mapClaims)
		c.Next()

	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)

		role := mapClaims["role"].(string)
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Message{
				Message: "your access to this resource is denied",
			})
			return
		}
		c.Next()
	}
}

func SetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet("map_claims")
		mapClaims := mc.(jwt.MapClaims)  

		userID := uint(mapClaims["id"].(float64))

		c.Set("user_id", userID)
		c.Next()

	}
}
