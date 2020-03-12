package middewares

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jotagesales/pkg/models"
	log "github.com/sirupsen/logrus"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var identifykey = "email"

type User struct {
	Name  string
	Email string
}

// NewAuth create a new jwt auth middleware
func NewAuth() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "test login",
		Key:           []byte("my secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   identifykey,
		TokenLookup:   "header: Authorization, query: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identifykey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{Email: claims[identifykey].(string)}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var user models.User

			db, _ := c.Keys["DB"].(*gorm.DB)
			db.Where("email = ? AND password = ?", loginVals.Email, loginVals.Password).Find(&user)

			if user.Email == "" {
				return "", jwt.ErrFailedAuthentication
			}

			return &User{Name: user.Name, Email: user.Email}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	return authMiddleware
}
