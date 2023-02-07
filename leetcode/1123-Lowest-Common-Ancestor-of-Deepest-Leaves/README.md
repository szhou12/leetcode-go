# [1123. Lowest Common Ancestor of Deepest Leaves](https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/description/)


## Solution idea

### LCA

* 用一个 helper function, return (LCA, depth)
    * return 的物理意义: subtree中是否找到LCA, 走subtree可以到达的最深有多深

* Base case: `root == nil`， return (nil, input depth)

* Recursion call:
    * 问问看: 是否能在左子树中找到LCA, 左子树最深可以到达哪里
    * 问问看: 是否能在右子树中找到LCA, 右子树最深可以到达哪里

* Return 的情况
    1. 如果左右子树返回的深度相同, 那么当前节点就是它们各自最深叶子节点的最近公共祖先。返回 root 以及 子树返回的深度
    2. 如果左右子树返回的深度不同, 那么当前节点不是LCA, 哪个子树走得深返回哪个子树返回的结果 (因为走得深的子树有最深的叶子节点)

Time complexity = $O(h)$ where $h = $ height of tree

## Resource

[Solution 1: Get Subtree Height and LCA - Java](https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/solutions/334577/java-c-python-two-recursive-solution/?orderBy=most_votes)

[【每日一题】LeetCode 1123. Lowest Common Ancestor of Deepest Leaves](https://www.youtube.com/watch?v=DUXvcoEZJqw&ab_channel=HuifengGuan)