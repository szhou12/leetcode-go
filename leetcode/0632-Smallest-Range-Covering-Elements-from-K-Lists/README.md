# [632. Smallest Range Covering Elements from K Lists](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/description/)

## Solution idea
### K sorted lists + Greedy + Heap
1. 所求的range `[a, b]`，其中`a, b`一定是自于k个list中的元素。如果不是，会出现浪费，一定有往里收缩的余地。
2. `k sorted lists`切入点：每个list最小的元素放一块看看有没有什么规律？跟其他元素的关系是什么？以它们为subgroup中如何找到符合题意的答案？可否往更大的元素推演转换？
    1. 最小元素放一块儿，放到minHeap里，单看这个subgroup，我们发现，满足题意的range是其中的最小值和最大值，这样，所有元素才可以覆盖到。
    2. 然后，从minHeap弹出最小值，然后怎么推进？记住，range必须覆盖所有list，所以，下一个要补进的元素一定是与弹出的元素所属同一个list。并且，补进的元素是这个list的次小值。此时minHeap里面的所有元素，就是当前全局最小值所对应的range (即恰好覆盖每个list)。再从minHeao中找出最小值和最大值，计算range，更新全局最小range。
    3. 每一时刻，重复找出minHeap中最大最小值算出range，直到某一个list的元素全用完了，不再能覆盖所有list，这时候就可以结束循环。
3. Go在实现上，需要自定义一个`maxVal`来随时记录当前minHeap里最大值是谁。

Time complexity = $O(l_{min}\cdot k\log k)$ where $l$ is the shortest list of all, $k$ is the number of lists.

## Resource
[【每日一题】632. Smallest Range Covering Elements from K Lists, 12/3/2020](https://www.youtube.com/watch?v=ejVD92bJe34&ab_channel=HuifengGuan)