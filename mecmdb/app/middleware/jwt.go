package middleware

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/6 15:06
 */

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mecmdb/app/constants"
	"mecmdb/app/utils"
	"net/http"
)

// JWTAuthorization 验证token
func JWTAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.CodeAuthticateFail,
				"message": constants.TokenIsInvalid,
			})
			c.Abort()
			return
		}

		zap.S().Debugf("get token: %#v", token)

		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.GetPayloadByToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    constants.CodeAuthticateFail,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// 存储认证的载荷信息和token，保留至后面使用。
		c.Set("claims", claims)
		c.Set("access_token", token)
		c.Next()
	}
}
