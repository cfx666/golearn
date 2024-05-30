package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log    select * from `users`  where `users`.`id` = 1
			Colorful:                  false,       // Disable color
		},
	)

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "lpx:lpxlpx@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	//db.AutoMigrate(&Product{})

	/*db.Create(&[]Product{
		{Code: "c++", Price: 1000},
		{Code: "java", Price: 2000},
		{Code: "php", Price: 3000},
		{Code: "python", Price: 5000},
	})*/

	db.Delete(&Product{}, 1) // update

	db.First(&Product{}, 1)

	//db.Unscoped().Delete(&Product{}, 1) //DELETE FROM `products` WHERE `products`.`id` = 1

	/*var user = User{ID: 8} //DELETE FROM `users` WHERE `users`.`id` = 6
	//db.Delete(&user)

	db.Where("name = ?", "jinzhu").Delete(&user) // DELETE FROM `users` WHERE name = 'jinzhu' AND `users`.`id` = 8*/
}

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // 一个常规字符串字段
	Email        *string        // 一个指向字符串的指针, allowing for null values
	Age          uint8          // 一个未签名的8位整数
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // 创建时间（由GORM自动管理）
	UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
}

/*type Product struct {
	gorm.Model
	Code  string
	Price uint
}*/

type Product struct {
	Id        uint
	DeletedAt gorm.DeletedAt
	Code      string
	Price     uint
}
