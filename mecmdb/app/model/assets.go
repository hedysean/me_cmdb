package model

import (
	. "mecmdb/app/database"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/2/7 10:10
 */

type IdcCompany struct {
	Base
	CompanyName string `json:"company_name" gorm:"type:varchar(255); unique_index"`
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
添加IDC名称
*/

func (company *IdcCompany) Create() error {
	err := Orm.Create(&company).Error
	return err
}

func (company IdcCompany) GetAll() ([]IdcCompanyInstance, error) {
	var idcCompanyInstanceList []IdcCompanyInstance
	//普通查询，并写入实例
	//err := Orm.Table(company.TableName()).Scan(&idcCompanyInstanceList).Error

	//倒序排列
	err := Orm.Table(company.TableName()).Where("status = ?", 1).Order("id desc").Scan(&idcCompanyInstanceList).Error
	return idcCompanyInstanceList, err
}

func (company *IdcCompany) GetOneById(id uint) error {
	err := Orm.First(&company, id).Error
	return err
}

//func (company *IdcCompany) Update(id uint) error {

func (company *IdcCompany) Delete(id uint) error {
	err := Orm.Table(company.TableName()).Where("id = ?", id).Update("status", 0).Error
	return err
}
