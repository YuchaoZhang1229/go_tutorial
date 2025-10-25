package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfo 结构体 --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//defer db.Close() // 确保程序退出前关闭连接

	// 创建表 自动迁移 (把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	u1 := UserInfo{
		ID:     1,
		Name:   "q1mi",
		Gender: "male",
		Hobby:  "蛙泳",
	}
	db.Create(&u1)

	// 查询
	var u2 UserInfo
	db.First(&u2) // 查询表中第一条数据保存到u2中
	fmt.Printf("u2: %v\n", u2)
	// 更新
	db.Model(&u2).Update("hobby", "双色球")
	// 删除
	db.Delete(&u2)

}
