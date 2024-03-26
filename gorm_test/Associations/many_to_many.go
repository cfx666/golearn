package main

import (
	"fmt"
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
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "lpx:lpxlpx@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&User{})

	// 插入
	user := User{
		Languages: []Language{{Name: "go"}, {Name: "java"}},
	}
	db.Create(&user)

	// 查询
	var user1 User
	db.Preload("Languages").First(&user1)
	for _, language := range user1.Languages {
		fmt.Printf("Language.Name: %s\r\n", language.Name)
	}

	// 有些时候并不一定需要取出languages，可以使用Association来延迟加载
	// 如果我已经取出一个用户来了，但是这个用户之前没有使用perload来加载对应的Languages，那么我可以使用以下方式来加载、
	var user2 User
	db.First(&user2)
	var languages []Language
	db.Model(&user2).Association("Languages").Find(&languages) // 执行的时join语句
	for _, Language := range languages {
		fmt.Printf("Language.Name: %s\r\n", Language.Name)
	}

}

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}
