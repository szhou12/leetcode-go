# [2973. Find Number of Coins to Place in Tree Nodes](https://leetcode.com/problems/find-number-of-coins-to-place-in-tree-nodes/description/)

## Solution idea
### DFS + Sorting
1. Tree 的题型，容易想到用 Recursion/DFS 来解决。
2. 难点在于如何让子树汇报自己的 max product?
3. 解决方法:
    1. 数学规律: 怎样从一堆可正可负的数中挑三个元素得到最大乘积？只可能出现在两端：a. 三个最大的正数相乘; b. 两个最小的负数再乘以最大的正数。
        * `arr[n-1] * arr[n-2] * arr[n-3]` or `arr[0] * arr[1] * arr[n-1]` (`arr` in ascending order)
        * 由此，我们可以发现：每次我们只需要五个元素。那么，在子树汇报时，就汇报这五个元素。
        * 当前节点的操作：收集每个子树的五个元素，加上自身，排序，再挑出符合条件的五个元素，返回给父节点。
    2. 每次汇报给父节点时，如何动态的记录子树的五个元素？
        * 对于要同时**跟踪多个最大最小元素**的情况，首先想到可以使用 **Heap**! 三元素 Max Heap 记录第一大、第二大、第三大的元素；二元素 Min Heap 记录最小的两个元素。
        * Heap 的本质上是起到一个排序的作用，所以也想到每一层使用 `sort` 来代替使用额外的数据结构。
4. 如何记录以当前节点为root的子树的对应五个元素？
    * 原则是这五个元素不能重叠。所以，我们要看一个总数 n = 所有子树汇报的元素个数 + 自己
    * Case 1: if n >= 1, 可以收集最小的元素 `arr[0]`
    * Case 2: if n >= 2, 才可以收集第二小的元素 `arr[1]`
    * Case 3: if n >= 5, 才可以收集第三大的元素 `arr[n-3]`。这样，第三大的元素才不会和第一小、第二小的元素重叠
    * Case 4: if n >= 4, 才可以收集第四大的元素 `arr[n-2]`
    * Case 5: if n >= 3, 才可以收集第五大的元素 `arr[n-1]`

Time complexity = $O(n * n\log n)$. DFS traverses all nodes, and for each node, we have a sorting operation.

## Resource
[【每日一题】LeetCode 2973. Find Number of Coins to Place in Tree Nodes](https://www.youtube.com/watch?v=Omfwj5wARUA&ab_channel=HuifengGuan)