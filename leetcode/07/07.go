package _7

import (
	_6 "leetcode/06"
)

func MergeKSortedList(lists []*_6.ListNode) *_6.ListNode {
	if len(lists) == 0 {
		return nil
	}

	var mergedLists []int

	for _, list := range lists {
		if list == nil {
			continue
		}

		listArray := _6.ListToArray(list)
		mergedLists = append(mergedLists, listArray...)
	}

	return _6.ArrayToList(_6.QuickSort(mergedLists))
}
