package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "mecmdb/app/model"
	"time"
)

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

func GetIdcCompanyList(ctx *gin.Context) ([]IdcCompanyInstance, error) {
	idcCompany := IdcCompany{}
	idcCompanyList, err := idcCompany.GetAll()
	return idcCompanyList, err
}

func DeleteIdcCompany(id uint) error {
	idcCompany := IdcCompany{}
	err := idcCompany.Delete(id)
	return err
}
