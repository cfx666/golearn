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

// gorm的BeforeCreate

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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "t_", // 表名前缀，`User` 的表名应该是 `t_users`
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	_ = db.AutoMigrate(&student{})

	db.Create(&student{
		Name: "lpx",
	})

}

type student struct {
	gorm.Model
	Name    string
	AddTime time.Time
	//AddTime sql.NullTime
}

func (s *student) BeforeCreate(tx *gorm.DB) (err error) {
	s.AddTime = time.Now()
	return
}
