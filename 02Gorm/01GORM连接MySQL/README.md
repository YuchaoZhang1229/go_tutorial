## GORM连接数据库
### 1. 连接数据库
GORM officially supports the databases `MySQL`, `PostgreSQL`, `GaussDB`, `SQLite`, `SQL Server`, and `TiDB`
```go
import (
	"gorm.io/driver/mysql"         // MySQL
    // "gorm.io/driver/sqlite"     // SQLite
    // "gorm.io/driver/postgres"   // PostgreSQL
    // "gorm.io/driver/gaussdb"    // GaussDB
    // "gorm.io/driver/sqlserver"  // SQL Server
	"gorm.io/gorm"
)
```

### 2. 连接MySQL
```go
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

// 方法一：
dsn := "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 方法二：
db, err := gorm.Open(mysql.New(mysql.Config{
  DSN: "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
  DefaultStringSize: 256, // string 类型字段的默认长度
  DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
  DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
  DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
  SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
}), &gorm.Config{})
```
**dsn参数解析**
- root:123456 - 用户名:密码 
- @tcp(127.0.0.1:3306) - 使用TCP协议连接本地MySQL(127.0.0.1)的3306端口 
- /dbname - 要连接到的具体数据库的名称。在连接之前，这个数据库应该已经被创建
  - CREATE DATABASE IF NOT EXISTS dbname CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
- charset=utf8mb4 - 支持完整UTF-8编码, 支持存储所有的 Unicode 字符，包括表情符号（Emoji）
- parseTime=True - 将数据库中的时间类型解析为 Go 的 time.Time 类型\
- loc=Local - 使用本地时区

### 3. 代码示例
```go
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

```

