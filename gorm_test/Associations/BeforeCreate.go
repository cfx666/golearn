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

	//_ = db.AutoMigrate(&student{})

	db.Create(&student{
		Name: "xdn",
	})
}

// gorm的BeforeCreate
type student struct {
	gorm.Model
	Name    string
	AddTime sql.NullTime
}

func (s *student) BeforeCreate(tx *gorm.DB) (err error) {
	s.AddTime = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	return
}
