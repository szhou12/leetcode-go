package leetcode

import (
	"container/heap"
	"math/big"
)

// Optimal Solution: DP + 基准排序
// ans[i][j] := only keeping i digits, the j-th smallest number's index
func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	m := len(nums)
	n := len(nums[0])
	ans := make([][]int, n+1)

	for k := 0; k <= n; k++ {
		ans[k] = make([]int, m)
	}

	for j := 0; j < m; j++ {
		// when keeping 0 digits, the order of numbers is the order of their indices
		ans[0][j] = j // this gives initial order
	}

	// for each num, i digits to keep
	for i := 1; i <= n; i++ { // O(nm)

		buckets := make([][]int, 10)
		// j-th num: j = index of this num
		for j := 0; j < m; j++ {
			idx := ans[i-1][j]
			bIdx := nums[idx][len(nums[idx])-i]
			buckets[bIdx-'0'] = append(buckets[bIdx-'0'], idx)
		}

		// Update ans[i][j]
		j := 0
		for b := 0; b <= 9; b++ {
			for _, idx := range buckets[b] {
				ans[i][j] = idx
				j++
			}
		}
	}

	res := make([]int, 0)

	for _, q := range queries { //O(|queries|)
		trim, k := q[1], q[0]
		res = append(res, ans[trim][k-1])
	}
	return res
}

/*********************************************************************/

// My Solution: Priority Queue
func smallestTrimmedNumbers_PQ(nums []string, queries [][]int) []int {
	result := make([]int, len(queries))

	for i, q := range queries { // O(m)
		k, trim := q[0], q[1]
		trimmedNums := trimming(nums, trim)        // O(n)
		kthIndex := getKthSmallest(trimmedNums, k) // O(nlogk)
		result[i] = kthIndex
	}

	return result
}

func trimming(nums []string, trim int) []string {
	res := make([]string, len(nums))
	for i, num := range nums {
		n := len(num)
		startIndex := n - trim
		newNum := num[startIndex:]
		res[i] = newNum
	}

	return res
}

func getKthSmallest(trimmedNums []string, k int) int {

	pq := &PQ{}
	heap.Init(pq)

	// O(nlogk)
	for i, num := range trimmedNums {
		// testing case will have very big int (e.g. 89288488870527604910029), strconv.Atoi won't work
		// numInt, _ := strconv.Atoi(num)
		numInt, _ := new(big.Int).SetString(num, 10)
		cur := &TrimmedNum{i, numInt}
		heap.Push(pq, cur)
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	res := heap.Pop(pq).(*TrimmedNum)
	return res.index
}

type TrimmedNum struct {
	index int
	val   *big.Int
}

type PQ []*TrimmedNum

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	if (pq[i].val).Cmp(pq[j].val) == 0 {
		return pq[i].index > pq[j].index
	}
	return (pq[i].val).Cmp(pq[j].val) == 1
}

func (pq *PQ) Push(x interface{}) {
	temp := x.(*TrimmedNum)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return temp
}
