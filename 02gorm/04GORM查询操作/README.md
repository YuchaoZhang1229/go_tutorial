## GORM 查询操作
### 1. 检索单个对象
```go
var user User
db.First(&user) // 获取第一条记录（主键升序）
db.Last(&user)  // 获取最后一条记录（主键降序）
db.Take(&user)  // 获取一条记录，没有指定排序字段
```

### 2. 检索全部对象
```go
var users []User
db.Find(&users)
```

### 3. 条件（Where）
```go
db.Where("name = ?", "jinzhu").Find(&users) // string 条件
db.Where(&User{Name: "jinzhu", Age: 0}.Find(&users) // struct条件(仅非零字段)
db.Where(map[string]interface{}{"name":"jinzhu", "age":20}).Find(&users) // Map条件
db.Where([]int64{20, 21, 22}).Find(&users) // 主键切片
db.Where(&User{Name: "jinzhu"}, "Age").Find(&users) // 指定结构体查询条件
db.Find(&users, "name = ?", "jinzhu") // 内联条件
db.Not("name", "jinzhu").First(&user) // Not条件
db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users) // Or条件
```

### 4. 选择特定字段 (Select)

```go
// 指定要查询的字段:
db.Select("name", "age").Find(&users)
db.Select([]string{"name", "age"}).Find(&users)
```


### 5. 排序 (Order)

```go
// 单字段排序
db.Order("age desc").Find(&users)  
// 多字段排序
db.Order("age desc, name").Find(&users)
db.Order("age desc").Order("name").Find(&users) 
```


### 6. Limit & Offset

```go
db.Limit(3).Find(&users)  // 限制数量
db.Offset(5).Find(&users) // 偏移量
db.Limit(10).Offset(5).Find(&users)  // 组合使用（常用于分页）
db.Limit(10).Find(&users1).Limit(-1).Find(&users2) // 取消 Limit 或 Offset 条件
```


### 7. Group By & Having
`待续`

### 8. Distinct
```go
// 查询不重复的值
db.Distinct("name").Find(&users)
```

### 7. Joins
`待续`


### 7. Group By & Having
```go
// 将结果扫描到结构体
db.Table("users").Select("name, age").Where("name = ?", "jinzhu").Scan(&result) 
// 原生SQL
db.Raw("SELECT name, age FROM users WHERE name = ?", "jinzhu").Scan(&result)
```


