package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	a1, a2 := []int{}, []int{}
	fmt.Println(getNext(l1))

	a1 = append(a1, getNext(l1).Val)
	a2 = append(a2, getNext(l2).Val)

	// v1, v2 := Calculate(l1.Val, l2.Val)
	// r.Val = v1

	fmt.Println(a1)
	fmt.Println(a2)
	return r
}

func getNext(l *ListNode) *ListNode {
	if l != nil {
		return l.Next
	}
	return nil
}

func Calculate(i1, i2 int) (int, int) {
	r := i1 + i2
	if r > 9 {
		return r - 10, 1
	}
	return r, 0
}
