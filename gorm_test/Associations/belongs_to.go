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

	/*db.Create(&User{
		Name: "lpx",
	}) // 可以看到插入失败了，外键约束，因为没有company_id

	db.Create(&User{
		Name: "lpx",
		Company: Company{
			Name: "家里蹲",
		},
	}) // 这样插入就会成功，因为会先插入company表，然后再插入user表

	db.Create(&User{
		Name:      "tom",
		CompanyID: 1,
	}) // 如果这里还是和上一条插入一样，就会导致多创建一个公司，所以这里选择对CompanyID赋值，而不是Company
	*/
	// 查询
	/*var user User
	db.First(&user, 2)
	fmt.Printf("user的Name为：%s\r\n", user.Name)
	fmt.Printf("user的Company.ID为：%d\r\n", user.Company.ID) //查询不到Company

	db.Preload("Company").First(&user, 2) //执行两个sql语句，一个查询company，一个查询user
	fmt.Printf("user的Company.ID为：%d\r\n", user.Company.ID)

	db.Joins("Company").First(&user, 2) //执行一个sql语句，使用join查询
	fmt.Printf("user的Company.ID为：%d\r\n", user.Company.ID)*/

}

/*type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}*/

type Company struct {
	ID   int
	Name string
}
