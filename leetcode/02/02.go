package _2

func IsPalindrome(x int) bool {
	reverseX := 0
	original := x

	for x > 0 {
		remainder := x % 10
		reverseX = reverseX*10 + remainder
		x /= 10
	}

	return reverseX == original
}
