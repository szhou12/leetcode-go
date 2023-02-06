# [2509. Cycle Length Queries in a Tree](https://leetcode.com/problems/cycle-length-queries-in-a-tree/description/)

## Solution idea

### LCA

* 连接树上任意两个nodes得到的cycle，寻找这个cycle的本质就是找到这两个nodes的 LCA

* 由于题目给的是完全二叉树，任意一个node X的parent可以容易回溯. 即, `parent(X) = X/2`

* 如何在不利用recursion的情况下寻找node a, b 的LCA?
    * a, b 谁的值大，说明那个node较深，优先回溯其parent.
    * 直到 parent(a) == parent(b)

* 一个node每回溯一次，就相当于找到cycle的一条边，所以 `count+1`

* 找到LCA后，`count+1` 表示连接 node a, b的那条边，这样，就得到了cycle的长度

Time complexity = $O(m*h)$ where $m = $ length of queries, $h = $ height of tree because LCA of any node is at most at root node, which is `height` farther. 

## Resource
[【每日一题】LeetCode 2509. Cycle Length Queries in a Tree](https://www.youtube.com/watch?v=PMKr2PKmyUY&ab_channel=HuifengGuan)