package gq

import (
	"fmt"
	"sort"
)

func main() {
	s1 := ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	s2 := ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	//res := mergeTwoLists(&s1, &s2)
	res := mergeKLists([]*ListNode{&s1, &s2})

	fmt.Println(res.Val)
	fmt.Println(res.Next.Val)
	fmt.Println(res.Next.Next.Val)
	fmt.Println(res.Next.Next.Next.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func mergeKLists1(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	default:
		tmp := lists[0]
		for i := 0; i < len(lists)-1; i++ {
			tmp = mergeTwoLists(tmp, lists[i+1])
		}
		return tmp
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	switch len(lists) {
	case 0:
		return nil
	case 1:
		return lists[0]
	case 2:
		return mergeTwoLists(lists[0], lists[1])
	default:
		var arr []int
		for _, e := range lists {
			for ; e != nil; e = e.Next {
				arr = append(arr, e.Val)
			}
		}
		sort.Ints(arr)
		var res = &ListNode{}
		var tmp *ListNode
		for i, e := range arr {
			if i == 0 {
				res.Val = e
				res.Next = &ListNode{}
				tmp = res.Next
			} else {
				tmp.Val = e
				if i != len(arr)-1 {
					tmp.Next = &ListNode{}
				} else {
					tmp.Next = nil
					break
				}
				tmp = tmp.Next
			}
		}
		return res
	}
}
