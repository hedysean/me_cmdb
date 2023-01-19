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
