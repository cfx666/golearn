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

	_ = db.AutoMigrate(&User{}, &CreditCard{})

	// 插入。单独插入
	/*user := User{}
	db.Create(&user)

	db.Create(&CreditCard{
		Number: "12",
		UserID: user.ID,
	})

	db.Create(&CreditCard{
		Number: "34",
		UserID: user.ID,
	})*/

	// 查询
	/*var user User = User{Model: gorm.Model{ID: 2}}
	db.Preload("CreditCards").First(&user)
	for _, creditCard := range user.CreditCards {
		fmt.Println("卡号:", creditCard.Number)
	}*/

}

/*type User struct {
	gorm.Model
	MemberNumber string       `gorm:"index:idx_member_number"`
	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}*/

type CreditCard struct {
	gorm.Model
	Number     string
	UserNumber uint
}
