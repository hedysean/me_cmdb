package config

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/goccy/go-json"
)

/*
* @author Yapeng
* @E-mail linuxsan@msn.com
* @date 2023/1/17 22:01
 */

//配置文件可以使用任何格式，例如：yaml,json,ini,本项目中使 json

//服务运行参数

type Config struct {
	Mode                string `json:"mode"`
	Port                int    `json:"port"`
	Logenv              string `json:"logenv"`
	SecretKey           string `json:"secret_key"`
	*LogConfig          `json:"log"`
	*DatabaseConfig     `json:"database"`
	*jwt.StandardClaims `json:"jwt"`
}

// 将日志输出到文件中
type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver          string `json:"driver"`
	Host            string `json:"host"`
	Port            string `json:"port"`
	Database        string `json:"database"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Charset         string `json:"charset"`
	MaximumConn     int    `json:"maximum_connection"`
	MaximumFreeConn int    `json:"maximum_free_connection"`
	TimeOut         int    `json:"timeout"`
}

var Conf = new(Config)

// Init 初始化配置，从指定位置加载配置文件

func Init(filepath string) error {
	/**
	filepath 配置文件路径
	*/
	b, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, Conf)
}
