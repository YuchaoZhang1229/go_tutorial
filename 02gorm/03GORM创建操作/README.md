## GORM 创建操作
### 1. 创建记录
- **功能**：向数据库插入单条新记录，会自动填充主键和追踪时间（CreatedAt/UpdatedAt）
- **语法**：

```go
db.Create(&user)
```

### 2. 用指定的字段创建记录
- **功能**：使用`Select`或`Omit`来指定只创建某些字段或忽略某些字段。
- **语法**：
```go
db.Select("Name", "Age").Create(&user)
db.Omit("Age").Create(&user)
```

### 3.批量插入
**功能**：通过切片批量插入记录，可指定每批次大小以提高效率。
**语法**：

```go
db.Create(&users)
db.CreateInBatches(&users, 100)
```

### 4. 创建钩子
- **功能**：在创建记录的生命周期（BeforeSave, BeforeCreate, AfterSave, AfterCreate）中执行自定义逻辑。
- **语法**：
```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) { 
	// 你的逻辑 
}
```

### 5. 根据 Map 创建
**功能**：直接使用`map[string]interface{}`或 `[]map[string]interface{}{}` 来创建记录。但注意，此方式不会执行钩子方法，也不会保存关联且不会回写主键。

**语法**：
```go
db.Model(&User{}).Create(map[string]interface{}{
"Name": "jinzhu", "Age": 18,
})

// batch insert from `[]map[string]interface{}{}`
db.Model(&User{}).Create([]map[string]interface{}{
{"Name": "jinzhu_1", "Age": 18},
{"Name": "jinzhu_2", "Age": 20},
})
```

### 6. 使用 SQL 表达式、Context Valuer 创建记录
`待续`

### 7. 高级选项
#### 关联创建
`待续`

#### 默认值

通过结构体tagd `efault` 为字段定义默认值。

注意：零值（如0、""、false）都还会触发默认值，如果想避免的话，建议用指针或Scanner/Valuer接口。

语法：
```go
type User struct { 
	Name *string `gorm:"default:'小王子'"`
}

type User struct {
    Name sql.NullString `gorm:"default:'小王子'"`
}
```


#### Upsert 及冲突

`待续`


### 8. 代码示例
```go
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
	db.Create(&u1)                                                                // 在数据库中创建了一行数据

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

```

