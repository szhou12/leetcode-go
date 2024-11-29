# [Priority Queue](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Standard Max/Min Heap](#standard-maxmin-heap)
* [Head And Tail Elements In Priority Queue](#head-and-tail-elements-in-priority-queue)
* [Custom Heap - Two Heaps At The Same Time](#custom-heap---two-heaps-at-the-same-time)
* [When Need heap.Init()?](#when-need-heapinit)
* [Pointer Receiver vs. Value Receiver](#pointer-receiver-vs-value-receiver)

## Standard Max/Min Heap
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

/** 
下面给出Top()的两种实现方式。
注：写法一在使用时，需要提前Check pq是否为空
**/
func (pq PQ) Top() interface{} {
    return pq[0]
}
func (pq PQ) Top() (int, bool) {
    if len(pq) == 0 {
        return 0, false
    }
    return pq[0], true
}


// Use this in code as follows:
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

/** 
下面给出Top()的两种实现方式。
注：写法一在使用时，需要提前Check pq是否为空
**/
func (pq PQ) Top() interface{} {
    return pq[0]
}
func (pq PQ) Top() (int, bool) {
    if len(pq) == 0 {
        return 0, false
    }
    return pq[0], true
}


// How To Use:
maxHeap := &PQ{}
heap.Init(maxHeap)

heap.Push(maxHeap, x)
x := heap.Pop(maxHeap).(int)

y := maxHeap.Top().(int)
```

## Head And Tail Elements In Priority Queue
In Go implementation of priority queue, the head element `pq[0]` is always the **highest-priority element**.

But why in `Pop()`, we are returning the last element `pq[n-1]`?

This is for the efficiency because if popping the head element, all the rest of elements will be shifted which requires $O(n)$.

Thus, when you call `Pop()` using the `container/heap` package, `Pop()` internally does the following:
1. Swap the Root `pq[0]` with the Last Element `pq[n-1]`.
2. Reheapify the Slice `pq[:n-1]`.

## Custom Heap - Two Heaps At The Same Time
A simplified way to define MaxHeap and MinHeap at the same time:
```go
type Heap struct {
	data     []int
	lessFunc func(a, b int) bool
}

// Can be omitted
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

// How To Use:
// 1. not to use constructor NewHeap()
maxHeap := &Heap{lessFunc: func(a, b int) bool { return a > b }}
minHeap := &Heap{lessFunc: func(a, b int) bool { return a < b }}

// 2. use constructor NewHeap()
maxHeap := NewHeap(func(a, b int) bool { return a > b })
minHeap := NewHeap(func(a, b int) bool { return a < b })
```

## When Need heap.Init()?
**What does heap.Init() do?** 
- `heap.Init()` is used to heapify a slice.

Therefore,
1. If start with an **empty slice**, there’s NO need to call `heap.Init` because `heap.Push` and `heap.Pop` will maintain the heap property from the very first operation.
2. If slice is already **pre-filled with elements**, you MUST call `heap.Init` to re-heapify the slice so that it satisfies the heap property.

Example Where `heap.Init` is NO Need:
```go
pq := &MinHeap{}
// pq starts empty. So you can directly call heap.Push and heap.Pop

heap.Push(pq, 3)
heap.Pop(pq)
```

Example Where `heap.Init` is MUST:
```go
pq := &MinHeap{
    &Pair{name: "a", score: 3},
    &Pair{name: "b", score: 1},
    &Pair{name: "c", score: 2},
}
// This slice is not guaranteed to satisfy the heap property.
heap.Init(pq) // Rearranges elements into a valid heap
```

## Pointer Receiver vs. Value Receiver
In Go, the decision to use a pointer receiver (`(pq *Heap)`) versus a value receiver (`(pq Heap)`) for methods depends on whether the method **modifies the receiver’s state** or simply **reads from it**.
### General Rule
1. Use `(pq *Heap)` if:
    - The method modifies the receiver (`Push()`, `Pop()`).
    - You want to avoid copying the receiver (e.g., when the receiver contains a large slice or struct).
2. Use `(pq Heap)` if:
    - The method ONLY reads from the receiver and does not modify it (`Len()`, `Less()`, `Swap()`, `Top()`).
    - The receiver is small (e.g., a slice or small struct).
### Analysis of Custom Heap Implementation
Given that:
```go
type Heap struct {
	data     []int
	lessFunc func(a, b int) bool
}
```
1. `Len()`:
    - **No modification**: `Len()` simply returns the length of `pq.data`.
    - Can use either a value receiver `(pq Heap)` or a pointer receiver `(pq *Heap)`. However, using a pointer receiver is common for consistency.
2. `Swap()`:
    - **No modification to the heap itself**: `Swap()` only swaps elements within the data slice but does not change the slice's structure or pointer.
    - Can use either a value receiver `(pq Heap)` or a pointer receiver `(pq *Heap)`. However, using a pointer receiver is common for consistency.
3. `Less()`:
    - **No modification**: `Less()` is just a comparison function.
    - Can use either a value receiver `(pq Heap)` or a pointer receiver `(pq *Heap)`. However, using a pointer receiver is common for consistency.
4. `Top()`:
    - **No modification**: `Top()` only reads the top element without modifying the heap.
    - Can use either a value receiver `(pq Heap)` or a pointer receiver `(pq *Heap)`. However, using a pointer receiver is common for consistency.
5. `Push()`:
    - **Modifies the heap**: `Push()` appends an element to data, which modifies the slice's underlying structure. A pointer receiver is necessary.
6. `Pop()`:
    - **Modifies the heap**: `Pop()` removes an element from data, changing the slice. A pointer receiver is necessary.
### Best Practice
For custom heap, it's generally better to **use `(pq *Heap)` for all methods** for:
- **Consistency**: All methods share the same receiver type.
- **Efficiency**: Avoid copying large slices or structs when accessing data.