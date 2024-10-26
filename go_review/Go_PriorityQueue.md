# [Priority Queue](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Max/Min Heap](#maxmin-heap)

## Max/Min Heap
[Is there a more generic way to implement 2 kinds of Heap (Max and Min) in Go Lang](https://stackoverflow.com/questions/23580285/is-there-a-more-generic-way-to-implement-2-kinds-of-heap-max-and-min-in-go-lan)
```go
/***** Min Heap *****/
type PQ []int

func (pq PQ) Len() int {
    return len(pq)
}

func (pq PQ) Less(i, j int) bool {
    return pq[i] < pq[j]
} // Min Heap

func (pq PQ) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
func (pq *PQ) Push(x interface{}) {
    // NOTE: 这里为什么要手动解引用 *pq？ 这是因为append()只作用于slice而不是slice 的 pointer; 同时, append()也对slice进行了长度改变
    *pq = append(*pq, x.(int))
}

func (pq *PQ) Pop() interface{} {
    // NOTE: 这里为什么要手动解引用 *pq？ 这是因为len()和slice indexing只作用于slice而不是slice pointer; 同时, *pq = (*pq)[:n-1]也对slice进行了长度改变
    n := len(*pq)
    temp := (*pq)[n-1]
    *pq = (*pq)[:n-1]
    return temp
}

func (pq PQ) Peek() (int, bool) {
    if len(pq) == 0 {
        return 0, false
    }
    return pq[0], true
}

minHeap := &PQ{}
heap.Init(minHeap)

heap.Push(minHeap, x)
x := heap.Pop(minHeap).(int)
```

```go
/***** Max Heap *****/
type PQ []int

func (pq PQ) Len() int {
    return len(pq)
}

func (pq PQ) Less(i, j int) bool {
    return pq[i] > pq[j]
} // Max Heap

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

func (pq PQ) Peek() (int, bool) {
    if len(pq) == 0 {
        return 0, false
    }
    return pq[0], true
}

maxHeap := &PQ{}
heap.Init(maxHeap)

heap.Push(maxHeap, x)
x := heap.Pop(maxHeap).(int)
```