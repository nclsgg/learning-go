package _4

func LongestCommonPrefix(words []string) string {
	prefix := []byte(words[0])

	for _, word := range words[1:] {
		if word == "" {
			return ""
		}
		for j, char := range word {
			if j >= len(prefix) {
				break
			}
			if byte(char) != prefix[j] {
				if j == 0 {
					prefix = prefix[:0]
				}
				prefix = prefix[:j]
				continue
			}

			if j == len(word)-1 {
				prefix = prefix[:j+1]
			}
		}
	}

	return string(prefix)
}
