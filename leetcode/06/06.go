package _6

import (
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListToArray(list *ListNode) []int {
	if list == nil {
		return []int{}
	}

	return append([]int{list.Val}, ListToArray(list.Next)...)
}

func ArrayToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	head := &ListNode{Val: arr[0]}
	current := head

	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	return head
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	p := arr[0]
	var left, right []int

	for _, i := range arr[1:] {
		if i <= p {
			left = append(left, []int{i}...)
		} else {
			right = append(right, []int{i}...)
		}
	}

	return slices.Concat(QuickSort(left), []int{p}, QuickSort(right))
}

func MergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	list1Array := ListToArray(list1)
	list2Array := ListToArray(list2)
	sortedArray := QuickSort(append(list1Array, list2Array...))

	list3 := ArrayToList(sortedArray)

	return list3
}
