package router

import (
	"github.com/gin-gonic/gin"
	"mecmdb/app/api"
	"mecmdb/app/utils"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/4 20:14
 */

func InitUserRouter(Router *gin.RouterGroup) {
	/**
	用户相关的路由组
	*/
	UserRouter := Router.Group("user")
	{
		// 用户认证登陆
		utils.Register(UserRouter, []string{"POST"}, "authenticate", api.UserAuthenticate)
		// 用户创建
		//utils.Register(UserRouter, []string{"POST"}, "", api.UserCreate)
		//utils.Register(UserRouter, []string{"POST"}, "", middleware.JWTAuthorization(), api.UserCreate)
	}
}
