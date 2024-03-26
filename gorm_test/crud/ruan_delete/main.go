package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	// 日志设置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// 数据库连接
	dsn := "lpx:lpxlpx@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 根据结构体创建表
	_ = db.AutoMigrate(&Product{})

	// 添加测试数据
	db.Create([]Product{
		{Code: "python", Price: 1000},
		{Code: "java", Price: 2000},
		{Code: "c++", Price: 3000},
		{Code: "c", Price: 4000},
	})

	// 删除，软删除
	db.Delete(&Product{}, 1) // 逻辑删除，可以看到是update语句

	var product Product
	db.Where("id = 1").Find(&product) // 查询不到数据
	db.Where("id = 2").Find(&product) // 查询到数据

	// 永久删除，硬删除
	db.Unscoped().Delete(&Product{}, 4) // 看到是delete语句

}

type Product struct {
	ID      uint
	Code    string
	Price   uint
	Deleted gorm.DeletedAt
}
