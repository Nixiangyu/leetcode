package main

func validPalindrome(s string) bool {
	var isSucessful bool
	for i:= 0; i <len(s) /2; i++ {
		if s[i] != s[len(s) - i - 1] {
			if !isSucessful {
				s = s[:i]+s[i+2:]
				isSucessful = true
				i--
				continue
			}
			return false
		}
	}

	return true
}

func main() {
}
