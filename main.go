package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"

	//"myfeed/model"
	"myfeed/routes"
	"time"
)

type User struct {
	UserID   uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username string `form:"username" binding:"required"` // 这里改为 form
	Password string `form:"password" binding:"required"` // 这里改为 form
	Email    string `form:"email" binding:"required"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/myfeed?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
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
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/login", routes.Login)
	r.POST("/login/auth", func(c *gin.Context) { //登录
		var user User
		c.ShouldBind(&user)
		res := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
		if res.RowsAffected == 0 {
			c.JSON(200, gin.H{
				"message": "登录失败",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "登录成功",
			})
		}
	})
	r.GET("/register", routes.Register)
	r.POST("/register/auth", func(c *gin.Context) { //注册TODO加上邮箱验证码
		var user User
		c.ShouldBind(&user)
		res := db.Where("username = ? AND password = ? AND email = ?", user.Username, user.Password, user.Email).First(&user)
		if res.RowsAffected != 0 {
			c.JSON(200, gin.H{
				"message": "注册失败",
			})
		} else {
			db.Create(&user)
			c.JSON(200, gin.H{
				"message": "注册成功",
			})
		}
	})

	r.GET("/password/reset", routes.Reset)
	r.GET("/", routes.Index)
	r.GET("/report", routes.Report)
	r.GET("/manual", routes.Manual)
	r.GET("/privacy", routes.Privacy)

	r.Run(":8080")
}
