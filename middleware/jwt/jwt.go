package jwt

import (
	"net/http"
	"vote_api/pkg/e"
	"vote_api/pkg/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// WxappJWT 小程序鉴权中间件
func WxappJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var claims *util.Claims
		var err error

		code = e.SUCCESS

		token := c.Request.Header.Get("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err = util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}

// AdminJWT 后台鉴权中间件
func AdminJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var claims *util.Claims
		var err error

		code = e.SUCCESS

		token, _ := c.Cookie("vue_admin_template_token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err = util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}
