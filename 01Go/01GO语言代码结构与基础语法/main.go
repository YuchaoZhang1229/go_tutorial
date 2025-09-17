package main

import "fmt"

func main() {
	// 这是单行注释
	/*
		这是多行注释
	*/
	fmt.Println("hello world")
}

// 1. 直接运行
// go run .\main.go

// 2. 生成二级制文件再运行
// go build
// .\main.exe

// 3. 格式化go代码
// go fmt .\main.go
