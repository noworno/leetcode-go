/*
  最长回文子串（一段连续的字符串）
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var index, length int
	var count int
	for i := 0; i < 2*len(s)-1; i++ {
		if i%2 == 0 {
			count = 2
			for i-count >= 0 && i+count < 2*len(s) && s[(i-count)/2] == s[(i+count)/2] {
				count += 2
			}
			if count == 2 {
				count = 1
			} else {
				count -= 1
			}
		} else {
			count = 0
			for i-1-count >= 0 && i+1+count < 2*len(s) && s[(i-1-count)/2] == s[(i+1+count)/2] {
				count += 2
			}
		}
		if length < count {
			index = i
			length = count
		}
	}

	return s[(index+1-length)/2 : (index+1+length)/2]
}
