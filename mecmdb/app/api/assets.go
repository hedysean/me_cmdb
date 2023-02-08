package api

import (
	"github.com/gin-gonic/gin"
	"mecmdb/app/constants"
	"mecmdb/app/services"
	"net/http"
	"strconv"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 14:51
 */

/**
添加idc公司
*/

func CreateIdcCompany(ctx *gin.Context) {
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

func GetIdcCompanyList(ctx *gin.Context) {
	idcCompanyList, err := services.GetIdcCompanyList()
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
			"idc_company_list": idcCompanyList,
		},
	})
}
func DeleteIdcCompany(ctx *gin.Context) {
	strId, _ := ctx.GetQuery("id")
	delId, _ := strconv.Atoi(strId)

	err := services.DeleteIdcCompany(uint(delId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeDelIdcCompanyFail,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    constants.CodeSuccess,
		"message": constants.Success,
	})
}
func UpdateIdcCompany(ctx *gin.Context) {
	newInstance, err := services.UpdateIdcCompany(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    constants.CodeUpdateIdcCompanyFail,
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
