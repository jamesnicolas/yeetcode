package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseUntil(head, tail *ListNode) *ListNode {
	if tail == nil {
		return head
	}
	if head != nil && head.Next != nil {
		reversed := head
		toDo := head.Next
		reversed.Next = tail
		tmp := toDo.Next
		for toDo != tail {
			tmp = toDo.Next
			toDo.Next = reversed
			reversed = toDo
			toDo = tmp
		}
		return reversed
	}
	return head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	kthNode := head
	for i := 0; i < k; i++ {
		if kthNode != nil && kthNode.Next != nil {
			kthNode = kthNode.Next
		} else {
			return head
		}
	}
	tmp := head
	head = reverseUntil(head, kthNode)
	tmp.Next = reverseKGroup(kthNode, k)
	return head
}

func cons(i int, n *ListNode) *ListNode {
	return &ListNode{i, n}
}

func main() {
	a := cons(1, cons(2, cons(3, cons(4, cons(5, nil)))))
	b := reverseKGroup(a, 3)
	fmt.Println(b)
	fmt.Println(b.Next)
	fmt.Println(b.Next.Next)
	fmt.Println(b.Next.Next.Next)
	fmt.Println(b.Next.Next.Next.Next)
}
