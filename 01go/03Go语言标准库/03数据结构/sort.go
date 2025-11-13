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

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func customSort() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
	}

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
