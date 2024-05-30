package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&User{})

	// 插入
	/*user := User{
		Languages: []Language{{Name: "go"}, {Name: "java"}},
	}
	db.Create(&user)*/

	/*var user1 User
	db.Preload("Languages").First(&user1)
	for _, language := range user1.Languages {
		fmt.Printf("Language.Name: %s\r\n", language.Name)
	}
	*/
	/*var user2 User
	db.First(&user2)

	var languages []Language
	db.Model(&user2).Association("Languages").Find(&languages) // 执行的时join语句

	for _, Language := range languages {
		fmt.Printf("Language.Name: %s\r\n", Language.Name)
	}*/
}

type User struct {
	gorm.Model
}

func (User) TableName() string {
	return "zidingyi_user"
}

type Language struct {
	gorm.Model
	Name string
}
