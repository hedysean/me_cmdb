package api

import (
	"github.com/gin-gonic/gin"
	"mecmdb/app/constants"
	"mecmdb/app/services"
	"net/http"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 14:51
 */

/**
添加idc公司
*/

func CreateProviderApi(ctx *gin.Context) {
	idcCompany, err := services.CreateProviderService(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeCreateProviderFail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data":    idcCompany,
	})
}

func GetProviderListApi(ctx *gin.Context) {
	idcCompanyList, err := services.GetProviderListService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeGetProviderFail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"idc_company_list": idcCompanyList,
		},
	})
}
func DeleteProviderApi(ctx *gin.Context) {
	err := services.DeleteProviderService(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeDelProviderFail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
	})
}
func UpdateProviderApi(ctx *gin.Context) {
	newInstance, err := services.UpdateProviderService(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeUpdateProviderFail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data":    newInstance,
	})
}

/**
添加机房信息相关
*/

func CreateIdcApi(ctx *gin.Context) {

	err := services.CreateIdcService(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeCreateIdcFail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
	})
}
