package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mecmdb/app/middleware"
	"mecmdb/app/router"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/1/17 22:36
 */

func InitRouter() *gin.Engine {
	//创建路由
	Router := gin.Default()
	//注册使用中间件
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	Router.Use(middleware.ExceptionMiddleware)

	//全局验证
	//Router.Use(middleware.JWTAuthorization())
	Router.Use(middleware.CORS)

	zap.S().Info("中间件注册完成...")
	//Api路由分组
	ApiGroup := Router.Group("/api")
	// 3. 初始化用户相关路由
	router.InitUserRouter(ApiGroup)
	router.InitAssetsRouter(ApiGroup)
	zap.S().Info("路由初始化完成....")

	return Router

}
