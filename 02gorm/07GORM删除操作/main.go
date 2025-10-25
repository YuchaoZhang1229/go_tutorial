package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 定义模型
type User struct {
	gorm.Model
	Name   string
	Age    int
	Active bool
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
	//	{Name: "jinzhu", Age: 28, Active: true},
	//	{Name: "jinzhu3", Age: 0, Active: false},
	//	{Name: "!jinzhu", Age: 18, Active: true},
	//	{Name: "jinzhu!", Age: 99, Active: false},
	//	{Name: "jinzhu@", Age: 16, Active: true},
	//}
	//
	//// 创建
	//db.Create(&u)

	// 查询
	var user User
	db.First(&user)

	// 1. 删除一条记录 ---------------------------------------------------
	db.Delete(&user)
	db.Where("name = ?", "jinzhu").Delete(&user)

	// 2. 根据主键删除
	db.Delete(&User{}, 3)
	//db.Delete(&User{}, "3")
	//db.Delete(&User{}, []int{1, 2, 3})

	// 3. 钩子函数-待续

	// 4. 批量删除
	//db.Where("name LIKE ?", "%jinzhu%").Delete(&User{})
	//db.Delete(&User{}, "name LIKE ?", "%jinzhu%")

	// 返回删除行的数据
	var users []User
	db.Clauses(clause.Returning{}).Where("name = ?", "jinzhu").Delete(&users)
	fmt.Println(users)

}
