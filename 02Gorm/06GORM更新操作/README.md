## GORM 查询操作
### 1. 检索单个对象
若主键有值则更新，无值则创建
```go
db.First(&user)
user.Name = "七米"
user.Age = 99
db.Save(&user)
```

### 2. 更新单个字段 
根据`&user`的`primary key`来找到对应行再更新字段;
```go
db.Model(&user).Update("name", "小王子")
db.Model(&user).Where("active=?", true).Update("name", "小王子")
```

### 3. 更新多个字段
Updates 方法支持 `struct`（非零）和 `map[string]interface{}` 参数
```go
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
```

### 4. 更新选定字段
```go
db.Model(&user).Select("Name").Updates(User{Name: "jinzhu", Age: 0})
db.Model(&user).Omit("Name").Updates(User{Name: "jinzhu", Age: 0})
```

### 5. 更新 Hook
支持的 hook 包括：BeforeSave, BeforeUpdate, AfterSave, AfterUpdate
```go
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
if u.Role == "admin" {
return errors.New("admin user not allowed to update")
}
return
}
```

### 6. 批量更新

**阻止全局更新**: 默认禁止无条件的批量更新，会返回 ErrMissingWhereClause错误。
```go
// 对符合条件的所有记录进行更新
db.Model(&User{}).Where("active = ?", "active").Updates(User{Name: "hello"})
```


**获取更新的记录数**: 通过 Result的 RowsAffected获取。

```go
result := db.Model(User{}).Where("active = ?", "active").Updates(User{Name: "hello"}); 
rowsAffected := result.RowsAffected
```

### 7. 高级选项
**使用SQL表达式更新**
```go
db.Model(&User{}).Where("active=?", true).Update("age", gorm.Expr("age + ?", 2))
```

**根据子查询进行更新**
`待续`

**不使用 Hook 和时间追踪**
```go
db.Model(&user).UpdateColumn("name", "hello")
db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
db.Model(&user).Select("name", "age").UpdateColumns(User{Name: "hello", Age: 0})
```
**返回修改行的数据**
`待续`

**检查字段是否有变更**
`待续`

**在 Update 时修改值**
`待续`