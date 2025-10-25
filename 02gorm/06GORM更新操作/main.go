package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型
type User struct {
	Name   string
	Age    int
	Active bool
	gorm.Model
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

	//u := []User{
	//	{Name: "Alice", Age: 28, Active: true},
	//	{Name: "Bob", Age: 0, Active: false},
	//	{Name: "Charlie", Age: 18, Active: true},
	//	{Name: "Diana", Age: 0, Active: false},
	//	{Name: "Evan", Age: 16, Active: true},
	//}
	//
	//// 创建
	//db.Create(&u)
	// 查询
	var user User
	db.First(&user)

	// 更新
	// 1. 更新所有字段----------------------------------------------------
	user.Name = "七米"
	user.Age = 99
	db.Save(&user)
	fmt.Println("更新所有字段", user) // 默认会修改所有字段

	// 2. 更新单个字段
	// 根据user的id来更新对应行的字段
	db.Model(&user).Update("name", "小王子")
	fmt.Println("更新单个字段", user)

	db.Model(&user).Where("active=?", true).Update("name", "hello") // 根据给定条件更新单个字段

	fmt.Println("更新单个字段", user)

	// 3. 更新多个字段
	db.Model(&user).Updates(User{
		Name:   "struct",
		Age:    666,
		Active: true,
	})
	fmt.Println("更新多个字段-struct", user)

	db.Model(&user).Updates(map[string]interface{}{
		"name":   "map",
		"age":    888,
		"active": false,
	})
	fmt.Println("更新多个字段-map", user)

	// 4. 更新选定字段
	// 选择除 Active 外的所有字段（包括零值字段的所有字段）
	db.Model(&user).Select("Name", "Age", "Active").Omit("Age").Updates(User{Name: "jinzhu", Age: 100, Active: true})
	fmt.Println("更新选定字段", user)

	// 5. 批量更新
	db.Model(&User{}).Where("active=?", true).Update("name", "hello")

	// 6. 高级选项
	// 使用SQL表达式更新
	// 让users表中所有用户的年龄在原来的基础上+2
	db.Model(&User{}).Where("active=?", true).Update("age", gorm.Expr("age + ?", 2))

}
