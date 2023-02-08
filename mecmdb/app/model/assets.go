package model

import (
	. "mecmdb/app/database"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 10:10
 */

// idc company相关

type IdcCompany struct {
	Base
	CompanyName string `gorm:"column:company_name; type:varchar(255); unique"`
}

/**
创建IDC公司使用接口实例
*/

type IdcCompanyInstance struct {
	ID          uint   `json:"id"`
	CompanyName string `json:"company_name"`
}

/**
设置IDC公司表名
*/

func (company IdcCompany) TableName() string {
	return "idc_company"
}

/**
添加IDC公司，操作
*/

func (company *IdcCompany) Create() error {
	err := Orm.Create(&company).Error
	return err
}

func (company IdcCompany) GetAll() ([]IdcCompanyInstance, error) {
	var idcCompanyList []IdcCompanyInstance
	//普通查询，并写入实例
	//err := Orm.Table(company.TableName()).Scan(&idcCompanyInstanceList).Error

	//倒序排列
	err := Orm.Table(company.TableName()).Where("status = ?", 1).Order("id desc").Scan(&idcCompanyList).Error
	return idcCompanyList, err
}

func (company *IdcCompany) GetOneById(id uint) (IdcCompanyInstance, error) {
	var instance IdcCompanyInstance
	err := Orm.First(&company, id).Scan(&instance).Error
	return instance, err
}

func (company *IdcCompany) Delete(id uint) error {
	err := Orm.Table(company.TableName()).Where("id = ?", id).Update("status", 0).Error
	return err
}

func (company *IdcCompany) Update() error {
	err := Orm.Table(company.TableName()).Where("id = ?", company.ID).Updates(&company).Error
	return err
}

//func (company *IdcCompany) GetCompanyIdByName(name string) (IdcCompanyInstance, error) {
//	var instance IdcCompanyInstance
//	err := Orm.Table(company.TableName()).Where("company_name = ?", name).Scan(&instance).Error
//	return instance, err
//}

// idc 相关操作封装

type Idc struct {
	Base
	IdcName   string `json:"idc_name" gorm:"varchar(255)"`
	CompanyID uint   `json:"company_id" gorm:"type:int"`
	Region    string `json:"region" gorm:"varchar(255)"`
}

type IdcInstance struct {
	ID          uint   `json:"id"`
	IdcName     string `json:"idc_name"`
	Region      string `json:"region"`
	CompanyName string `json:"company_name"`
}

func (idc Idc) TableName() string {
	return "idc"
}

func (idc *Idc) Create() error {
	err := Orm.Create(&idc).Error
	return err
}

func (idc *Idc) GetAll() ([]IdcInstance, error) {
	var idcList []IdcInstance
	err := Orm.Table(idc.TableName()).Where("status = ?", 1).Order("id desc").Scan(&idcList).Error

	return idcList, err
}
