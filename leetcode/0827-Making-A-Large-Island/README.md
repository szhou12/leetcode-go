# [827. Making A Large Island](https://leetcode.com/problems/making-a-large-island/)

## Solution idea

### DFS - 岛屿沉没类

* 总体思路分两步:
    1. 搜索每个岛屿, 并计算每个岛屿的面积
        * 每搜索到一个岛屿, 对其染色 (i.e. 给一个编号/id)
        * 用一个map记录每个岛屿对应的面积 (key=编号, value=面积)

    2. 遍历每一个为0的cell
        * 把0变成1, 连接起上、下、左、右可能出现的邻居岛屿
        * 统计连接后形成的最大面积
        * 遍历所有 0 之后，就可以得出 选一个0变成1 之后的最大面积。

Time complexity = $O(n^2 + n^2) = O(n^2)$

## Resource
[代码随想录-827. 最大人工岛](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0827.%E6%9C%80%E5%A4%A7%E4%BA%BA%E5%B7%A5%E5%B2%9B.md)