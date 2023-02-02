package main

func main() {
	println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	res := ""
	lens := len(s)
	for i := 0; i < lens; i++ {
		left := palindrome(s, i, i)
		right := palindrome(s, i, i+1)
		leftLens := len(left)
		rightLens := len(right)
		resLens := len(res)
		if leftLens >= rightLens {
			if resLens < leftLens {
				res = left
			}
		} else {
			if resLens < rightLens {
				res = right
			}
		}
	}
	return res
}

func palindrome(s string, left, right int) string {
	lens := len(s)
	for left >= 0 && right < lens && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
