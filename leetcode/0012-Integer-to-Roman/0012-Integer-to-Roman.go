package leetcode

func intToRoman(num int) string {
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	M := []string{"", "M", "MM", "MMM"}
	thousands := M[num/1000] // cut off last 3 digits
	hundreds := C[(num%1000)/100] // get last 3 digits, then cut off last 2 digits
	tens := X[(num%100)/10] // get last 2 digits, then cut off last digit
	ones := I[num%10] // get last digit

	return thousands + hundreds + tens + ones
}