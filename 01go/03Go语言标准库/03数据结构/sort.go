package main

import (
	"fmt"
	"sort"
)

func builtInSort() {
	// 整数排序
	nums := []int{4, 2, 7, 1, 9}
	sort.Ints(nums)
	fmt.Println("排序后整数:", nums)

	// 字符串排序
	strings := []string{"banana", "apple", "cherry"}
	sort.Strings(strings)
	fmt.Println("排序后字符串:", strings)

	// 检查是否已排序
	fmt.Println("是否已排序:", sort.IntsAreSorted(nums))
}

type Person struct {
	Name string
	Age  int
}

// 按年龄排序
type ByAge []Person

// 实现sort.Interface
func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func customSort() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
	}

	// 排序过程大致如下：
	// 1. 调用 Len() 知道有3个元素
	// 2. 多次调用 Less() 比较元素
	//    - Less(0,1): 25<20? false → Bob应该在Alice前面
	//    - Less(1,2): 20<30? true → Bob应该在Charlie前面
	// 3. 根据需要调用 Swap() 调整位置
	sort.Sort(ByAge(people))
	fmt.Println("\n按年龄排序:")
	for _, p := range people {
		fmt.Printf("%s: %d岁\n", p.Name, p.Age)
	}

	// 使用sort.Slice进行更简单的自定义排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name > people[j].Name // 按姓名倒序
	})

	fmt.Println("\n按姓名倒序:")
	for _, p := range people {
		fmt.Printf("%s: %d岁\n", p.Name, p.Age)
	}
}

func main() {
	builtInSort()
	customSort()

}
