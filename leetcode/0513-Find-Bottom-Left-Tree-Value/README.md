# [513. Find Bottom Left Tree Value](https://leetcode.com/problems/find-bottom-left-tree-value/)

## Solution idea

### BFS - Level-order traversal
* 层序遍历解这道题非常直观
* 只需要每一层第一个元素记录到一个全局变量，loop结束后，这个全局变量记录的就是最底一层最左边的节点
* 用BFS处理这道题也没有Edge Case需要处理


### DFS
* 深度搜索也可以解这道题, 需要一个变量记录树的深度 
* 主要是掌握 Base Case 的条件:
    1. 要到达叶子节点
    2. 同时 curHeight > maxHeight (保证当前层是第一次递归到达，也就保证此时当前节点一定是最底一层最左边的节点)

Time complexity = $O(n)$

## Resource

[代码随想录-513.找树左下角的值](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0513.%E6%89%BE%E6%A0%91%E5%B7%A6%E4%B8%8B%E8%A7%92%E7%9A%84%E5%80%BC.md)

[halfrost/LeetCode-Go](https://github.com/halfrost/LeetCode-Go/blob/master/leetcode/0513.Find-Bottom-Left-Tree-Value/513.%20Find%20Bottom%20Left%20Tree%20Value.go)
