package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型
type User struct {
	gorm.Model
	Description string
	Name        sql.NullString `gorm:"default:'小王子'"`
	Age         int
}

func main() {
	// 连接MySQL数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 1.创建记录-单条
	u1 := User{Description: "1.创建记录-单条", Name: sql.NullString{"", true}, Age: 18} // 在代码层面创建一个User对象
	db.Create(&u1)                                                                      // 在数据库中创建了一行数据

	// 2.用指定的字段创建记录-db.Select
	u2 := User{Description: "2.用指定的字段创建记录-db.Select", Name: sql.NullString{"q1mi", true}, Age: 18}
	db.Select("Description", "Name").Create(&u2)

	// 3.用指定的字段创建记录-db.Omit
	u3 := User{Description: "3.用指定的字段创建记录-db.Omit", Age: 18} // 没有指定Name, 会被设置成默认值'小王子'
	db.Omit("Age").Create(&u3)

	// 4.批量插入-db.Create
	u4 := []User{
		{Description: "4.批量插入", Name: sql.NullString{"q1mi_1", true}, Age: 18},
		{Description: "4.批量插入", Name: sql.NullString{"q1mi_2", true}, Age: 19},
	}
	//db.Create(&u4)
	db.CreateInBatches(&u4, 5)

	// 5.根据Map创建
	db.Model(&User{}).Create(map[string]interface{}{
		"Description": "5.根据Map创建1",
		"Name":        sql.NullString{"q1mi", true},
		"Age":         18,
	})

	db.Model(&User{}).Create([]map[string]interface{}{
		{"Description": "5.根据Map创建2", "Name": sql.NullString{"q1mi_1", true}, "Age": 18},
		{"Description": "5.根据Map创建2", "Name": sql.NullString{"q1mi_2", true}, "Age": 20},
	})
	// 注意当使用map来创建时，钩子方法不会执行，关联不会被保存且不会回写主键

}
