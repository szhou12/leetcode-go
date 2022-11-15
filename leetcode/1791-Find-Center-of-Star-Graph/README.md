# [1791. Find Center of Star Graph](https://leetcode.com/problems/find-center-of-star-graph/description/)

## Solution idea

### My Solution
* 中心节点一定出现在每条边

Time complexity = $O(n)$

### 入度
* 统计各个节点的度（无向图所以没有区别入度和出度），如果某个节点的度等于这个图边的数量。 那么这个节点一定是中心节点。

## Resource
[代码随想录-1791.找出星型图的中心节点](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1791.%E6%89%BE%E5%87%BA%E6%98%9F%E5%9E%8B%E5%9B%BE%E7%9A%84%E4%B8%AD%E5%BF%83%E8%8A%82%E7%82%B9.md)