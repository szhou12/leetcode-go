package leetcode

// My Solution: two pointers
func romanToInt(s string) int {

	dict := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	l, r := 0, 1
	res := 0
	for r < len(s) {
		curRoman := s[l : r+1]
		if val, ok := dict[curRoman]; ok {
			res += val
			l += 2
			r += 2
		} else {
			res += dict[s[l:r]]
			l += 1
			r += 1
		}
	}

	if l < len(s) {
		res += dict[s[l:r]]
	}

	return res

}

// Another Solution: look at one Roman letter once at a time
var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt2(s string) int {
	if s == "" {
		return 0
	}
	num, lastint, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i] // 从末位的 roman letter 往前看
		num = roman[char]
		if num < lastint { // 两个roman letter 可以组合的情况只有：当前一个的值 < 后一个的值
			total = total - num
		} else {
			total = total + num
		}
		lastint = num
	}
	return total
}
