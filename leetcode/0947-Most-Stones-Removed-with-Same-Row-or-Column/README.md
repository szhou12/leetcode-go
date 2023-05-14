# [947. Most Stones Removed with Same Row or Column](https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/description/)

## Solution idea
### Union Find
#### 思路总结
1. 用 Union Find 来解决的思路不难想到：同行/同列的元素相消 :arrow_right: 同行/同列的元素属于同一个家族。
2. 难点在于实现：如何预处理可以把同行/同列的元素聚集起来 :arrow_right: 使用两个 `hashmap`，一个只看同列`Y`坐标的元素，另一个只看同行`X`坐标的元素，这样，分享同一个key的元素们就属于同一个家族，应该被 union 起来。
3. 算需要消去的节点个数 = 总个数 - 家族个数。因为每个家族相消最后一定只剩下一个家族首领
4. 因为 节点 是二维的，UnionFind之前需要压缩为 1-D index = `i*N + j` where `N=10000`

Time complexity = $O(n)$ where `n = |stones|`

## Resource
[【每日一题】947. Most Stones Removed with Same Row or Column, 10/04/2019](https://www.youtube.com/watch?v=HAaik49m0q0&ab_channel=HuifengGuan)