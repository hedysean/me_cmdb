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

func CreateProviderService(ctx *gin.Context) (ProviderInstance, error) {
	p := Provider{}

	// 添加创建时间
	p.CreateTime = time.Now()

	//添加更新时间，第一次创建时间和更新时间一致
	p.UpdateTime = time.Now()
	var instance ProviderInstance
	var err error
	if err = ctx.ShouldBindJSON(&p); err != nil {
		return instance, err
	}

	//数据写入到数据库
	err = p.Create()

	instance = ProviderInstance{
		ID:           p.ID,
		ProviderName: p.ProviderName,
	}
	fmt.Println("instance", instance)
	return instance, err
}

/**
获取所有idc_company列表
*/

func GetProviderListService() ([]ProviderInstance, error) {
	p := Provider{}
	pList, err := p.GetAll()
	return pList, err
}

func DeleteProviderService(ctx *gin.Context) error {
	p := Provider{}
	//id 直接在这里写入，后端直接获取
	var err error
	if err = ctx.ShouldBindJSON(&p); err != nil {
		return err
	}
	fmt.Println("打印传输过来得ID", p.ID)
	err = p.Delete(p.ID)
	return err
}

func UpdateProviderService(ctx *gin.Context) (ProviderInstance, error) {
	p := Provider{}

	p.UpdateTime = time.Now()
	fmt.Println(p.ID, reflect.TypeOf(p.ID))

	var instance ProviderInstance
	var err error

	//id 直接在这里写入，后端直接获取
	if err = ctx.ShouldBindJSON(&p); err != nil {
		return instance, err
	}

	//写入后查询数据是否正确，并返回给前端
	err = p.Update()
	fmt.Println(p.ID, reflect.TypeOf(p.ID))
	instance, err = p.GetOneById(p.ID)

	return instance, err
}

func CreateIdcService(ctx *gin.Context) error {
	idc := Idc{}

	//更新时间
	idc.CreateTime = time.Now()
	idc.UpdateTime = time.Now()

	var err error

	if err = ctx.ShouldBindJSON(&idc); err != nil {
		return err
	}
	fmt.Println("绑定后数据机房名称：", idc.IdcName)
	fmt.Println("绑定后数据公司ID：", idc.ProviderID)
	fmt.Println("绑定后数据机房位置：", idc.Region)
	err = idc.CreateIdc()

	return err
}
