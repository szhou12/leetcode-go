package leetcode

import "container/heap"

type SORTracker struct {
	left  *MinHeap
	right *MaxHeap
}

func Constructor() SORTracker {
	pq1 := &MinHeap{}
	heap.Init(pq1)
	pq2 := &MaxHeap{}
	heap.Init(pq2)

	return SORTracker{
		left: pq1,
		right: pq2,
	}
}

func (this *SORTracker) Add(name string, score int) {
	cur := &Pair{
		name: name,
		score: score,
	}
	// compare to the left heap's top
	if this.left.Len() > 0 && cur.Greater(this.left.Top().(*Pair)) {
		heap.Push(this.left, cur)
		heap.Push(this.right, heap.Pop(this.left))
	} else {
		heap.Push(this.right, cur)
	}
}

func (this *SORTracker) Get() string {
	heap.Push(this.left, heap.Pop(this.right))
	return this.left.Top().(*Pair).name
}

type Pair struct {
	name  string
	score int
}

func (p1 *Pair) Greater(p2 *Pair) bool {
	if p1.score == p2.score {
		return p1.name < p2.name
	}
	return p1.score > p2.score
}

type MinHeap []*Pair

func (pq MinHeap) Len() int {
	return len(pq)
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq MinHeap) Less(i, j int) bool {
	if pq[i].score == pq[j].score {
		return pq[i].name > pq[j].name
	}
	return pq[i].score < pq[j].score
}

func (pq *MinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(*Pair))
}

func (pq *MinHeap) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}

func (pq MinHeap) Top() interface{} {
	return pq[0]
}

type MaxHeap []*Pair

func (pq MaxHeap) Len() int {
	return len(pq)
}

func (pq MaxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq MaxHeap) Less(i, j int) bool {
	if pq[i].score == pq[j].score {
		return pq[i].name < pq[j].name
	}
	return pq[i].score > pq[j].score
}

func (pq *MaxHeap) Push(x interface{}) {
	*pq = append(*pq, x.(*Pair))
}

func (pq *MaxHeap) Pop() interface{} {
	n := len(*pq)
	temp := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return temp
}

func (pq MaxHeap) Top() interface{} {
	return pq[0]
}

/**
 * Your SORTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(name,score);
 * param_2 := obj.Get();
 */
