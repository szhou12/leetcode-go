# [235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)

## Solution idea

### LCA

* LCA 经典题型: 两个input nodes `p, q` 一定存在于树中

* 利用 BST 性质: **node X is LCA of p, q iff X's value is between p and q's values.**
    1. 如果 current node `X` 比 `p, q` 都小, 说明 current node `X` 太小了, 往它的右子树--增大`X`值的方向看看呢？
    2. 如果 current node `X` 比 `p, q` 都大, 说明 current node `X` 太大了, 往它的左子树--缩小`X`值的方向看看呢？
    3. 否则, 看来current node `X` 恰好夹在`p, q`中间, 说明 current node `X`正合适, 就它了