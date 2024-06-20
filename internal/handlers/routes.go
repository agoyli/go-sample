package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func LoadApiV1Routes(routes *gin.Engine) {
	v1 := routes.Group("/api/v1")
	{
		// users
		v1.GET("users", UserListHandler)
		v1.POST("users", UserCreateHandler)
		v1.PUT("users/{id}", UserUpdateHandler)
		v1.DELETE("users", UserDeleteHandler)
	}
}

func Routes(routes *gin.Engine) {
	isProd := true

	// CORS
	routes.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "RefreshToken", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		AllowWebSockets:  true,
		// MaxAge:           12 * time.Hour,
	}))

	LoadApiV1Routes(routes)

	routes.Static("/uploads", "./uploads")
	if !isProd {
		routes.StaticFile("/docs", "./docs/docs.html")
		routes.StaticFile("/openapi.yaml", "./docs/openapi.yaml")
		routes.StaticFile("/openapi.json", "./docs/openapi.json")
	}
}
