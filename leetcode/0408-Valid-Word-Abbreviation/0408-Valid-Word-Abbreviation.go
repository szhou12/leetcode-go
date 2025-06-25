package leetcode

func validWordAbbreviation(word string, abbr string) bool {
	i, j := 0, 0

	for i < len(word) && j < len(abbr) {
		if '0' <= abbr[j] && abbr[j] <= '9' {
			// if encounters a number, j always starts at the leftmost digit
			// can't allow leading zero
			if abbr[j] == '0' {
				return false
			}

			num := 0
			for j < len(abbr) && ('0' <= abbr[j] && abbr[j] <= '9') {
				num = num * 10 + int(abbr[j] - '0')
				j++
			}
			// make i skip by num step
			i += num
		} else { // both i, j point to an English letter
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++

		}
	}

	return i == len(word) && j == len(abbr)
}