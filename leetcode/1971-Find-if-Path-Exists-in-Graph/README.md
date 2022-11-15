# [1971. Find if Path Exists in Graph](https://leetcode.com/problems/find-if-path-exists-in-graph/)

## Solution idea

### DFS 暴力解 - 会超时！！！

### Union Find

* 把连通的vertex都添加到一个集合中
* 再判断 `source` 和 `destination` 是否有用一个根节点 (ie. 在同一个集合中)
    * 如果有同一个根节点，说明有路径
    * 否则，没有

Time complexity = $O(n)$ where $n=$ # of edges

[代码随想录-1971. 寻找图中是否存在路径](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1971.%E5%AF%BB%E6%89%BE%E5%9B%BE%E4%B8%AD%E6%98%AF%E5%90%A6%E5%AD%98%E5%9C%A8%E8%B7%AF%E5%BE%84.md)

