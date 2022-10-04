package leetcode

func reverseWords(s string) string {

	sArray := []rune(s)

	// Step 1: Remove extra spaces: leading spaces, trailing spaces, extra spaces between words
	sArrayTrim := removeSpaces(sArray, ' ')

	// Step 2a: Reverse all chars all at once
	reverse(&sArrayTrim, 0, len(sArrayTrim)-1)
	// Step 2b: Reverse all chars within each word, one word at a time
	l, r := 0, 0
	for l < len(sArrayTrim) {
		if r == len(sArrayTrim) || sArrayTrim[r] == ' ' {
			reverse(&sArrayTrim, l, r-1)
			l = r + 1
		}
		r++
	}

	return string(sArrayTrim)

}

func removeSpaces(array []rune, target rune) []rune {
	left := 0 // 物理意义：挡板左边都是保留下来的 char
	for right := 0; right < len(array); right++ {
		if array[right] != target {
			array[left], array[right] = array[right], array[left]
			left++
		} else if left > 0 && array[left-1] != target { // allow one target char if it prev is not target
			// 这一个条件允许单词间留有一个space
			// 注意！！！this condition will lead to one trailing whitespace
			array[left], array[right] = array[right], array[left]
			left++
		}
	}

	arrayTrim := array[:left]
	// post-processing: 去除第二条件造成的trailing whitespace
	i := len(arrayTrim) - 1
	for ; i >= 0 && arrayTrim[i] == target; i-- {

	}
	// 退出时，i指向最后一个合法的char
	return arrayTrim[:i+1]

}

func reverse(array *[]rune, l int, r int) {
	for l <= r {
		(*array)[l], (*array)[r] = (*array)[r], (*array)[l]
		l++
		r--
	}
}
