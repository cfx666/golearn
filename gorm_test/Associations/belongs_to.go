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

	_ = db.AutoMigrate(&User{})

	// 插入
	/*user := User{Name: "lpx"}
	db.Create(&user)*/ // 插入失败。CompanyID默认是0，这个字段有外键约束，所以插入失败

	/*user := User{
		Name: "lpx",
		Company: Company{
			Name: "家里蹲公司",
		},
	}
	db.Create(&user)*/ // 插入成功，执行的是两条sql，先插入Company，再插入User

	/*user := User{
		Name:      "ddf",
		CompanyID: 1,
	}
	db.Create(&user)*/ // 插入成功，只执行一条sql。执行的是插入User的sql

	/*var user1 User
	db.First(&user1)  //执行的是一条sql，查询user表
	fmt.Printf("公司名字：%s\r\n", user1.Company.Name)*/

	/*var user1 User
	db.Preload("Company").First(&user1) //执行的是两条sql，先查询Company，再查询User
	fmt.Printf("公司名字：%s\r\n", user1.Company.Name)*/

	/*var user1 User
	db.Joins("Company").First(&user1) //执行的是一条sql，使用join查询，left join
	fmt.Printf("公司名字：%s\r\n", user1.Company.Name)*/

	/*type Result struct {
		Name string
		Id   int
	}
	var result Result
	db.Model(&User{Model: gorm.Model{ // 自己写join查询来实现查询出company
		ID: 2,
	}}).Select("companies.id, companies.name").Joins("left join companies on companies.id = users.company_id").Scan(&result)
	fmt.Println(result.Id, result.Name)*/
}

/*type User struct {
	gorm.Model
	Name        string
	CompanyName int
	Company     Company `gorm:"references:Name"`
}*/

type Company struct {
	ID   int
	Name string `gorm:"index:idx_name"`
}
