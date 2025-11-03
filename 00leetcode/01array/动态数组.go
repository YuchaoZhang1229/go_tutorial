package main

import (
	"fmt"
	"time"
)

func basicOperations() {
	fmt.Println("\n=== 切片基本操作 ===")

	// 创建切片
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("初始: %v, len=%d, cap=%d\n", nums, len(nums), cap(nums))

	// 1. 访问元素
	fmt.Printf("第一个元素: %d\n", nums[0])
	fmt.Printf("最后一个元素: %d\n", nums[len(nums)-1])

	// 2. 修改元素
	nums[0] = 100
	fmt.Printf("修改第一个元素为100后: %v\n", nums)

	// 3. 切片遍历
	fmt.Print("索引遍历: ")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()

	fmt.Print("range遍历: ")
	for index, value := range nums {
		fmt.Printf("nums[%d]=%d ", index, value)
	}
	fmt.Println()

	// 4. 切片截取
	subSlice := nums[1:4] // 索引1到3 [2,3,4]
	fmt.Printf("子切片: %v\n", subSlice)

	// 5. 检查切片是否为空
	var emptySlice []int
	fmt.Printf("emptySlice为空: %t\n", emptySlice == nil)
	fmt.Printf("emptySlice长度: %d\n", len(emptySlice))
}

func advancedOperations() {
	fmt.Println("\n=== 切片高级操作 ===")

	// 创建初始切片
	arr := make([]int, 0)

	fmt.Println("1. 尾部追加操作")
	for i := 0; i < 5; i++ {
		arr = append(arr, i)
		fmt.Printf("追加 %d: %v, len=%d, cap=%d\n", i, arr, len(arr), cap(arr))
	}

	fmt.Println("\n2. 中间插入操作")
	// 在索引2的位置插入666
	arr = append(arr[:2], append([]int{666}, arr[2:]...)...)
	fmt.Printf("中间插入后: %v\n", arr)

	fmt.Println("\n3. 头部插入操作")
	arr = append([]int{-1, -2}, arr...)
	fmt.Printf("头部插入后: %v\n", arr)

	fmt.Println("\n4. 删除操作")
	// 删除尾部元素
	arr = arr[:len(arr)-1]
	fmt.Printf("删除尾部后: %v\n", arr)

	// 删除头部元素
	arr = arr[1:]
	fmt.Printf("删除头部后: %v\n", arr)

	// 删除中间元素（索引2）
	arr = append(arr[:2], arr[3:]...)
	fmt.Printf("删除中间后: %v\n", arr)

	fmt.Println("\n5. 切片拷贝")
	// 浅拷贝（共享底层数组）
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2[0] = 100
	fmt.Printf("浅拷贝 - slice1: %v, slice2: %v\n", slice1, slice2)

	// 深拷贝
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, len(slice3))
	copy(slice4, slice3)
	slice4[0] = 200
	fmt.Printf("深拷贝 - slice3: %v, slice4: %v\n", slice3, slice4)

	fmt.Println("\n6. 切片比较")
	// 切片不能直接比较，需要手动比较元素
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Printf("切片相等: %t\n", sliceEqual(a, b))

	fmt.Println("\n7. 切片过滤")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := filter(numbers, func(x int) bool {
		return x%2 == 0
	})
	fmt.Printf("偶数过滤: %v -> %v\n", numbers, evenNumbers)
}

// 切片相等比较
func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 切片过滤
func filter(slice []int, test func(int) bool) []int {
	result := make([]int, 0)
	for _, v := range slice {
		if test(v) {
			result = append(result, v)
		}
	}
	return result
}

func sliceInternals() {
	fmt.Println("\n=== 切片内部原理 ===")

	// 演示切片扩容机制
	fmt.Println("切片扩容演示:")
	var s []int
	for i := 0; i < 20; i++ {
		oldCap := cap(s)
		s = append(s, i)
		newCap := cap(s)
		if newCap != oldCap {
			fmt.Printf("长度: %2d, 新容量: %2d (扩容系数: %.2f)\n",
				len(s), newCap, float64(newCap)/float64(oldCap))
		}
	}

	fmt.Println("\n内存共享问题演示:")
	// 多个切片可能共享底层数组
	original := []int{1, 2, 3, 4, 5}
	sliceA := original[1:4] // [2,3,4]
	sliceB := original[2:5] // [3,4,5]

	fmt.Printf("修改前 - original: %v, sliceA: %v, sliceB: %v\n", original, sliceA, sliceB)

	// 修改共享的底层数组
	sliceA[0] = 100
	fmt.Printf("修改sliceA[0]后 - original: %v, sliceA: %v, sliceB: %v\n", original, sliceA, sliceB)

	// 避免内存共享问题
	fmt.Println("\n避免内存共享:")
	safeCopy := make([]int, len(original))
	copy(safeCopy, original)
	safeCopy[0] = 999
	fmt.Printf("深拷贝后 - original: %v, safeCopy: %v\n", original, safeCopy)
}

