package main

import (
	"fmt"
	"strconv"
)

// -------------------------控制流程-------------------------
// 136. 只出现一次的数字
// 通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素
func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, value := range nums {
		_, ok := m[value]
		if ok {
			m[value] += 1
		} else {
			m[value] = 1
		}
	}
	var res int
	for k, v := range m {
		if v == 1 {
			res = k
		}
	}
	return res
}

// 9. 回文数
func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}
	rev := 0
	for rev < x/10 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev == x || rev == x/10
}

// 转换为字符串
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// -------------------------字符串-------------------------
// 20. 有效括号
// 考察：字符串处理、栈的使用
func isValid(s string) bool {

	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []rune{}
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')', '}', ']':
			top := len(stack) - 1
			if len(stack) == 0 || (stack[top] != m[v]) {
				return false
			}

			if stack[top] == m[v] {
				stack = append(stack[:top], stack[top+1:]...)
			}
		}
	}
	return len(stack) == 0

}

// 14. 最长公共前缀
// 考察：字符串处理、循环嵌套
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// -------------------------基本值类型-------------------------
// 考察：数组操作、进位处理

// -------------------------引用类型：切片-------------------------
// 26. 删除有序数组中的重复项

// 56. 合并区间

// -------------------------基础-------------------------
// 1. 两数之和
// 考察：数组遍历、map使用
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for index, num := range nums {
		value, ok := m[target-num]
		if ok {
			// 判断元素之前是否存在，如果存在，则返回两个元素的索引
			return []int{value, index}
		}
		// 如果元素之前不存在，则将key为元素值，value为索引存入map
		m[num] = index
	}
	return []int{}
}

func main() {
	// 控制流程
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(isValid("()[]{}"))

	// 字符串
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome2(1221))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	// 基本值类型

	// 引用类型：切片

	// 基础
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

}
