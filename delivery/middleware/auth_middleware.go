package middleware

import (
	"field_work/shared/service"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	RequireToken(roles ...string) gin.HandlerFunc
}

type authMiddleware struct {
	jwtService service.JwtService
}
type AuthHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}


// RequireToken implements AuthMiddleware.
func (a *authMiddleware) RequireToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context){
		var authHeader AuthHeader
		if err := ctx.ShouldBindHeader(&authHeader); err != nil{
			log.Printf("RequireToken.authHeader: %v \n", err.Error())
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenHeader := strings.Replace(authHeader.AuthorizationHeader, "Bearer ", "", -1)
		if tokenHeader == ""{
			log.Printf("RequireToken.tokenHeader \n")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := a.jwtService.ParseToken(tokenHeader)
		if err != nil {
			log.Printf("RequireToken.ParseToken: %v \n", err.Error())
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Set("user", claims["username"])

		validRole := false
		for _, role := range roles{
			if role == claims["role"]{
				validRole = true
				break
			}
		}

		if !validRole {
			log.Printf("RequireToken.validRole \n")
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}
		ctx.Next()
	}
}


func NewAuthMiddleware(jwtService service.JwtService) AuthMiddleware {
	return &authMiddleware{jwtService: jwtService}
}
