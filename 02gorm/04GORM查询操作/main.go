package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型
type User struct {
	ID   int
	Name string
	Age  int
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
	//	{Name: "Alice", Age: 28},
	//	{Name: "Bob", Age: 0},
	//	{Name: "Charlie", Age: 18},
	//	{Name: "Diana", Age: 0},
	//	{Name: "Evan", Age: 16},
	//}
	//
	//// 创建
	//db.Create(&u)

	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("--------------------1. 检索单个对象--------------------")
	// 获取第一条记录（主键升序）
	// result.RowsAffected 返回找到的记录数;
	//  result.Error returns error or nil
	var user User
	result := db.First(&user) // SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println("user_struct:", user, "result.RowsAffected:", result.RowsAffected, "result.Error:", result.Error)

	result1 := map[string]interface{}{}
	db.Model(&User{}).First(&result1)
	fmt.Println("user_map:", result1)

	// 获取一条记录，没有指定排序字段
	var user2 User
	db.Take(&user2) // SELECT * FROM users LIMIT 1;
	fmt.Println("user2:", user2)

	// 获取最后一条记录（主键降序）
	var user3 User
	db.Last(&user3) // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println("user3:", user3)

	// 根据主键检索
	fmt.Println("----------根据主键检索----------")

	var user4 User
	db.First(&user4, 3) // SELECT * FROM users WHERE id = 3;
	//db.First(&user4, "3")
	fmt.Println("user4:", user4)

	var user5 User
	db.First(&user5, []int{3, 4, 5}) // SELECT * FROM users WHERE id IN (3, 4, 5);
	fmt.Println("user5:", user5)

	// 当目标对象有一个主键值时，将使用主键构建查询条件
	var user6 = User{ID: 3}
	db.First(&user6) // SELECT * FROM users WHERE id = 3;
	fmt.Println("user6:", user6)

	var user7 User
	db.Debug().Model(User{ID: 3}).First(&user7) // SELECT * FROM users WHERE id = 3;
	fmt.Println("user7:", user7)

	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("--------------------2. 检索全部对象--------------------")
	var user8 []User
	result2 := db.Find(&user8) // SELECT * FROM users;
	fmt.Println("user8:", user8, "result.RowsAffected:", result2.RowsAffected, "result.Error:", result2.Error)

	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("--------------------3. 条件--------------------")
	fmt.Println("----------string条件----------")
	// =
	var user9 []User
	db.Where("name = ?", "Charlie").Find(&user9)
	fmt.Println("user9:", user9)

	// !=
	var user10 []User
	db.Where("name <> ?", "Charlie").Find(&user10) // SELECT * FROM users WHERE name != 'Charlie';
	fmt.Println("user10:", user10)

	// IN
	var user11 []User
	db.Where("name IN ?", []string{"Alice", "Diana"}).Find(&user11)
	fmt.Println("user11:", user11)

	// LIKE
	var user12 []User
	db.Where("name LIKE ?", "%a%").Find(&user12)
	fmt.Println("user12:", user12)

	// AND
	var user13 []User
	db.Where("name = ? AND age >= ?", "Charlie", "18").Find(&user13)
	fmt.Println("user13:", user13)

	// >
	var user14 []User
	db.Where("age  > ?", 18).Find(&user14)
	fmt.Println("user14:", user14)

	// Between
	var user15 []User
	db.Where("age BETWEEN ? AND ?", 1, 18).Find(&user15)
	fmt.Println("user15:", user15)

	var user16 = User{Age: 18}
	db.Where("id = ?", 3).First(&user16)
	// SELECT * FROM users WHERE id = 3 and age = 18 ORDER BY id ASC LIMIT 1
	fmt.Println("user16:", user16)

	fmt.Println("----------Struct & Map条件----------")

	// Struct
	var user17 []User
	db.Where(&User{Name: "Charlie", Age: 18}).Find(&user17)
	fmt.Println("user17:", user17)

	// Map
	var user18 []User
	db.Where(map[string]interface{}{"name": "Charlie", "age": 18}).Find(&user18)
	fmt.Println("user18:", user18)

	// Slice of primary keys
	var user19 []User
	db.Where([]int64{3, 5}).Find(&user19)
	fmt.Println("user19:", user19)

	// 当使用结构体进行查询时，GORM 只会查询非零值字段。
	// 这意味着如果字段的值为 0、''（空字符串）、false 或其他零值，该字段将不会被用于构建查询条件，例如
	var user20 []User
	db.Where(&User{Name: "Diana", Age: 0}).Find(&user20) // SELECT * FROM users WHERE name = "Diana";
	fmt.Println("user20:", user20)

	// 要在查询条件中包含零值，您可以使用 map，它会将所有键值对作为查询条件包含在内，例如：
	var user21 []User
	db.Where(map[string]interface{}{"Name": "Diana", "Age": 0}).Find(&user21) // SELECT * FROM users WHERE name = "Diana" AND age = 0;
	fmt.Println("user21:", user21)

	fmt.Println("----------指定结构体查询字段----------")
	// 当使用结构体进行查询时，您可以通过向 Where()方法传递相关字段名或其数据库列名来指定查询条件中应使用该结构体的哪些特定值，例如：
	var user22 []User
	db.Where(&User{Name: "Diana"}, "name", "Age").Find(&user22) // 指定查询条件用name和age
	// SELECT * FROM users WHERE name = "Diana" AND age = 0;
	fmt.Println("user22:", user22)

	var user23 []User
	db.Where(&User{Name: "Bob"}, "Age").Find(&user23) // 指定查询条件只用age
	// SELECT * FROM users WHERE age = 0;
	fmt.Println("user23:", user23)

	fmt.Println("----------内联条件----------")
	// 查询条件可以内联到 First和 Find之类的方法中，其方式与 Where类似。
	var user24 []User
	db.Find(&user24, "name = ?", "Charlie") //  SELECT * FROM `users` WHERE name = 'Charlie'
	fmt.Println("user24:", user24)

	var user25 []User
	db.Find(&user25, "name <> ? AND age > ?", "Charlie", 18) //  SELECT * FROM `users` WHERE name <> 'Charlie' AND age > 18
	fmt.Println("user25:", user25)

	// Struct
	var user26 []User
	db.Find(&user26, User{Age: 18}) // SELECT * FROM users WHERE age = 18;
	fmt.Println("user26:", user26)

	// Map
	var user27 []User
	db.Find(&user27, map[string]interface{}{"age": 18}) // SELECT * FROM users WHERE age = 18;
	fmt.Println("user27:", user27)

	fmt.Println("----------Not 条件----------")
	// Build NOT conditions, works similar to Where
	var user28 []User
	db.Not("name = ?", "Alice").First(&user28) //  SELECT * FROM `users` WHERE NOT name = 'Alice' ORDER BY `users`.`id` LIMIT 1
	fmt.Println("user28:", user28)

	// Not In
	var user29 []User
	db.Not(map[string]interface{}{"name": []string{"Diana", "Diana"}}).Find(&user29)
	// SELECT * FROM users WHERE name NOT IN ("Diana", "Diana");
	fmt.Println("user29:", user29)

	// Struct
	var user30 []User
	db.Not(User{Name: "Alice", Age: 18}).Find(&user30)
	// SELECT * FROM users WHERE name <> "Alice" AND age <> 18;
	fmt.Println("user30:", user30)

	// Not In slice of primary keys
	var user31 []User
	db.Not([]int64{1, 2, 3}).First(&user31)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
	fmt.Println("user31:", user31)

	fmt.Println("----------Or 条件----------")
	var user32 []User
	db.Where("name = ?", "Charlie").Or("name = ?", "Alice").Find(&user32)
	// SELECT * FROM users WHERE name = 'Charlie' OR name = 'Alice';
	fmt.Println("user32:", user32)

	// Struct
	var user33 []User
	db.Where("name = 'Alice'").Or(User{Name: "Charlie", Age: 18}).Find(&user33)
	// SELECT * FROM users WHERE name = 'Alice' OR (name = 'Charlie' AND age = 18);
	fmt.Println("user33:", user33)

	// Map
	var user34 []User
	db.Where("name = 'Alice'").Or(map[string]interface{}{"name": "Charlie", "age": 18}).Find(&user34)
	// SELECT * FROM users WHERE name = 'Alice' OR (name = 'Charlie' AND age = 18);
	fmt.Println("user34:", user34)

	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("---------------4. 选择特定字段---------------")
	// Select 允许您指定要从数据库中检索的字段。否则，GORM 默认会选择所有字段

	var user35 []User
	db.Select("name", "age").Find(&user35) // SELECT name, age FROM users;
	fmt.Println("user35:", user35)

	var user36 []User
	db.Select([]string{"name", "age"}).Find(&user36) // SELECT name, age FROM users;
	fmt.Println("user36:", user36)

	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("---------------5. 排序---------------")
	var user37 []User
	db.Order("age desc, name").Find(&user37) // SELECT * FROM users ORDER BY age desc, name;
	fmt.Println("user37:", user37)

	// Multiple orders
	var user38 []User
	db.Order("age desc").Order("name").Find(&user38) // SELECT * FROM users ORDER BY age desc, name;
	fmt.Println("user38:", user38)

	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("---------------6. Limit & Offset---------------")
	var user39 []User
	db.Limit(3).Find(&user39) // SELECT * FROM users LIMIT 3;
	fmt.Println("user39:", user39)

	// Cancel limit condition with -1
	var user40 []User
	var user41 []User
	db.Limit(2).Find(&user40).Limit(-1).Find(&user41)
	// SELECT * FROM users LIMIT 3; (users1)
	// SELECT * FROM users; (users2)
	fmt.Println("user40:", user40)
	fmt.Println("user41:", user41)

	var user42 []User
	db.Limit(4).Offset(2).Find(&user42) // SELECT * FROM `users` LIMIT 4 OFFSET 2;
	fmt.Println("user42:", user42)

	fmt.Println("-----------------------------------------------------------------------------------------------")
	fmt.Println("---------------6. Group By & Having---------------")
	var user43 []User
	fmt.Println("user43:", user43)

	fmt.Println("---------------7. Distinct---------------")
	var user44 []User
	db.Distinct("age").Order("age desc").Find(&user44)
	fmt.Println("user44:", user44)

	fmt.Println("---------------8. Joins---------------")
	fmt.Println("----------Joins 预加载----------")
	// 您可以使用Joins预加载关联，通过单条SQL实现，例如：
	// db.Joins("Company").Find(&users)

	// inner join
	//db.InnerJoins("Company").Find(&users)

	// Join with conditions
	//db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)
	fmt.Println("----------Joins 一个衍生表----------")
	// You can also use Joins to join a derived table.

	fmt.Println("---------------9. Scan---------------")
	// Scanning results into a struct works similarly to the way we use Find
	type Result struct {
		Name string
		Age  int
	}

	var res1 []Result
	db.Table("users").Select("name", "age").Where("name = ?", "Charlie").Scan(&res1)
	fmt.Println("res1:", res1)

	// Raw SQL
	var res2 []Result
	db.Raw("SELECT name, age FROM users WHERE name = ?", "Charlie").Scan(&res2)
	fmt.Println("res2:", res2)

}
