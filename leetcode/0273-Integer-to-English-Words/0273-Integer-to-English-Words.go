package leetcode

import "strings"

var GROUPS = []string{"", "Thousand", "Million", "Billion"}
var TENS = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
var LESS_THAN_20 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
	"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	words := ""
	i := 0 // every 3-digit group

	for num > 0 {
		// convert last 3 digits
		if num%1000 != 0 {
			words = convertDigits(num%1000) + GROUPS[i] + " " + words
		}
		num /= 1000
		i++
	}

	return strings.TrimSpace(words)

}

func convertDigits(num int) string {
	if num == 0 {
		return ""
	} else if num < 20 { // two digits and <20
		return LESS_THAN_20[num] + " "
	} else if num < 100 { // two digits and >= 20
		return TENS[num/10] + " " + convertDigits(num%10) // if use LESS_THAN_20[num] here, will miss " " at the end; Thus, use recursive call
	} else { // three digits
		return LESS_THAN_20[num/100] + " Hundred " + convertDigits(num%100)
	}
}