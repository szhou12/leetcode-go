# [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

## Solution idea

### Solution - Min Heap
Use `min heap` to store `k` largest element: every time the top element in min heap is smaller than the current element, pop then push.

Time complexity = $O(n\log k)$

### Solution - Binary Search by Value
1. 用二分法猜一个数`mid`。每猜一次，拿`mid`与`nums[]`中每个元素比对，统计`nums[]`中大于等于`mid`的元素个数。如果符合条件的个数 < k个，说明猜大了，`right = mid - 1`；如果符合条件的个数 >= k个，说明`mid`有可能是答案，`left = mid`。
2. CAUTION: 左右边界初始值为`MinInt` 和 `MaxInt`，注意Go中这样的初始值，会导致在计算`mid`时溢出。解决办法可以是: `left, right := math.MinInt32, math.MaxInt32` 或者 `left, right := math.MinInt / 2, math.MaxInt / 2`。

Time complexity = $O(n\log C)$ where $C = 2^{32}$

### Solution - Quick Select

Detailed explanation refers to **Resource** starting at `17:45`. Use the **Rainbow Sort** technique (refer to **Leetcode 0075**)

#### Partition - 3 Pointers
1. Initialize pointers in `nums[left:right]` (inclusive)
```
i := left
t := left
j := right
```
2. Syntax Meaning:
```
[left, i) < pivot
[i, t) == pivot
(j, right] > pivot
```
3. General Pattern:
```
pivot = O
S    S    S    O    O    X    X    X    L    L
left           ^         ^         ^        right
               i         t         j     
```
4. Loop Execution: 
```
主要移动 t:
if nums[t] < pivot:
    nums[t], nums[i] = nums[i], nums[t]
    i++
    t++
else if nums[t] > pivot:
    nums[t], nums[j] = nums[j], nums[t]
    j--
else:
    t++
```
5. Loop End Condition: `t >= j`
6. Loop End Pattern:
```
S    S    S    O    O    O    L    L    L    L
left           ^         ^    ^            right
               i         j    t  
```

#### Find Kth Largest Element
```
     k?             k?             k?
S    S    S    O    O    O    L    L    L    L
-----------    -----------    ----------------
     a              b                 c

a := number of elements < (pivot=O)
b := number of elements == (pivot=O)
c := number of elements > (pivot=O)

if c >= k:
    Find k-th largest in segment [L]
else if b+c >= k:
    k-th largest = O
else:
    Find (k-(b+c))-th largest in segment [S]
```

Time complexity = $O(n)$ on average. Worst case $O(n^2)$

## Resource
[【每日一题】215. Kth Largest Element in an Array, 7/25/2021](https://www.youtube.com/watch?v=dMq9_EkfSEc&t=713s&ab_channel=HuifengGuan)