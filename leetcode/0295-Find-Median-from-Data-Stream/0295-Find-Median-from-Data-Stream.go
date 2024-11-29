package leetcode

import "container/heap"

type MedianFinder struct {
	left  *Heap
	right *Heap
}

func Constructor() MedianFinder {
	// maxHeap stores smallest
	pq1 := &Heap{
		lessFunc: func(a, b int) bool { return a > b },
	}
	// minHeap stores largest
	pq2 := &Heap{
		lessFunc: func(a, b int) bool { return a < b },
	}

	return MedianFinder{
		left:  pq1,
		right: pq2,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// if n:n, will become n+1:n from n:n -> n:n+1 -> n+1:n
	if this.left.Len() == this.right.Len() {
		heap.Push(this.right, num)
		heap.Push(this.left, heap.Pop(this.right))
	} else {
		// if n+1:n, will become n+1:n+1 from n+1:n -> n+2:n -> n+1:n+1
		heap.Push(this.left, num)
		heap.Push(this.right, heap.Pop(this.left))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// if n:n, return avg(left.top, right.top)
	if this.left.Len() == this.right.Len() {
		return float64(this.left.Top().(int) + this.right.Top().(int)) / 2
	} else {
		// if n+1:n, return left.top
		return float64(this.left.Top().(int))
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type Heap struct {
	data     []int
	lessFunc func(a, b int) bool
}

func NewHeap(less func(a, b int) bool) *Heap {
	return &Heap{lessFunc: less}
}

func (pq *Heap) Len() int {
	return len(pq.data)
}

func (pq *Heap) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *Heap) Less(i, j int) bool {
	return pq.lessFunc(pq.data[i], pq.data[j])
}

func (pq *Heap) Push(x interface{}) {
	pq.data = append(pq.data, x.(int))
}

func (pq *Heap) Pop() interface{} {
	n := len(pq.data)
	temp := pq.data[n-1]
	pq.data = pq.data[:n-1]
	return temp
}

func (pq *Heap) Top() interface{} {
	return pq.data[0]
}
