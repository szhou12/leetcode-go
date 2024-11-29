# [2102. Sequentially Ordinal Rank Tracker](https://leetcode.com/problems/sequentially-ordinal-rank-tracker/description/)

## Solution idea

### MinHeap Stores Largest & MaxHeap Stores Smallest
维护两个堆: MinHeap 和 MaxHeap。按照题目定义：更大的元素是首先score更大，如果score相等，name更小。

MinHeap 始终保持前 i-1 个最大的元素，因为是小顶堆，所以出口是第 i-1 大的元素。

MaxHeap 存从第i大开始剩下的所有元素。

执行`Add()`操作时：如果有新元素X进来，我们与MinHeap的出口元素Y比较：
1. 如果 MinHeap不为空 并且 X > Y：说明X应该是前 i-1 大的一员，而 Y 是第i大。所以，我们先把X放进MinHeap，这时MinHeap已有i个元素，出口为第i大，然后把MinHeap的出口元素扔进MaxHeap。
2. 否则：直接把X放进MaxHeap。MaxHeap自动更新出口元素为第i大。

执行`Get()`操作时：我们知道MaxHeap的出口元素是第i大，所以，先把MaxHeap的出口元素扔进MinHeap，然后返回MinHeap的出口元素。这样，MinHeap中就包含前 i 个最大的元素。在下一轮，第 i+1 次操作时，MinHeap就存够了前 i 个最大的元素。

Time complexity = $O(\log n)$

### 一图胜千言
![2102-demo](https://github.com/user-attachments/assets/fd45f785-a40f-465a-a8d1-fc2410e2ac8b)



## Resource
- [[Golang] Two Heaps (Min and Max)](https://leetcode.com/problems/sequentially-ordinal-rank-tracker/solutions/3249914/golang-two-heaps-min-and-max/)
- [【每日一题】LeetCode 2102. Sequentially Ordinal Rank Tracker](https://www.youtube.com/watch?v=a0wE2YP6R2s&t=769s&ab_channel=HuifengGuan)
