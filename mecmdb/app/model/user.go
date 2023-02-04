package model

import (
	"github.com/jinzhu/gorm"
	. "mecmdb/app/database"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/1/18 21:56
 */

/**
用户实体
*/

type User struct {
	gorm.Model
	Username     string `validate:"required" label:"账户名" gorm:"unique_index;size:255;comment:'账户名'" json:"username" sql:"index" json:"username,omitempty"`
	HashPassword string `gorm:"not null;size:255;comment:'密码'" json:"-" json:"hashPassword,omitempty"`
	Password     string `validate:"required" label:"登陆密码" gorm:"-" json:"password,omitempty" json:"password,omitempty"`
	Nickname     string `gorm:"size:255;comment:'昵称'" json:"nickname" sql:"index" json:"nickname,omitempty"`
	Mobile       string `gorm:"index;size:15;comment:'手机';" json:"mobile" json:"mobile,omitempty"`
	Email        string `validate:"omitempty,email" valid:"email" gorm:"index;size:255;comment:'邮箱';" json:"email" json:"email,omitempty"`
	Avatar       string `gorm:"size:255;comment:'头像'" json:"avatar" json:"avatar,omitempty"`
	Sex          bool   `gorm:"type:boolean;default:true;comment:'性别'" json:"sex" json:"sex,omitempty"`
	Ip           string `validate:"omitempty,ipv4" gorm:"size:255;comment:'IP地址';" json:"ip" json:"ip,omitempty"`
	Status       bool   `gorm:"type:boolean;default:true;comment:'状态'" json:"status" json:"status,omitempty"`
}

/**
设置表名
*/

func (User) TableName() string {
	return "user_info"
}

/**
创建用户
*/

func (user User) Insert() (id uint, err error) {
	//添加数据
	result := Orm.Create(&user)
	return user.ID, result.Error
}

/**
根据指定ID获取用户
*/

func (user *User) GetOneById(id uint) {
	Orm.First(&user, id)
}

/**
根据账户信息(用户名、手机、邮箱)获取用户
*/

func (user *User) GetOneByAccount(account string) {
	Orm.First(&user, "username = ? or mobile = ? or email= ?", account, account, account)
}

/**
获取所有用户
*/

func (user User) GetAll() []User {
	var users []User
	Orm.Find(&users)
	return users
}

/**
更新密码
*/

func (user User) ChangePassword(HashPasswrd string) {
	user.HashPassword = HashPasswrd
	Orm.Save(&user)
}
