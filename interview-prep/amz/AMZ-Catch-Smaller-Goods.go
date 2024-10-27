package amz

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
items = [1, 2, 7, 4, 5], start = [0, 0, 1], end = [1, 2, 2], query = [2, 4]
	order 0: start=0, end=1, items[0:1] = [1, 2]
	order 1: start=0, end=2, items[0:2] = [1, 2, 7]
	order 2: start=1, end=2, items[1:2] = [2, 7]
query 0: keep items from all 3 orders whose value < 2: there are two 1's. Thus, the answer is 2
query 1: keep items from all 3 orders whose value < 4: there are two 1's and three 2's. Thus, the answer is 5.
Output: [2, 5]
*/

func catchSmallerGoods(items []int, start []int, end []int, query []int) []int {
	
}