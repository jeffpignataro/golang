package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateListNode(a []int) *ListNode {
	var l *ListNode

	for i := len(a) - 1; i >= 0; i-- {
		l = &ListNode{Val: a[i], Next: l}
	}

	return l
}

func main() {
	a1, a2 := []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}

	l1 := generateListNode(a1)
	l2 := generateListNode(a2)

	l := addTwoNumbers(l1, l2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		l3, l4 *ListNode
		carry  int
	)

	for l1 != nil || l2 != nil {
		var sum int

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry

		carry = sum / 10
		sum = sum % 10

		if l3 == nil {
			l3 = &ListNode{Val: sum}
			l4 = l3
		} else {
			l3.Next = &ListNode{Val: sum}
			l3 = l3.Next
		}
	}

	if carry > 0 {
		l3.Next = &ListNode{Val: carry}
	}

	return l4
}
