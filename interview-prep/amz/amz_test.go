package amz

import (
	"testing"
)

/*
Terminal Commands:

	Run all tests that contains func_name:
	/amz $ go test -run TestMatchString

	Print out result (t.Logf) in console:
	/amz $ go test -v -run TestMatchString

	Run the only test named exactly as func_name:
	/amz $ go test -run ^TestMatchString$

	Run the only test named as TestMatchString from root directory 
		(NOTE: MUST use regex otherwise TestMatchString will be mistakenly treated as a package name):
	/leetcode-go $ go test -run ./interview-prep/amz ^TestMatchString$
*/
func TestMatchString(t *testing.T) {
	text := []string{"code", "coder"}
	pattern := []string{"co*d", "co*er"}
	expected := []bool{false, true}

	result := matchString(text, pattern)

	t.Logf("result: %v", result) // Print out the result

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("matchString(%v, %v) = %v; want %v", text, pattern, result, expected)
		}
	}
}

func TestCatchSmallerGoods(t *testing.T) {
	items := []int{1, 2, 7, 4, 5}
	start := []int{0, 0, 1}
	end := []int{1, 2, 2}
	query := []int{2, 4}

	expected := []int{2, 5}
	result := catchSmallerGoods(items, start, end, query)

	t.Logf("result: %v", result)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("catchSmallerGoods(%v, %v, %v, %v) = %v; want %v", items, start, end, query, result, expected)
		}
	}
}
