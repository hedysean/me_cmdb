package api

import (
	"github.com/gin-gonic/gin"
	"mecmdb/app/constants"
	"mecmdb/app/services"
	"mecmdb/app/utils"
	"net/http"
	"strconv"
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
	// 用户登陆认证业务
	//user, err := UserLogin(ctx)
	user := services.UserLogin(ctx)

	//if err != nil  {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code":    constants.CodeNoSuchUser,
	//		"message": err.Error(),
	//	})
	//	return
	//}

	// 生成token
	newJwt := utils.NewJWT()
	publicClaims := utils.PublicClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}

	token, err := newJwt.AccessToken(publicClaims)
	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

/**
用户认证登录
*/

/**
创建用户
*/

func UserCreate(ctx *gin.Context) {

	id, err := services.CreateUser(ctx)
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
