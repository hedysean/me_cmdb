package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"mecmdb/app/config"
	"time"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/1/18 9:32
 */

var Orm *gorm.DB

func InitDB(cfg *config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset)
	var err error
	Orm, err = gorm.Open(cfg.Driver, dsn)
	if err != nil {
		zap.S().Errorf("database connection fail: %v", err.Error())
		panic(err.Error())
	}
	// 最大链接数
	Orm.DB().SetMaxOpenConns(cfg.MaximumConn)
	// 最大空闲连接数
	Orm.DB().SetMaxIdleConns(cfg.MaximumFreeConn)
	// 每个链接
	//的最大生命周期
	Orm.DB().SetConnMaxLifetime(time.Duration(cfg.TimeOut))
	if Orm.Error != nil {
		zap.S().Errorf("database error: %v", Orm.Error)
		panic(Orm.Error)
	}

	return Orm
}

//定义初始化SQL必须字段
