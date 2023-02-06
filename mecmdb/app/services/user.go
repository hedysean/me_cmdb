package services

import (
	"github.com/gin-gonic/gin"
	. "mecmdb/app/model"
	. "mecmdb/app/utils"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/1/18 22:19
 */
/**
创建用户
*/

func CreateUser(ctx *gin.Context) (uint, error) {
	user := User{}
	var err error
	if err = ctx.ShouldBindJSON(&user); err != nil {
		return 0, err
	}

	user.HashPassword, err = MakeHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	return user.Insert()
}

/*
*
获取指定ID账户
*/
func GetUserById(id uint) (user User) {
	user = User{}
	user.GetOneById(id)
	return user
}

/**
根据账户信息(用户名、手机、邮箱)获取用户
*/

func GetUserByAccount(account string) (user User) {
	user = User{}
	user.GetOneByAccount(account)
	return user
}

/**
获取所有用户
*/

func GetAllUser() []User {
	user := User{}
	return user.GetAll()
}

/**
更新密码
*/

func ChangeUserPassword(user User, RawPassword string) {
	password, _ := MakeHashPassword(RawPassword)
	user.ChangePassword(password)
}

// 测试用

type TestUserinfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func UserLogin(ctx *gin.Context) (user TestUserinfo) {
	user = TestUserinfo{
		ID:       1,
		Username: "san",
		Password: "123456",
		Nickname: "abcd",
		Avatar:   "ddd",
	}

	return user
}

//func UserLogin(ctx *gin.Context) (user User, err error) {
//	user = User{}
//	//if err = ctx.ShouldBindJSON(&user); err != nil {
//	//	return user, err
//	//}
//
//	//if err = validator.UserValidator(&user); err != nil {
//	//	return user, err
//	//}
//
//	//user.GetOneByAccount(user.Username)
//	//if user.ID < 1 {
//	//	return user, errors.New(constants.NoSuchUser)
//	//}
//	//
//	//ret := CheckPassword(user.HashPassword, user.Password)
//	//
//	//if !ret {
//	//	err = errors.New(constants.PasswordError)
//	//	return
//	//}
//
//	return
//}
