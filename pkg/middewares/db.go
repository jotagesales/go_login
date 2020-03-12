package middewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ContextDB set DB key in request context
func ContextDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
