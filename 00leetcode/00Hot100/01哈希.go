package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

// Leetcode 1 两数之和
func twoSum(nums []int, target int) []int {
	maps := make(map[int]int, len(nums))

	for i, v := range nums {
		preIndex, ok := maps[target-v]
		if ok {
			return []int{preIndex, i}
		} else {
			maps[v] = i
		}
	}

	return nil
}

// Leetcode 49 字母异位词分组
// 排序法
// 1. 创建一个映射，key为 string（排序后），value为[]string。
// 2. 遍历每个字符串：
// - 将字符串排序, 得到"标准键"
// - 以这个"标准键"去map中查找：
// 3. 遍历map中的所有值，将这些字符串切片组合成一个二维切片作为最终结果
func groupAnagrams(strs []string) [][]string {
	maps := make(map[string][]string)
	for _, str := range strs {
		tmp := []byte(str) // 转换成slice
		//fmt.Println("str->slice", tmp)
		slices.Sort(tmp) // 排序slice
		//fmt.Println("sorted slice", tmp)
		sortedStr := string(tmp) // 转换会str
		//fmt.Println("sorted str", sortedStr)
		maps[sortedStr] = append(maps[sortedStr], str)
		//fmt.Println("")
	}

	// 创建结果二维切片
	result := make([][]string, 0, len(maps))
	// 遍历哈希表，将分组结果添加到结果中
	for _, group := range maps {
		result = append(result, group)
	}

	return result
}

// 辅助函数：对字符串进行排序
func sortString(str string) string {
	chars := strings.Split(str, "") // string -> []string
	sort.Strings(chars)
	return strings.Join(chars, "") // []string -> str
}

// Leetcode 128 最长连续序列
// 1. 去重处理：将所有数字存入哈希集合，自动去除重复元素
// 2. 寻找序列起点：遍历集合，只处理那些是连续序列起点的数字（即 num-1不在集合中）
// 3. 扩展序列：从起点向后查找连续数字，统计序列长度
// 4. 更新最大值：比较并记录最长序列长度
func longestConsecutive(nums []int) int {
	maps := make(map[int]bool)
	for _, num := range nums {
		maps[num] = true
	}
	res := 0
	for x := range maps {
		if maps[x-1] { // 如果x-1也在map表中,那么这个值就不是序列起点
			continue
		}
		// x是序列起点
		y := x + 1
		for maps[y] {
			y++
		}
		res = max(res, y-x)
	}
	return res
}

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println("twoSum", res)

	res1 := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println("groupAnagrams", res1)

	res2 := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Println("longestConsecutive", res2)
}
