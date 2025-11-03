package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 输入一个数组，转换为一条单链表
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

// 在单链表头部插入新元素
func insertAtHead(head *ListNode, val int) *ListNode {
	// 在单链表头部插入一个新节点 0
	newNode := &ListNode{Val: val}
	newNode.Next = head
	// 现在链表变成了 0 -> 1 -> 2 -> 3 -> 4 -> 5
	return newNode
}

// 在单链表尾部插入新元素
func insertAtTail(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	if head == nil {
		return newNode
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head
}

// 在指定位置插入节点
func insertAtIndex(head *ListNode, index int, val int) *ListNode {
	if index < 0 {
		return head
	}
	if index == 0 {
		return insertAtHead(head, val)
	}

	newNode := &ListNode{Val: val}
	current := head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		return head // 位置超出链表长度，不插入
	}

	newNode.Next = current.Next
	current.Next = newNode
	return head
}

// 删除指定位置的节点
func deleteAtIndex(head *ListNode, index int) *ListNode {
	if head == nil || index < 0 {
		return head
	}
	if index == 0 {
		return head.Next
	}

	current := head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil || current.Next == nil {
		return head // 位置超出链表长度，不删除
	}

	current.Next = current.Next.Next
	return head
}

// 获取链表长度
func getLength(head *ListNode) int {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// 查找节点值
func searchValue(head *ListNode, target int) bool {
	current := head
	for current != nil {
		if current.Val == target {
			return true
		}
		current = current.Next
	}
	return false
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		nextTemp := current.Next
		current.Next = prev
		prev = current
		current = nextTemp
	}
	return prev
}

func main() {
	// 创建一条单链表
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Print("原始链表: ")
	printList(head)

	// 在头部插入节点
	head = insertAtHead(head, 0)
	fmt.Print("头部插入0后: ")
	printList(head)

	// 在尾部插入节点
	head = insertAtTail(head, 6)
	fmt.Print("尾部插入6后: ")
	printList(head)

	// 在指定位置插入节点
	head = insertAtIndex(head, 3, 99)
	fmt.Print("在位置3插入99后: ")
	printList(head)

	// 删除指定位置节点
	head = deleteAtIndex(head, 3)
	fmt.Print("删除位置3的节点后: ")
	printList(head)

	// 反转链表
	head = reverseList(head)
	fmt.Print("反转链表后: ")
	printList(head)

	// 获取链表长度
	fmt.Printf("链表长度: %d\n", getLength(head))

	// 查找节点
	fmt.Printf("是否存在值为3的节点: %t\n", searchValue(head, 3))
	fmt.Printf("是否存在值为10的节点: %t\n", searchValue(head, 10))

}
