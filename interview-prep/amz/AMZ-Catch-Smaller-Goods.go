package amz

import "sort"

/*
Given:
1. an array items[] of length n. items[i] := value of item i
2. an array start[] of length m. start[j] := the start indes of order j
3. an array end[] of length m. end[j] := the end index of order j
	For an order j, we pick items from items[start[j] : end[j]] (inclusive)
4. an array query[] of length q. query[k] := we keep items from all m orders whose value < query[k]
Find:
for each query k, the count of items kept.

Example:
items = [1, 2, 7, 3, 15], start = [0, 0, 1], end = [1, 2, 2], query = [2, 4]
	order 0: start=0, end=1, items[0:1] = [1, 2]
	order 1: start=0, end=2, items[0:2] = [1, 2, 7]
	order 2: start=1, end=2, items[1:2] = [2, 7]
query 0: keep items from all 3 orders whose value < 2: there are two 1's. Thus, the answer is 2
query 1: keep items from all 3 orders whose value < 4: there are two 1's and three 2's. Thus, the answer is 5.
Output: [2, 5]
*/

func catchSmallerGoods(items []int, start []int, end []int, query []int) []int {
	n := len(items)
	m := len(start)

	// Step 1: count the frequency of each item index that is being included in orders' range
	// Use difference array and its prefix sum to count the freq
	// freq[] now is diff array
	freq := make([]int, n+1) // +1 to avoid out of bound
	for i := 0; i < m; i++ { // O(m)
		freq[start[i]]++
		freq[end[i]+1]--
	}
	// now, convert freq[] to prefix sum of diff array
	for i := 1; i < n; i++ { // O(n)
		freq[i] += freq[i-1]
	}

	// Step 2: Map the item index to item value.
	// for each unique item value, aggregate the freq
	totalCounts := make(map[int]int) // {item value: total counts for being included}
	for i := 0; i < n; i++ { // O(n)
		// skip 0 bc it means this item never been included
		if freq[i] > 0 {
			totalCounts[items[i]] += freq[i]
		}
	}

	// Step 3: Get all keys (item values) from totalCounts{} and sort in ascending order
	includedItems := make([]int, 0)
	for itemVal, _ := range totalCounts { // O(n)
		includedItems = append(includedItems, itemVal)
	}
	sort.Ints(includedItems)
	// Following the order of sorted includedItems, construct prefix sum array that accumulates counts as of item value = includedItems[i]
	prefixSum := make([]int, len(includedItems))
	for i, itemVal := range includedItems { // O(n)
		if i == 0 {
			prefixSum[i] = totalCounts[itemVal]
		} else {
			prefixSum[i] = prefixSum[i-1] + totalCounts[itemVal]
		}
	}

	// Step 4: for each query q (threshold item value), use Binary Search to find the last includedItems index whose item value < q
	res := make([]int, 0)
	for _, q := range query { // O(q*log(n))
		// first index whose item value >= q
		idx := sort.Search(len(includedItems), func(i int) bool {return includedItems[i] >= q})
		if idx == 0 { // no item value < q
			res = append(res, 0)
		} else {
			res = append(res, prefixSum[idx-1])
		}
	}

	return res
}
