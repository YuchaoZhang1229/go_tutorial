package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 创建链表
func createLinkedList(arr []int) *ListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

// 打印整个链表
func printList(head *ListNode) {
	if head == nil {
		fmt.Println("空链表")
		return
	}
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d", p.Val)
		if p.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

// LeeCode 21 合并两个有序列表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	p1 := list1
	p2 := list2
	for p1 != nil && p2 != nil {
		if p1.Val >= p2.Val {
			dummy.Next = p2
			p2 = p2.Next

		} else {
			dummy.Next = p1
			p1 = p1.Next
		}

		dummy = dummy.Next
	}

	if p1 != nil {
		dummy.Next = p1
	}

	if p2 != nil {
		dummy.Next = p2
	}

	return head.Next
}

// Leetcode 86 分解链表
func partition(head *ListNode, x int) *ListNode {

}

func main() {
	list1 := createLinkedList([]int{6, 7, 10})
	list2 := createLinkedList([]int{1, 5, 8})

	merged := mergeTwoLists(list1, list2)
	printList(merged)

}
