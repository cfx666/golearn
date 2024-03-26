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

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         `gorm:"column:myname"` // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
}

func main() {
	//日志配置
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

	// 数据库连接
	dsn := "lpx:lpxlpx@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 根据结构体创建表
	//_ = db.AutoMigrate(&User{})

	// 增
	/*var user User = User{
		Name: "lpx",
	} // 创建一个user，没赋值的字段会使用默认值

	fmt.Printf("主键id为：%d", user.ID) // 0
	result := db.Create(&user)      // 通过数据的指针来创建

	//获取主键id
	fmt.Printf("主键id为：%d\r\n", user.ID) // 1
	// 通过result获取插入的结果
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数*/

	//使用*string可以插入零值
	/*var empty string = ""
	db.Model(&User{ID: 2}).Updates(User{Email: &empty})*/

	/*// 批量插入，传入切片
	users := []User{{Name: "lpx1"}, {Name: "lpx2"}, {Name: "lpx3"}}
	db.Create(&users)
	// 也可以使用CreateInBatches批量插入
	db.CreateInBatches(users, 100) // 每次插入100条记录，直到插入完所有记录。为什么需要分批插入？因为一次性插入大量记录会导致内存占用过高，且mysql有最大包大小限制

	// map插入
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})

	// batch insert from `[]map[string]interface{}{}`
	db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu_1", "Age": 18},
		{"Name": "jinzhu_2", "Age": 20},
	})*/

	// 查

	// 检索单个对象
	// 获取第一条记录（主键升序）
	/*var user1 User
	result := db.First(&user1) // SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&user1) // SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 获取一条记录，没有指定排序字段
	db.Take(&user1) // SELECT * FROM users LIMIT 1;

	errors.Is(result.Error, gorm.ErrRecordNotFound) //判断错误是否是记录未找到

	// 根据主键检索
	db.First(&user1, 10) // SELECT * FROM users WHERE id = 10;

	db.First(&user1, "10") // SELECT * FROM users WHERE id = 10;

	result = db.Find(&user1, []int{1, 2, 3}) // SELECT * FROM users WHERE id IN (1,2,3);
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println("记录未找到")
	}*/
	/*
		var user = User{ID: 10}
		db.First(&user) // SELECT * FROM users WHERE id = 10;

		var result User
		db.Model(User{ID: 10}).First(&result) // SELECT * FROM users WHERE id = 10;*/

	var user User
	db.Where("Name = ?", "lpx").First(&user) //大小写不敏感，name，Name都可以
	// SELECT * FROM users WHERE name = 'lpx' ORDER BY id LIMIT 1;

	db.Where("name = ?", "lpx").First(&user)
	// SELECT * FROM users WHERE name = 'lpx' ORDER BY id LIMIT 1;
	db.Where(&User{Name: "lpx"}).First(&user)
	db.Where(map[string]interface{}{"name": "lpx"}).Find(&user)

	var users []User
	db.Where([]int64{20, 21, 22}).Find(&users) // SELECT * FROM users WHERE id IN (20, 21, 22);

	db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users) // SELECT * FROM users WHERE name = "jinzhu";

	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users) // SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20) // SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;
	// 检索全部对象
	/*var users []User
	result = db.Find(&users) // SELECT * FROM users;
	fmt.Printf("记录数：%d\r\n", result.RowsAffected)
	for _, user := range users {
		fmt.Println(user.ID)
	}*/

	// Update - 将 product 的 price 更新为 200
	//db.Model(&User{ID: 1}).Update("Name", "")
	//db.Model(&User{ID: 1}).Updates(User{Name: ""})

	db.Save(&User{Name: "lpx", Age: 100})        //INSERT INTO
	db.Save(&User{ID: 1, Name: "lpx", Age: 100}) //UPDATE

	// Delete - 删除 product   逻辑删除，可以看到是update语句

	// 带额外条件的删除
	db.Where("name = ?", "lpx").Delete(&user) // DELETE from emails where id = 10 AND name = "jinzhu";

	//db.Delete(&user1, 1)

}
