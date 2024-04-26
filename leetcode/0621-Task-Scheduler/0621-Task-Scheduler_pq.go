package leetcode

import "container/heap"

// Solution: MaxHeap{freq, task}
func leastInterval_maxheap(tasks []byte, n int) int {
	n++ // n+1 means no repeated tasks in consecutive n+1 slots

	// Use map to count freq of each task
	count := make(map[byte]int)
	for _, task := range tasks {
		count[task]++
	}

	maxHeap := &PQ{}
	heap.Init(maxHeap)
	for task, freq := range count {
		heap.Push(maxHeap, Pair{freq: freq, task: task})
	}

	res := 0

	// prioritize arranging top k distinct tasks
	for maxHeap.Len() > 0 {
		k := min(n, maxHeap.Len()) // maxHeap may not have enought n elements
		queue := make([]Pair, 0)   // temporarily store popped tasks

		for i := 0; i < k; i++ {
			cur := heap.Pop(maxHeap).(Pair)
			cur.freq--
			if cur.freq > 0 {
				queue = append(queue, cur)
			}
		}

		if len(queue) > 0 { 
			// n may include idle
			// still have tasks to be arranged in next round
			// thus, this round needs to be filled up with n
			res += n
		} else { // 最后一次安排任务不再需要idle补空缺
			res += k
		}

		for _, e := range queue {
			heap.Push(maxHeap, e)
		}
	}

	return res
}

type Pair struct {
	freq int
	task byte
}

type PQ []Pair // {Pair{freq, task}, ...}

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].freq > pq[j].freq
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Pair))
}

func (pq *PQ) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}
