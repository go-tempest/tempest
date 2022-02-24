package gorm

import (
	"fmt"
	"github.com/go-tempest/tempest/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var DB *gorm.DB

// Initialize 初始化Mysql数据库
func Initialize() {
	m := config.TempestConfig.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	c := logConfig(&m)
	var err error
	if DB, err = gorm.Open(mysql.New(mysqlConfig), c); err != nil {
		fmt.Println("MySQL启动异常", zap.Any("err", err)) //todo 替换成logger
		os.Exit(0)
	} else {
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleCons)
		sqlDB.SetMaxOpenConns(m.MaxOpenCons)
	}
}

// config 根据配置决定是否开启日志
func logConfig(m *config.Mysql) (c *gorm.Config) {
	if m.LogMode {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	} else {
		c = &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}
	return
}
