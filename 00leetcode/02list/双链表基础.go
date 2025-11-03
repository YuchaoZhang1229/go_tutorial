package main

import "fmt"

type DoublyListNode struct {
	Val        int
	Prev, Next *DoublyListNode
}

func NewDoublyListNode(x int) *DoublyListNode {
	return &DoublyListNode{Val: x}
}

func CreateDoublyLinkedList(arr []int) *DoublyListNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := NewDoublyListNode(arr[0])
	cur := head
	for i := 1; i < len(arr); i++ {
		newNode := NewDoublyListNode(arr[i])
		cur.Next = newNode
		newNode.Prev = cur
		cur = cur.Next
	}
	return head
}

// 在链表头部插入节点
func InsertAtHead(head *DoublyListNode, val int) *DoublyListNode {
	newNode := NewDoublyListNode(val)
	if head == nil {
		return newNode
	}
	newNode.Next = head
	head.Prev = newNode
	return newNode
}

// 在链表尾部插入节点
func InsertAtTail(head *DoublyListNode, val int) *DoublyListNode {
	newNode := NewDoublyListNode(val)
	if head == nil {
		return newNode
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	newNode.Prev = current
	return head
}

// 在指定位置插入节点
func InsertAtIndex(head *DoublyListNode, index int, val int) *DoublyListNode {
	if index < 0 {
		return head
	}
	if index == 0 {
		return InsertAtHead(head, val)
	}

	newNode := NewDoublyListNode(val)
	current := head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	if current == nil {
		return head // 位置超出链表长度，不插入
	}

	newNode.Next = current.Next
	newNode.Prev = current
	if current.Next != nil {
		current.Next.Prev = newNode
	}
	current.Next = newNode
	return head
}

// 删除头部节点
func DeleteAtHead(head *DoublyListNode) *DoublyListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	newHead := head.Next
	newHead.Prev = nil
	return newHead
}

// 删除尾部节点
func DeleteAtTail(head *DoublyListNode) *DoublyListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Prev.Next = nil
	return head
}

// 删除指定值的节点
func DeleteNode(head *DoublyListNode, val int) *DoublyListNode {
	if head == nil {
		return nil
	}

	// 处理头节点匹配的情况
	if head.Val == val {
		if head.Next != nil {
			head.Next.Prev = nil
		}
		return head.Next
	}

	current := head
	for current != nil && current.Val != val {
		current = current.Next
	}

	if current == nil {
		return head // 未找到节点
	}

	// 更新前后节点的指针[2,3](@ref)
	if current.Prev != nil {
		current.Prev.Next = current.Next
	}
	if current.Next != nil {
		current.Next.Prev = current.Prev
	}

	return head
}

// 查找节点
func FindNode(head *DoublyListNode, val int) *DoublyListNode {
	current := head
	for current != nil {
		if current.Val == val {
			return current
		}
		current = current.Next
	}
	return nil
}

// 获取链表长度
func GetLength(head *DoublyListNode) int {
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// 反转双向链表
func ReverseList(head *DoublyListNode) *DoublyListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *DoublyListNode
	current := head

	for current != nil {
		// 交换前后指针
		nextTemp := current.Next
		current.Next = prev
		current.Prev = nextTemp

		prev = current
		current = nextTemp
	}

	return prev
}

// 打印链表（向前）
func PrintForward(head *DoublyListNode) {
	if head == nil {
		fmt.Println("空链表")
		return
	}
	current := head
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Next != nil {
			fmt.Print(" <-> ")
		}
		current = current.Next
	}
	fmt.Println()
}

// 打印链表（向后）
func PrintBackward(tail *DoublyListNode) {
	if tail == nil {
		fmt.Println("空链表")
		return
	}
	current := tail
	for current != nil {
		fmt.Printf("%d", current.Val)
		if current.Prev != nil {
			fmt.Print(" <-> ")
		}
		current = current.Prev
	}
	fmt.Println()
}

func main() {
	// 创建双向链表
	head := CreateDoublyLinkedList([]int{1, 2, 3, 4, 5})

	fmt.Print("原始链表: ")
	PrintForward(head)

	// 在头部插入
	head = InsertAtHead(head, 0)
	fmt.Print("头部插入0后: ")
	PrintForward(head)

	// 在尾部插入
	head = InsertAtTail(head, 6)
	fmt.Print("尾部插入6后: ")
	PrintForward(head)

	// 在指定位置插入
	head = InsertAtIndex(head, 3, 99)
	fmt.Print("在位置3插入99后: ")
	PrintForward(head)

	// 删除节点
	head = DeleteNode(head, 99)
	fmt.Print("删除值为99的节点后: ")
	PrintForward(head)

	// 在头部删除
	head = DeleteAtHead(head)
	fmt.Print("头部删除6后: ")
	PrintForward(head)

	// 在尾部删除
	head = DeleteAtTail(head)
	fmt.Print("尾部删除6后: ")
	PrintForward(head)

	// 获取链表长度
	fmt.Printf("链表长度: %d\n", GetLength(head))

	// 查找节点
	if node := FindNode(head, 3); node != nil {
		fmt.Printf("找到值为3的节点\n")
	}

	// 反转链表
	head = ReverseList(head)
	fmt.Print("反转链表后: ")
	PrintForward(head)

	// 获取尾节点并向后打印
	tail := head
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}
	fmt.Print("从尾向前打印: ")
	PrintBackward(tail)
}
