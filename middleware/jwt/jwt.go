package jwt

import (
	"github.com/gin-gonic/gin"
	"time"
	"word/pkg/app"
	"word/pkg/e"
	"word/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.GetHeader("token")

		if token == "" {
			code = e.ERROR_AUTH
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

			if code == e.SUCCESS {
				c.Set("userId", claims.UserId)
				c.Set("username", claims.Username)
			}
		}

		if code != e.SUCCESS {
			appG := app.Gin{c}
			appG.Response(401, code, e.GetMsgLabel(code), nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
