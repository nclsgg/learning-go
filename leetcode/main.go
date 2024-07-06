package main

import (
	"fmt"
	_1 "leetcode/01"
	_2 "leetcode/02"
	_3 "leetcode/03"
	_4 "leetcode/04"
	_5 "leetcode/05"
	_6 "leetcode/06"
	_7 "leetcode/07"
)

func main() {
	nums := []int{3, 2, 4}
	fmt.Println("Two Sum", _1.TwoSum(nums, 6))
	fmt.Println("isPalindrome", _2.IsPalindrome(-121))
	fmt.Println("Roman To Int", _3.RomanToInt("MCM"))
	words := []string{"flower", "flow", "flight"}
	fmt.Println("Longest Prefix", _4.LongestCommonPrefix(words))
	fmt.Println("Valid Parentheses", _5.ValidParentheses("({}}"))

	list1 := &_6.ListNode{Val: 1}
	list1.Next = &_6.ListNode{Val: 2}
	list1.Next.Next = &_6.ListNode{Val: 5}
	list1.Next.Next.Next = &_6.ListNode{Val: 2}
	list1.Next.Next.Next.Next = &_6.ListNode{Val: 7}
	list1.Next.Next.Next.Next.Next = &_6.ListNode{Val: 9}

	list2 := &_6.ListNode{Val: 4}
	list2.Next = &_6.ListNode{Val: 1}
	list2.Next.Next = &_6.ListNode{Val: 7}

	list3 := &_6.ListNode{Val: 1}
	list3.Next = &_6.ListNode{Val: 4}
	list3.Next.Next = &_6.ListNode{Val: 5}

	list4 := &_6.ListNode{Val: 1}
	list4.Next = &_6.ListNode{Val: 3}
	list4.Next.Next = &_6.ListNode{Val: 4}

	list5 := &_6.ListNode{Val: 2}
	list5.Next = &_6.ListNode{Val: 6}

	fmt.Println("Merge Two Sorted Lists", _6.MergeTwoList(list1, list2))
	fmt.Println("Merge K Sorted Lists", _7.MergeKSortedList([]*_6.ListNode{list3, list4, list5}))
}