// 性能测试：预分配容量 vs 动态扩容
func performanceTest() {
	fmt.Println("\n=== 性能对比: 预分配 vs 动态扩容 ===")

	// 方法1: 动态扩容（性能较差）
	start := time.Now()
	var slice1 []int
	for i := 0; i < 100000; i++ {
		slice1 = append(slice1, i)
	}
	time1 := time.Since(start)

	// 方法2: 预分配容量（性能优化）
	start = time.Now()
	slice2 := make([]int, 0, 100000)
	for i := 0; i < 100000; i++ {
		slice2 = append(slice2, i)
	}
	time2 := time.Since(start)

	fmt.Printf("动态扩容耗时: %v\n", time1)
	fmt.Printf("预分配耗时: %v\n", time2)
	fmt.Printf("性能提升: %.2f%%\n", (float64(time1-time2)/float64(time1))*100)
}

func practicalExamples() {
	fmt.Println("\n=== 实战示例 ===")

	// 示例1: 栈实现
	fmt.Println("1. 切片实现栈:")
	stack := []int{}

	// 入栈
	stack = append(stack, 1) // push 1
	stack = append(stack, 2) // push 2
	stack = append(stack, 3) // push 3
	fmt.Printf("入栈后: %v\n", stack)

	// 出栈
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Printf("出栈: %d, 剩余: %v\n", top, stack)

	// 示例2: 队列实现（简单版）
	fmt.Println("\n2. 切片实现队列:")
	queue := []int{}

	// 入队
	queue = append(queue, 1) // enqueue 1
	queue = append(queue, 2) // enqueue 2
	queue = append(queue, 3) // enqueue 3
	fmt.Printf("入队后: %v\n", queue)

	// 出队
	front := queue[0]
	queue = queue[1:]
	fmt.Printf("出队: %d, 剩余: %v\n", front, queue)

	// 示例3: 切片批处理
	fmt.Println("\n3. 批量操作:")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 批量删除前3个元素
	data = data[3:]
	fmt.Printf("删除前3个: %v\n", data)

	// 批量添加多个元素
	data = append(data, []int{11, 12, 13}...)
	fmt.Printf("批量添加: %v\n", data)

	// 示例4: 二维切片
	fmt.Println("\n4. 二维切片:")
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
		for j := range matrix[i] {
			matrix[i][j] = i*3 + j + 1
		}
	}
	fmt.Println("二维矩阵:")
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func bestPractices() {
	fmt.Println("\n=== 最佳实践总结 ===")

	// 1. 预分配容量
	fmt.Println("1. 预分配容量:")
	// 不好：可能多次扩容
	var badSlice []int
	for i := 0; i < 1000; i++ {
		badSlice = append(badSlice, i)
	}

	// 好：一次性分配足够容量
	goodSlice := make([]int, 0, 1000)
	for i := 0; i < 1000; i++ {
		goodSlice = append(goodSlice, i)
	}

	// 2. 避免内存泄漏
	fmt.Println("\n2. 避免内存泄漏:")
	bigData := make([]int, 1000000)
	// 错误：大数组无法被GC
	// smallPart := bigData[0:10]

	// 正确：复制需要的数据
	smallPart := make([]int, 10)
	copy(smallPart, bigData[0:10])
	bigData = nil // 释放大数组

	// 3. 使用完整切片表达式控制容量
	fmt.Println("\n3. 控制切片容量:")
	data := []int{1, 2, 3, 4, 5}
	limitedSlice := data[1:3:3] // len=2, cap=2
	fmt.Printf("受控切片: len=%d, cap=%d\n", len(limitedSlice), cap(limitedSlice))

	// 4. 批量操作优化
	fmt.Println("\n4. 批量操作优化:")
	// 批量追加
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...) // 一次追加多个元素

	// 5. 安全遍历和修改
	fmt.Println("\n5. 安全遍历:")
	numbers := []int{1, 2, 3, 4, 5}

	// 遍历时修改原始切片是危险的
	// 应该创建新切片或使用索引遍历
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			numbers = append(numbers[:i], numbers[i+1:]...)
			i-- // 调整索引
		}
	}
}

func main() {
	basicOperations()
	advancedOperations()
	//sliceInternals()
	//performanceTest()
	//practicalExamples()
	//bestPractices()
}
