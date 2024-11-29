# [295. Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/description/)

## Solution idea
### MinHeap Stores Largest & MaxHeap Stores Smallest
维护两个堆: left (MaxHeap) 和 right (MinHeap)。

left (MaxHeap) 始终保持 前一半最小的元素，因为是大顶堆，所以出口是第 n/2 小的元素。人为规定：如果奇数个元素，调整left比right多存一个元素。i.e., `left:right = n:n or n+1:n`

right (MinHeap) 存后一半较大的所有元素。

执行`Add()`操作时：我们看left和right的size：
1. 如果`n:n`：说明已有偶数个元素，加进来新元素后将成为奇数个，需调整为`n+1:n`：流程为`n:n -> n:n+1 -> n+1:n`。先把新元素放进right"涮"一下，然后把Mright的出口元素放进left。
2. 否则：说明已有奇数个元素，加进来新元素后将成为偶数个，根据规定，left多一个，需调整为`n+1:n+1`：流程为`n+1:n -> n+2:n -> n+1:n+1`。先把新元素放进left"涮"一下，然后把left的出口元素放进right。

执行`Get()`操作时：如果left和right刚好相等，那么两者的出口元素取平均就是中位数。否则，根据规定，left多一个元素，所以left的出口元素就是中位数。

Time complexity = $O(\log n)$

## Resource
- [Clean 2 Heaps Go Solution](https://leetcode.com/problems/find-median-from-data-stream/solutions/1331923/clean-2-heaps-go-solution/)
