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

func IdcCompanyCreate(ctx *gin.Context) {
	idcCompany, err := services.CreateIdcCompany(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeCreateIdcCompanyFail,
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

func IdcCompanyInstanceGet(ctx *gin.Context) {
	idcCompanyInstanceList, err := services.GetIdcCompanyInstanceList(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeGetIdcCompanyFail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
		"data": map[string]interface{}{
			"idc_company_instance_list": idcCompanyInstanceList,
		},
	})
}
