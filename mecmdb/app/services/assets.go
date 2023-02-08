package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "mecmdb/app/model"
	"reflect"
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

func GetIdcCompanyList() ([]IdcCompanyInstance, error) {
	idcCompany := IdcCompany{}
	idcCompanyList, err := idcCompany.GetAll()
	return idcCompanyList, err
}

func DeleteIdcCompany(id uint) error {
	idcCompany := IdcCompany{}
	err := idcCompany.Delete(id)
	return err
}

func UpdateIdcCompany(ctx *gin.Context) (IdcCompanyInstance, error) {
	idcCompany := IdcCompany{}

	idcCompany.UpdateTime = time.Now()
	fmt.Println(idcCompany.ID, reflect.TypeOf(idcCompany.ID))

	var instance IdcCompanyInstance
	var err error

	//id 直接在这里写入，后端直接获取
	if err = ctx.ShouldBindJSON(&idcCompany); err != nil {
		return instance, err
	}

	//写入后查询数据是否正确，并返回给前端
	err = idcCompany.Update()
	fmt.Println(idcCompany.ID, reflect.TypeOf(idcCompany.ID))
	instance, err = idcCompany.GetOneById(idcCompany.ID)

	return instance, err
}
