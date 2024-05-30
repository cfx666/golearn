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

	//_ = db.AutoMigrate(&User{})

	/*user := User{Name: "lpx"}
	fmt.Println("userId:", user.ID)
	result := db.Create(&user)
	fmt.Println("userId:", user.ID)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("影响的记录数：", result.RowsAffected)*/

	/*users := []User{
		{Name: "lpx1"},
		{Name: "lpx2"},
		{Name: "lpx3"},
	}
	result := db.Create(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("影响的记录数：", result.RowsAffected)
	for _, user := range users {
		fmt.Println("userId:", user.ID)
	}*/

	/*users := []User{
		{Name: "lpx4"},
		{Name: "lpx5"},
		{Name: "lpx6"},
	}
	result := db.CreateInBatches(&users, 100)
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Println("影响的记录数：", result.RowsAffected)
	for _, user := range users {
		fmt.Println("userId:", user.ID)
	}*/
	/*
		db.Model(&User{}).Create(map[string]interface{}{
			"Name": "jinzhu", "Age": 18,
		})
	*/
	/*db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu1", "Age": 18},
		{"Name": "jinzhu2", "Age": 18},
		{"Name": "jinzhu3", "Age": 18},
	})*/

	//var user User
	//db.First(&user) // SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	//fmt.Printf("user:%v", user)

	//db.Last(&user) //SELECT * FROM `users` ORDER BY `users`.`id` DESC LIMIT 1
	//fmt.Printf("user:%v", user)

	//db.Take(&user) //SELECT * FROM `users` LIMIT 1
	//fmt.Printf("user:%v", user)

	/*db.First(&user, "2")
	fmt.Printf("user:%v", user)*/

	/*var users []User
	db.Find(&users, []int{1, 2, 3})
	for _, user := range users {
		fmt.Printf("user:%v \r\n", user)
	}*/

	/*var user = User{ID: 10}
	db.First(&user)
	fmt.Printf("user:%v", user)*/

	/*var result User
	db.Model(User{ID: 10}).First(&result)
	fmt.Printf("user:%v", result)*/

	/*var users []User
	db.Find(&users)  //SELECT * FROM `users`
	for _, user := range users {
		fmt.Printf("user:%v \r\n", user)
	}
	*/

	//var user User
	/*db.Where("name = ?", "jinzhu").First(&user) //SELECT * FROM `users` WHERE name = 'jinzhu' ORDER BY `users`. `id` LIMIT 1
	fmt.Printf("user:%v", user)*/

	// Struct
	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user) // SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu' AND `users`.`age` = 20 ORDER BY `users`.`id` LIMIT 1

	// Map
	/*var users []User
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users) //SELECT * FROM `users` WHERE `age` = 20 AND `name` = 'jinzhu'
	*/
	/*var users []User
	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users) // SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu'
	*/
	/*var users []User
	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users) //SELECT * FROM `users` WHERE `Age` = 0 AND `Name` = 'jinzhu'
	*/

	/*var user User
	db.First(&user)

	user.Name = ""
	user.Age = 100
	db.Save(&user)*/

	/*db.Save(&User{Name: "lpx", Age: 18})
	db.Save(&User{ID: 8, Name: "lpx", Age: 18})*/

	//db.Model(&User{ID: 6}).Update("name", "lpx")
	//db.Model(&User{}).Where("id = ?", 6).Update("name", "ddf")

	//db.Model(&User{ID: 10}).Where("name = ?", "ddf").Update("name", "xdn")
	/*var r string
	db.Model(&User{ID: 10}).Updates(User{Name: &r, Age: 18})*/
	//db.Model(&User{ID: 10}).Updates(map[string]interface{}{"name": ""})

	db.Model(&User{ID: 10}).Select("name").Updates(User{Name: "", Age: 18})
	//db.Model(&User{ID: 6}).Update("name", "")

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
