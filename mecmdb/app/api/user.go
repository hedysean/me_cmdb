package api

import (
	"github.com/gin-gonic/gin"
	"mecmdb/app/constants"
	"net/http"

	. "mecmdb/app/services"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/1/18 22:28
 */

/**
用户认证登陆
*/

func UserAuthenticate(ctx *gin.Context) {
	var data map[string]string = map[string]string{
		"userame": "xiaoming",
		"age":     "16",
	}
	ctx.JSON(http.StatusOK, data)
}

/**
创建用户
*/

func UserCreate(ctx *gin.Context) {

	id, err := CreateUser(ctx)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    constants.CodeCreateUserFail,
			"message": constants.CreateUserFail,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.CreateUserSuccess,
		"data":    id,
	})
}
