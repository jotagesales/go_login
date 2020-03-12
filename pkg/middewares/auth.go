package middewares

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
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
			// var user models.User

			// db, exists := c.Get("DB")
			// fmt.Println(db)
			// fmt.Println(exists)

			// fmt.Println(loginVals.Email)

			// db.Where("email = ?", loginVals.Email).Find(&user)
			// TODO: added here response if email or password is not valid
			return &User{Name: "test", Email: "test@login.com"}, nil
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
