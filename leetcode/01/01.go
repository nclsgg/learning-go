package _1

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for index, num := range nums {
		if idx, value := m[target-num]; value {
			return []int{idx, index}
		}

		m[num] = index
	}

	return nil
}
