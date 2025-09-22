## GORM 创建操作
### 1. 创建记录
- **功能**：向数据库插入单条新记录，会自动填充主键和追踪时间（CreatedAt/UpdatedAt）
- **语法**：
db.Create(&user)

### 2. 用指定的字段创建记录
- **功能**：使用`Select`或`Omit`来指定只创建某些字段或忽略某些字段。
- **语法**：
  - db.Select("Name", "Age").Create(&user)
  - db.Omit("Age").Create(&user)

### 3. 批量插入
- 功能：在创建记录的生命周期（BeforeSave, BeforeCreate, AfterSave, AfterCreate）中执行自定义逻辑。
- 语法：
```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) { 
	// 你的逻辑 
}
```

### 4. 根据 Map 创建
功能：直接使用`map[string]interface{}`或 `[]map[string]interface{}{}` 来创建记录。但注意，此方式不会执行钩子方法，也不会保存关联且不会回写主键。
语法：
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

### 5. 使用 SQL 表达式、Context Valuer 创建记录
`待续`

### 6. 高级选项
- 关联创建-`待续`
- 默认值-`待续`
- Upsert 及冲突-`待续`

