package leetcode

import "container/heap"

/*
* Implements a priority queue
 */
type wordCount struct {
	word  string
	count int
}

type PQ []*wordCount

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq PQ) Less(i, j int) bool {
	if pq[i].count == pq[j].count {
		// lexicographical order
		return pq[i].word > pq[j].word
	}
	return pq[i].count < pq[j].count
}

func (pq *PQ) Push(x interface{}) {
	temp := x.(*wordCount)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}

/*
* problem solving
 */
func topKFrequent(words []string, k int) []string {
	// step 1: use a map that maps each unique word to its counts
	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}
	// step 2: create a minHeap
	pq := &PQ{}
	heap.Init(pq)
	for w, c := range wordMap {
		heap.Push(pq, &wordCount{w, c})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	// step 3: return k elements from PQ
	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		wc := heap.Pop(pq).(*wordCount)
		res[i] = wc.word
	}
	return res
}
