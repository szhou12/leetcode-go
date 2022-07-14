package leetcode

// One pass check all failure cases
func canChange(start string, target string) bool {
	j := 0 // pointer for target

	// i = pointer for start
	for i := 0; i < len(start); i++ {
		// if start[i] = '_', keep incrementing i
		if start[i] == '_' {
			continue
		}

		// if target[j] = '_', keep incrementing j
		for j < len(target) && target[j] == '_' {
			j++
		}

		// Check cases that cause two key properties to fail:
		// if j finishes walking but i hasn't, return false
		if j == len(target) {
			return false
		}

		// if i,j both now point to a letter, they must point to the same one L/R
		if start[i] != target[j] {
			return false
		}

		// if i now points to L, then the starting position of L must be to the right of the ending position
		// i.e. i > j
		if start[i] == 'L' && i < j {
			return false
		}

		// if i now points to R, then the starting position of R must be to the left of the ending position
		if start[i] == 'R' && i > j {
			return false
		}

		// passed all failure cases, j can now move on to the next L/R
		j++
	}

	// post-processing
	// if i finishes walking but j hasn't, then if any upcoming letter that j will point to cannot be L/R
	// bc i has no corresponding letter anymore
	for j < len(target) {
		if target[j] != '_' {
			return false
		}
		j++
	}

	return true
}
