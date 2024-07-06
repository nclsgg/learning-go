package _5

type Bracket struct {
	Match rune
	Type  string
}

func ValidParentheses(s string) bool {
	var stack []rune
	match := map[rune]Bracket{
		'(': Bracket{')', "open"},
		')': Bracket{'(', "close"},
		'{': Bracket{'}', "open"},
		'}': Bracket{'{', "close"},
		'[': Bracket{']', "open"},
		']': Bracket{'[', "close"},
	}

	for _, bracket := range s {
		if match[bracket].Type == "close" {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == match[bracket].Match {
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		}

		stack = append(stack, bracket)
	}

	return len(stack) == 0
}
