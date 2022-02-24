package db

import "gorm.io/gorm"

type Interface interface {
	Initialize()                          //初始化
	DB() *gorm.DB                         //获取DB
	RegisterTables(models ...interface{}) //注册数据库表
}
