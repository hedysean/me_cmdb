package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
import . "mecmdb/app/model"

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 11:06
 */

/**
创建idc_company
*/

func CreateIdcCompany(ctx *gin.Context) (IdcCompanyInstance, error) {
	idcCompany := IdcCompany{}

	// 添加创建时间
	idcCompany.CreateTime = time.Now()

	//添加更新时间，第一次创建时间和更新时间一致
	idcCompany.UpdateTime = time.Now()
	var instance IdcCompanyInstance
	var err error
	fmt.Println("请求数据", ctx.Request.Body)
	if err = ctx.ShouldBindJSON(&idcCompany); err != nil {
		return instance, err
	}

	//数据写入到数据库
	err = idcCompany.Create()

	instance = IdcCompanyInstance{
		ID:          idcCompany.ID,
		CompanyName: idcCompany.CompanyName,
	}
	fmt.Println("instance", instance)
	return instance, err
}

/**
获取所有idc_company列表
*/

func GetIdcCompanyInstanceList(ctx *gin.Context) ([]IdcCompanyInstance, error) {
	idcCompany := IdcCompany{}
	idcCompanyInstanceList, err := idcCompany.GetAll()
	return idcCompanyInstanceList, err
}
