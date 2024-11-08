package amz

/*
Given an array of length n: segments[][] where segments[i] = [start index, end index, value].
Imagine you have a number line indexing from 1 to +inf. 
segments[i] means each corresponding indexed cells from start to end have the value.
e.g. segment[0] = [1, 4, 2] means cells 1, 2, 3, 4 have value 2.
Now, we have a constraint: any two segments segments[i] and segments[j] cannot intersect.
Given a window size k, find the maximum value from this number line of size k.

Example:
segments = [[1, 4, 2], [6, 6, 5], [7, 7, 7], [7, 9, 1], [9, 10, 1]]
k = 5

The number line will look like this:
index: 1 2 3 4 5 6 7 8 9 10
value: 2 2 2 2 0 5 7 0 1 1
           ---------

The max value of size k = [3-7] = 16
Output: 16
*/

/*
Key Observation: 
1. segments[][] is obviously "range update" -> we can use difference array (diff[] and presum[]) to form what the number line looks like.
2. window size k is "range query" -> we can use prefix sum to efficiently calculate the sum of a range k.
3. Consider the constraint: if a segments[i] intersects with the previous segment, its value should NOT be added. In the given example, [7, 9, 1] intersects with [7, 7, 7]. So, we should skip adding 1 to cell: 7, 8, 9.

Solution:
Step 1: find the max ending index n.
Step 2: Assume segments[][] is sorted. Use a bool array keep[] where keep[i] = true if segments[i] should be added; otherwise, keep[i] = false if segments[i] intersects with segments[i-1] and its value should NOT be added.
Step 3: create length n+1 diff[] and presum[] to form the values on the number line.
Step 4: create length n prefixSum[]. range sum of size k = prefixSum[i+k-1] - prefixSum[i-1].

Time complexity = O(n)
*/
func FindMaxValueSizeK(segments [][]int, k int) int {
	m := len(segments)

	// Step 1
	n := -1
	for i := 0; i < m; i++ {
		n = max(n, segments[i][1])
	}

	// Step 2
	keep := make([]bool, m)
	for i := 0; i < m; i++ {
		if i == 0 {
			keep[i] = true
		} else {
			if segments[i][0] > segments[i-1][1] {
				keep[i] = true
			} else {
				keep[i] = !keep[i-1]
			}
		}
	}

	// Step 3
	diff := make([]int, n+1)
	for i := 0; i < m; i++ {
		if keep[i] {
			diff[segments[i][0]-1] += segments[i][2]
			diff[segments[i][1]] -= segments[i][2]
		}
	}
	presum := make([]int, n+1)
	presum[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		presum[i] = presum[i-1] + diff[i]
	}

	// Step 4
	prefixSum := make([]int, n)
	prefixSum[0] = presum[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + presum[i]
	}

	res := -1
	// (n-1) - s + 1 = k ==> s = n-k
	for s := 0; s <= n-k; s++ {
		if s == 0 {
			res = max(res, prefixSum[s+k-1])
			continue
		}
		// e-s+1 = k ==> e = s+k-1
		res = max(res, prefixSum[s+k-1] - prefixSum[s-1])
	}

	return res

}