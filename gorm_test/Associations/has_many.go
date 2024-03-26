package main

import (
	"gorm.io/gorm"
)

func main() {
	/*newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)*/

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	/*dsn := "lpx:lpxlpx@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}*/

	//_ = db.AutoMigrate(&User{})
	//_ = db.AutoMigrate(&CreditCard{})

	// 插入
	/*user := User{}
	db.Create(&user)*/

	/*db.Create(&CreditCard{
		Number: "12",
		UserID: user.ID,
	})*/

	/*db.Create(&CreditCard{
		Number: "34",
		UserID: user.ID,
	})*/

	// 查询
	/*var user1 User
	db.Preload("CreditCards").First(&user1, 1)
	for _, card := range user1.CreditCards {
		fmt.Printf("CreditCard.Number: %s\r\n", card.Number)
	}*/
}

// User 有多张 CreditCard
/*type User struct {
	gorm.Model
	CreditCards []CreditCard
}*/

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
