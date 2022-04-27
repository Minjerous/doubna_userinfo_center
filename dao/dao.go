package dao

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB
var DB *sql.DB
var DNS string
func init() {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DNS, // DSN data source name
		DefaultStringSize:         171,                                                                                                                                // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                                                                                                                              // 根据当前 MySQL 版本自动配置, &gorm.Config{})
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, //表名为复数
		},
		DisableForeignKeyConstraintWhenMigrating: true, //逻辑外键
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	Db = db
	//db.First(&model.UserInfo{})
	fmt.Println("数据库链接成功")
}

func init() {
	db, err := sql.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
	fmt.Println("数据库链接成功")
}
