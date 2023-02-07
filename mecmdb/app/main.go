package main

import (
	"fmt"
	"mecmdb/app/config"
	"mecmdb/app/initialize"
	"mecmdb/app/model"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	. "mecmdb/app/database"
)

func main() {

	//获取一个基于整个目录入口所在得路径
	dir, err := filepath.Abs(filepath.Dir("."))
	// vscdoe和 jet 目录逻辑不一致，这里需要根据编辑器修改
	//fmt.Println("dir", dir)
	if err != nil {
		panic(err.Error())
	}

	//初始化配置
	err = config.Init(fmt.Sprintf("%s/mecmdb/app/config/config.json", dir))
	fmt.Println("当前目录路径", dir)
	if err != nil {
		panic(err.Error())
	}

	//设置调试模式
	gin.SetMode(config.Conf.Mode)

	// 日志初始化
	if err := initialize.InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	zap.S().Debugf("调试信息:%d", config.Conf.Port)

	//创建路由
	Router := initialize.InitRouter()

	//打印启动信息
	zap.S().Infof("服务端启动...端口：%d", config.Conf.Port)

	// 6. 数据库初始化
	Orm := InitDB(config.Conf.DatabaseConfig)
	// 数据迁移
	Orm.AutoMigrate(&model.User{})

	//迁移idc公司数据库
	Orm.AutoMigrate(&model.IdcCompany{})

	// 禁用复数ll
	Orm.SingularTable(true)
	defer Orm.Close()

	//监听端口，默认在8080,
	err = Router.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		zap.S().Panic("服务端启动失败：", err.Error())
	}
}
