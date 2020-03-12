package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jotagesales/pkg/handlers"
	"github.com/jotagesales/pkg/middewares"
)

// GetRoutes return gin router with all api routes
func GetRoutes(router *gin.Engine, db *gorm.DB) *gin.Engine {
	// routers API V1
	authMiddleware := middewares.NewAuth()

	router.Use(middewares.ContextDB(db))
	v1 := router.Group("/api/v1")
	v1.POST("/login", authMiddleware.LoginHandler)

	router.Use(authMiddleware.MiddlewareFunc())
	v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	v1.GET("/", handlers.Home)
	return router
}
