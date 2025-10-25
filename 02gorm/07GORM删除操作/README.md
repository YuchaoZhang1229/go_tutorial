## GORM 删除操作
### 1. 删除一条记录
删除单条记录，需指定主键或条件
```go
db.Delete(&user)
db.Where("name = ?", "jinzhu").Delete(&user)
```

### 2. 根据主键删除
```go
db.Delete(&User{}, 3)
db.Delete(&User{}, "3")
db.Delete(&User{}, []int{1, 2, 3})
```


### 3. 钩子函数
对于删除操作，GORM 支持 BeforeDelete、AfterDelete Hook，在删除记录时会调用这些方法，查看 Hook 获取详情
```go
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
    if u.Role == "admin" {
        return errors.New("admin user not allowed to delete")
    }
    return
}
```
### 4. 批量删除
一次性删除多条记录，支持条件筛选或主键列表
```go
db.Where("name LIKE ?", "%jinzhu%").Delete(&User{})
db.Delete(&users, []int{1,2,3})
```
阻止全局删除
```go
db.Delete(&User{})// 会返回 gorm.ErrMissingWhereClause错误
db.Where("1 = 1").Delete(&User{})// 允许
```
返回删除行的数据
```go
db.Clauses(clause.Returning{}).Where("role = ?", "admin").Delete(&users)
```

### 5. 软删除
非物理删除，通过设置 DeletedAt时间戳标记删除，默认查询不可见
```go
db.Unscoped().Find(&users)// 查询包括软删除的记录
db.Unscoped().Delete(&user)// 永久删除
```