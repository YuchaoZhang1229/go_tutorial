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
	// 存放小于 x 的链表的虚拟头结点
	dummy1 := &ListNode{-1, nil}
	// 存放大于等于 x 的链表的虚拟头结点
	dummy2 := &ListNode{-1, nil}
	// p1, p2 指针负责生成结果链表
	p1, p2 := dummy1, dummy2
	// p 负责遍历原链表，类似合并两个有序链表的逻辑
	// 这里是将一个链表分解成两个链表
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		// 不能直接让 p 指针前进，
		// p = p.Next
		// 断开原链表中的每个节点的 next 指针
		temp := p.Next
		p.Next = nil
		p = temp
	}
	// 连接两个链表
	p1.Next = dummy2.Next

	return dummy1.Next
}

func main() {
	//list1 := createLinkedList([]int{6, 7, 10})
	//list2 := createLinkedList([]int{1, 5, 8})
	//merged := mergeTwoLists(list1, list2)
	//printList(merged)

	list3 := createLinkedList([]int{1, 4, 3, 2, 5, 2})
	partition(list3, 3)
	printList(list3)
}
