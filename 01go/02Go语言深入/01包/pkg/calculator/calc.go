// pkg/calculator/calc.go
package calculator

import "fmt"

// 包级变量
var operationCount int

func init() {
	fmt.Println("计算器包初始化完成")
}

// 导出函数
func Add(a, b float64) float64 {
	operationCount++
	return a + b
}

func Subtract(a, b float64) float64 {
	operationCount++
	return a - b
}

// 导出函数获取操作次数
func GetOperationCount() int {
	return operationCount
}
