package _3

func RomanToInt(s string) int {
	sum := 0

	rim := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, v := range s {
		sum += rim[string(v)]
		if i != 0 {
			if rim[string(s[i-1])] < rim[string(v)] {
				sum -= 2 * rim[string(s[i-1])]
			}
		}
	}

	return sum
}
