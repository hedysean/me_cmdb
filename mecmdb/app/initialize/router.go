package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mecmdb/app/middleware"
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

	zap.S().Info("日志中间件注册完成...")
	//Api路由分组
	//ApiGroup := Router.Group("/api")

	return Router

}
