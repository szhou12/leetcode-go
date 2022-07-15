package leetcode

import (
	"container/heap"
	"sort"
)

// Optimal solution: Pigeonhole Principle
func fillCups_optimized(amount []int) int {
	sort.Slice(amount, func(i, j int) bool {
		return amount[i] > amount[j]
	})

	total := 0
	for _, x := range amount {
		total += x
	}

	if amount[0] >= total/2+1 {
		return amount[0]
	} else {
		// return (total + 1) / 2
		if total%2 == 0 { // if total even
			return total / 2
		} else { // if total odd, then last second can only fill one cup
			return total/2 + 1
		}
	}
}

// Priority Queue
/*
x x x x x
o o o o
y y y
---- (x o)
x x x x
o o o
y y y
---- (x o) (x o)
x x x
o o
y y y
---- (x o) (x o) (x y)
x x
o o
y y
---- (x o) (x o) (x y) (x o)
x
o
y y
---- (x o) (x o) (x y) (x o) (y o)
x

y
---- (x o) (x o) (x y) (x o) (y o)

---- (x o) (x o) (x y) (x o) (y o) (x y)
*/
func fillCups(amount []int) int {
	maxHeap := &PQ{}
	heap.Init(maxHeap)

	for _, x := range amount { // O(3)
		if x > 0 {
			heap.Push(maxHeap, x) // O(log3)
		}
	}

	res := 0
	for maxHeap.Len() >= 2 { // O(nlog3): bc if a number is n, then we will pop and push n times until n decreases to 0
		a := heap.Pop(maxHeap).(int)
		b := heap.Pop(maxHeap).(int)
		res++
		if a-1 > 0 {
			heap.Push(maxHeap, a-1)
		}
		if b-1 > 0 {
			heap.Push(maxHeap, b-1)
		}
	}

	if maxHeap.Len() == 1 { // O(log3)
		c := heap.Pop(maxHeap).(int)
		res += c
	}

	return res
}

type PQ []int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	(*pq) = (*pq)[:n-1]
	return temp
}
