package main

import "fmt"

func main() {
	// 1. 不同的数组声明和初始化方式
	fmt.Println("=== 数组声明方式 ===")

	// 方式1: 先声明后赋值
	var arr1 [10]int
	arr1[0] = 1
	arr1[1] = 2

	// 方式2: 声明时初始化
	arr2 := [5]int{1, 2, 3, 4, 5}

	// 方式3: 让编译器计算数组长度
	arr3 := [...]int{1, 2, 3, 4, 5, 6}

	// 方式4: 部分初始化，其余为0值
	arr4 := [5]int{1, 2} // [1, 2, 0, 0, 0]

	// 方式5: 按索引初始化
	arr5 := [5]int{0: 10, 3: 40} // [10, 0, 0, 40, 0]

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v\n", arr3)
	fmt.Printf("arr4: %v\n", arr4)
	fmt.Printf("arr5: %v\n", arr5)

	// 2. 数组遍历
	fmt.Println("\n=== 数组遍历 ===")

	// 方式1: 传统for循环
	fmt.Print("传统for循环: ")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("%d ", arr2[i])
	}
	fmt.Println()

	// 方式2: range遍历（推荐）
	fmt.Print("range遍历: ")
	for index, value := range arr2 {
		fmt.Printf("arr2[%d]=%d ", index, value)
	}
	fmt.Println()

	// 方式3: 只获取值
	fmt.Print("只获取值: ")
	for _, value := range arr2 {
		fmt.Printf("%d ", value)
	}
	fmt.Println()

	// 3. 数组操作
	fmt.Println("\n=== 数组操作 ===")

	// 数组长度
	fmt.Printf("arr2的长度: %d\n", len(arr2))

	// 数组比较（只有相同类型和长度的数组可以比较）
	arr6 := [3]int{1, 2, 3}
	arr7 := [3]int{1, 2, 3}
	arr8 := [3]int{1, 2, 4}
	fmt.Printf("arr6 == arr7: %t\n", arr6 == arr7)
	fmt.Printf("arr6 == arr8: %t\n", arr6 == arr8)

	// 4. 多维数组
	fmt.Println("\n=== 多维数组 ===")

	// 二维数组声明和初始化
	var matrix1 [2][3]int
	matrix1[0] = [3]int{1, 2, 3}
	matrix1[1] = [3]int{4, 5, 6}

	matrix2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("matrix1:")
	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[i]); j++ {
			fmt.Printf("%d ", matrix1[i][j])
		}
		fmt.Println()
	}

	fmt.Println("matrix2:")
	for _, row := range matrix2 {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	// 5. 数组作为函数参数（值传递）
	fmt.Println("\n=== 数组作为函数参数 ===")
	testArray := [3]int{1, 2, 3}
	fmt.Printf("修改前: %v\n", testArray)
	modifyArray(testArray) // 数组是值类型，函数内修改不影响原数组
	fmt.Printf("修改后: %v\n", testArray)

	// 6. 数组的局限性
	demonstrateArrayLimitations()
}

// 修改数组的函数
func modifyArray(arr [3]int) {
	arr[0] = 100
	fmt.Printf("函数内修改: %v\n", arr)
}

// 演示数组的局限性
func demonstrateArrayLimitations() {
	fmt.Println("\n=== 数组的局限性 ===")

	// 数组长度是类型的一部分，不同长度的数组是不同的类型
	arr1 := [3]int{1, 2, 3}
	arr2 := [4]int{1, 2, 3, 4}

	fmt.Printf("arr1 类型: %T\n", arr1)
	fmt.Printf("arr2 类型: %T\n", arr2)

	// 数组大小固定，无法动态扩展
	fmt.Println("数组大小固定，无法添加更多元素")

	// 实际开发中更多使用切片(slice)
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5) // 切片可以动态扩展
	fmt.Printf("切片可以扩展: %v\n", slice)
}
