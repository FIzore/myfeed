package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"myfeed/utils"
	"time"
)

var db *gorm.DB
var err error

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbAddress,
		utils.DbPort,
		utils.DbName)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		SkipDefaultTransaction: false, // 禁用默认事务
		NamingStrategy: schema.NamingStrategy{ // 命名策略
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			//TablePrefix: 表名前缀
		},
		DisableForeignKeyConstraintWhenMigrating: true, //不使用外键约束
	})
	_ = db.AutoMigrate(&Auditgoals{}, &Accounts{}) //todo 编写数据库
	if err != nil {
		log.Fatal("数据库初始化失败", err)
	}
	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
