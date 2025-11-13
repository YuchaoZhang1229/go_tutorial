package main

import "fmt"

// 1. 基本函数定义与调用
func add(a, b int) int {
	return a + b
}

// 2. 多返回值函数
func swap(x, y string) (string, string) {
	return y, x
}

// 3. 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 裸返回，自动返回x,y
}

// 4. 闭包示例 - 计数器
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 5. 闭包示例 - 斐波那契数列
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	// 基本函数调用
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// 多返回值调用
	a, b := swap("hello", "world")
	fmt.Printf("swap: %s %s\n", a, b)

	// 命名返回值调用
	x, y := split(17)
	fmt.Printf("split 17: %d, %d\n", x, y)

	// 闭包使用 - 计数器
	count := counter()
	fmt.Println("Counter:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", count())
	}
	fmt.Println()

	// 闭包使用 - 斐波那契
	fib := fibonacci()
	fmt.Println("Fibonacci:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()
}
