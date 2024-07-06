package main

import (
	"fmt"
	_1 "leetcode/01"
	_2 "leetcode/02"
	_3 "leetcode/03"
)

func main() {
	nums := []int{3, 2, 4}
	fmt.Println("Two Sum", _1.TwoSum(nums, 6))
	fmt.Println("isPalindrome", _2.IsPalindrome(-121))
	fmt.Println("Roman To Int", _3.RomanToInt("MCM"))
}
