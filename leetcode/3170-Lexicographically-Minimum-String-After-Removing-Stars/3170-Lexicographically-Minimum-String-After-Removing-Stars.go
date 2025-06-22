package leetcode

import (
	"container/heap"
	"strings"
)

// minHeap: store characters seen so far. Top = smallest character
// map: {char : [i1, i2, ...]}
// mark[]: []bool. mark[i] = true if s[i] should be removed
func clearStars(s string) string {
	n := len(s)
	mark := make([]bool, n)
	dict := make(map[int][]int)
	minHeap := &PQ{}
	heap.Init(minHeap)

	for i, c := range s {
		if c != '*' {
			charIdx := int(c - 'a')
			heap.Push(minHeap, charIdx)
			if _, ok := dict[charIdx]; !ok {
				dict[charIdx] = []int{i}
			} else {
				dict[charIdx] = append(dict[charIdx], i)
			}
		} else {
			mark[i] = true
			if minHeap.Len() > 0 {
				cur := heap.Pop(minHeap).(int)
				mark[dict[cur][len(dict[cur])-1]] = true
				dict[cur] = dict[cur][:len(dict[cur])-1]
			}

		}
	}

	var sb strings.Builder
	for i, removed := range mark {
		if !removed {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()

}

type PQ []int

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
