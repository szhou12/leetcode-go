# [1005. Maximize Sum Of Array After K Negations](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/description/)

## Solution idea

### Greedy Algorithm
* 贪心的思路，局部最优：让绝对值大的负数变为正数，当前数值达到最大，整体最优：整个数组和达到最大。
* 局部最优可以推出全局最优。

* 步骤:
    * Step 1: 将数组按照绝对值大小**从大到小**排序，注意要**按照绝对值的大小**
    * Step 2: 从前向后遍历，遇到负数将其变为正数，同时K--
    * Step 3: 如果K还大于0，那么反复转变数值最小的元素，将K用完
    * Step 4: 求和

Time complexity = $O(n\log n + n + n) = O(n\log n)$

### Min Heap
* 每次 negate 堆顶元素，再入堆
* 反复直到k用完

### DFS - naive solution
* 尝试所有地方，找到最大和的那一个，会超时！！！