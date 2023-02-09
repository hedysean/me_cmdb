package model

import (
	. "mecmdb/app/database"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 10:10
 */

// idc provider 相关
// provider 供应者，服务者

type Provider struct {
	Base
	ProviderName string `json:"Provider_name" gorm:"type:varchar(255); unique"`
}

/**
创建IDC供应商使用接口实例

*/

type ProviderInstance struct {
	ID           uint   `json:"id"`
	ProviderName string `json:"Provider_name"`
}

/**
设置IDC公司表名
*/

func (pro Provider) TableName() string {
	return "provider"
}

/**
添加IDC公司，操作
*/

func (pro *Provider) Create() error {
	err := Orm.Create(&pro).Error
	return err
}

func (pro *Provider) GetAll() ([]ProviderInstance, error) {
	var ProviderList []ProviderInstance
	//普通查询，并写入实例
	//err := Orm.Table(Provider.TableName()).Scan(&idcProviderInstanceList).Error

	//倒序排列
	err := Orm.Table(pro.TableName()).Where("status = ?", 1).Order("id desc").Scan(&ProviderList).Error
	return ProviderList, err
}

func (pro *Provider) GetOneById(id uint) (ProviderInstance, error) {
	var instance ProviderInstance
	err := Orm.First(&pro, id).Scan(&instance).Error
	return instance, err
}

func (pro *Provider) Delete(id uint) error {
	err := Orm.Table(pro.TableName()).Where("id = ?", id).Update("status", 0).Error
	return err
}

func (pro *Provider) Update() error {
	err := Orm.Table(pro.TableName()).Where("id = ?", pro.ID).Updates(&pro).Error
	return err
}

/**
idc 相关操作封装
*/

/**
配置数据库字段
*/

type Idc struct {
	Base
	IdcName    string `json:"idc_name" gorm:"varchar(255)"`
	ProviderID uint   `json:"provider_id" gorm:"type:int"`
	Region     string `json:"region" gorm:"varchar(255)"`
}

/**
声明查询返回数据字段
*/

type IdcInstance struct {
	ID           uint   `json:"id"`
	IdcName      string `json:"idc_name"`
	Region       string `json:"region"`
	ProviderName string `json:"Provider_name"`
}

func (idc Idc) TableName() string {
	return "idc"
}

func (idc *Idc) CreateIdc() error {
	err := Orm.Create(&idc).Error
	return err
}

func (idc *Idc) GetIdcAll() ([]IdcInstance, error) {
	var idcList []IdcInstance
	err := Orm.Table(idc.TableName()).Where("status = ?", 1).Order("id desc").Scan(&idcList).Error

	return idcList, err
}
