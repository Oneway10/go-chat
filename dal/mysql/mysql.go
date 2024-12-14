package mysql

import (
	"chat/config"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Init() {
	if db != nil {
		return
	}
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Mysql.Username,
			config.Mysql.Password,
			config.Mysql.ServerHost,
			config.Mysql.ServerPort,
			config.Mysql.Database)
		DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true, // 自动创建表的时候不创建外键
			NamingStrategy: schema.NamingStrategy{ // 自动创建表时候表名的配置
				SingularTable: true,
			},
		})
		if err != nil {
			panic(err)
		}
		db = DB
		hlog.Info("init mysql success")
	})
}

func DB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
