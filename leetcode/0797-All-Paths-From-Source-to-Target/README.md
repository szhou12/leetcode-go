# [797. All Paths From Source to Target](https://leetcode.com/problems/all-paths-from-source-to-target/description/)

## Solution idea

### DFS
* 找出所有可以从`node 0` 到`node n-1`的路径: `[0, ..., n - 1]`
    * 到达`node n-1`就是一条合法路径，不管中间经过多少nodes
* 题目给出的是有向无环图 (DAG)，所以不用担心回到visited node

Time complexity = $O(V+E)$

## Resource
[代码随想录-797.所有可能的路径](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0797.%E6%89%80%E6%9C%89%E5%8F%AF%E8%83%BD%E7%9A%84%E8%B7%AF%E5%BE%84.md)

[图论基础及遍历算法](https://labuladong.github.io/algo/2/22/50/)

