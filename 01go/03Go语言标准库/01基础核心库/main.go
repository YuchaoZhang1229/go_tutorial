package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 用户结构体
type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

// 自定义错误
var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidAge   = errors.New("invalid age")
)

// fmt 库示例
func fmtExamples() {
	fmt.Println("=== fmt 库示例 ===")

	// 基本打印
	name := "Alice"
	age := 25
	fmt.Print("Print: 不换行打印\n")
	fmt.Println("Println: 自动换行")
	fmt.Printf("Printf: 格式化打印，姓名：%s，年龄：%d\n", name, age)

	// 格式化字符串
	formatted := fmt.Sprintf("格式化字符串：%s今年%d岁", name, age)
	fmt.Println(formatted)

	// 输入扫描
	fmt.Print("请输入你的名字和年龄（用空格分隔）: ")
	var inputName string
	var inputAge int
	_, err := fmt.Scan(&inputName, &inputAge)
	if err != nil {
		fmt.Printf("输入错误: %v\n", err)
	} else {
		fmt.Printf("你好 %s, 你今年 %d 岁\n", inputName, inputAge)
	}

	fmt.Println()
}

// strings 库示例
func stringsExamples() {
	fmt.Println("=== strings 库示例 ===")

	// 字符串比较
	fmt.Println("=== 字符串比较:")
	fmt.Println("Compare(\"abc\", \"abc\")", strings.Compare("abc", "abc")) // 0 (相等)
	fmt.Println("Compare(\"abc\", \"abd\")", strings.Compare("abc", "abd")) // -1 (第一个字符串小于第二个)
	fmt.Println("EqualFold(\"Go\", \"go\")", strings.EqualFold("Go", "go")) // true (不区分大小写的比较)
	fmt.Println()

	text := "Hello, World! Welcome to Go Programming!"
	fmt.Printf("原始文本: %s\n", text)
	fmt.Println()

	// 查找操作
	fmt.Println("=== 字符串查找:")
	fmt.Printf("包含 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("'World' 位置: %d\n", strings.Index(text, "World"))
	fmt.Printf("前缀检查: %t\n", strings.HasPrefix(text, "Hello"))
	fmt.Printf("后缀检查: %t\n", strings.HasSuffix(text, "!"))
	fmt.Println()

	// 替换和大小写
	fmt.Println("=== 替换和大小写:")
	replaced := strings.ReplaceAll(text, "Go", "Golang")
	fmt.Printf("替换后: %s\n", replaced)
	fmt.Printf("转大写: %s\n", strings.ToUpper(text))
	fmt.Printf("转小写: %s\n", strings.ToLower(text))
	fmt.Println()

	// 分割和连接
	fmt.Println("=== 分割和连接:")
	words := strings.Split(text, " ")
	fmt.Printf("分割单词: %v\n", words)
	fmt.Printf("连接字符串: %s\n", strings.Join(words, "-"))
	fmt.Println()

	// 修剪和重复
	// 修剪:Trim, TrimSpace, TrimPrefix, TrimSuffix
	fmt.Println("=== 修剪和重复:")
	spacedText := "   hello world   "
	fmt.Printf("修剪前后空格: '%s'\n", strings.TrimSpace(spacedText))
	fmt.Printf("重复字符串: %s\n", strings.Repeat("Go", 3))
	fmt.Println()

	// Builder 高效构建字符串
	fmt.Println("=== Builder 高效构建字符串:")
	var builder strings.Builder
	builder.WriteString("开始构建...")
	builder.WriteString(" 添加更多内容...")
	builder.WriteString(" 完成!")
	result := builder.String()
	fmt.Printf("Builder结果: %s\n", result)

	fmt.Println()
}

// strconv 库示例
func strconvExamples() {
	fmt.Println("=== strconv 库示例 ===")

	// 字符串转数字
	strNum := "123"
	num, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Printf("Atoi错误: %v\n", err)
	} else {
		fmt.Printf("字符串 '%s' 转数字: %d\n", strNum, num)
	}

	// 数字转字符串
	num2 := 456
	strNum2 := strconv.Itoa(num2)
	fmt.Printf("数字 %d 转字符串: '%s'\n", num2, strNum2)

	// 解析其他类型
	strFloat := "3.14"
	floatNum, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Printf("ParseFloat错误: %v\n", err)
	} else {
		fmt.Printf("字符串 '%s' 转浮点数: %.2f\n", strFloat, floatNum)
	}

	// 格式化其他类型
	boolStr := strconv.FormatBool(true)
	fmt.Printf("布尔值转字符串: %s\n", boolStr)

	// 解析布尔值
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Printf("ParseBool错误: %v\n", err)
	} else {
		fmt.Printf("字符串转布尔值: %t\n", b)
	}

	fmt.Println()
}

// errors 库示例
func errorsExamples() {
	fmt.Println("=== errors 库示例 ===")

	// 创建用户
	user, err := createUser("Bob", "bob@example.com", 30)
	if err != nil {
		fmt.Printf("创建用户失败: %v\n", err)
	} else {
		fmt.Printf("创建用户成功: %+v\n", user)
	}

	// 故意创建无效用户
	invalidUser, err := createUser("Charlie", "charlie@example.com", -5)
	if err != nil {
		fmt.Printf("创建用户失败: %v\n", err)
	} else {
		fmt.Printf("创建用户成功: %+v\n", invalidUser)
	}

	// 查找用户
	foundUser, err := findUser(1)
	if err != nil {
		fmt.Printf("查找用户失败: %v\n", err)
	} else {
		fmt.Printf("找到用户: %+v\n", foundUser)
	}

	// 查找不存在的用户
	notFoundUser, err := findUser(999)
	if err != nil {
		fmt.Printf("查找用户失败: %v\n", err)
	} else {
		fmt.Printf("找到用户: %+v\n", notFoundUser)
	}

	fmt.Println()
}

// 创建用户函数 - 演示错误处理
func createUser(name, email string, age int) (*User, error) {
	if age <= 0 || age > 150 {
		return nil, ErrInvalidAge
	}

	user := &User{
		ID:    len(name) + len(email), // 简单模拟ID生成
		Name:  name,
		Email: email,
		Age:   age,
	}

	return user, nil
}

// 查找用户函数 - 演示错误处理
func findUser(id int) (*User, error) {
	// 模拟数据库查找
	if id == 1 {
		return &User{
			ID:    1,
			Name:  "Test User",
			Email: "test@example.com",
			Age:   25,
		}, nil
	}

	return nil, ErrUserNotFound
}

func main() {
	fmt.Println("Go标准库学习Demo")
	fmt.Println("================\n")

	// 运行各个示例
	//fmtExamples()
	stringsExamples()
	strconvExamples()
	errorsExamples()

	fmt.Println("Demo结束！")
}
