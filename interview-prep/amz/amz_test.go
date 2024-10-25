package amz

import (
	"testing"
)

/*
Terminal Commands:

	Run all tests that contains func_name:
	/amz $ go test -run func_name

	Print out result (t.Logf) in console:
	/amz $ go test -v -run func_name

	Run the only test named exactly as func_name:
	/amz $ go test -run ^func_name$

	Run the only test named as func_name from root directory 
		(NOTE: MUST use regex otherwise func_name will be mistakenly treated as a package name):
	/leetcode-go $ go test -run ./interview-prep/amz ^func_name$
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
