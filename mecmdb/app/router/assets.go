package router

import (
	"github.com/gin-gonic/gin"
	"mecmdb/app/api"
	"mecmdb/app/utils"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 11:05
 */

func InitAssetsRouter(Router *gin.RouterGroup) {
	HostRouter := Router.Group("assets")
	//HostRouter.Use(middleware.JWTAuthorization())
	utils.Register(HostRouter, []string{"POST"}, "company", api.CreateProviderApi)
	utils.Register(HostRouter, []string{"GET"}, "company", api.GetProviderListApi)
	utils.Register(HostRouter, []string{"DELETE"}, "company", api.DeleteProviderApi)
	utils.Register(HostRouter, []string{"PUT"}, "company", api.UpdateProviderApi)

	//IDC相关操作
	utils.Register(HostRouter, []string{"POST"}, "idc", api.CreateIdcApi)
}
