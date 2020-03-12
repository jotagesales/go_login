package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home endpoint
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "status ok"})
}
