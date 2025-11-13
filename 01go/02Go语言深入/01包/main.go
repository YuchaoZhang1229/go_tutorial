// main.go
package main

import (
	"fmt"
	"github.com/myapp/pkg/calculator"
	"github.com/myapp/pkg/logger"
)

func main() {
	logger.Log("应用程序启动")

	result1 := calculator.Add(10, 5)
	result2 := calculator.Subtract(10, 5)

	fmt.Printf("加法结果: %.2f\n", result1)
	fmt.Printf("减法结果: %.2f\n", result2)
	fmt.Printf("总操作次数: %d\n", calculator.GetOperationCount())

	logger.Log("应用程序结束")
}
