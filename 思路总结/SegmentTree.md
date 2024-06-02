# Segment Tree

## Template
[Segment Tree Template](https://github.com/szhou12/leetcode-go/blob/main/template/SegmentTree.go)

## Overview
线段树的维护仅需要使用小区间的值更新大区间。
线段树是平衡二叉树，时间复杂度总是在$O(\log n)$。
线段树的局限性：解决的问题必须要满足“区间加法”才能将大问题划分为子问题来解决。
区间加法：对于[L, R]区间，它的答案可以由[L, M]和[M+1, R]的答案合并求出。
满足的问题：区间求和，区间最大值/最小值 etc.
不满足的问题：区间的众数，区间最长连续问题，最长不下降问题 etc.

线段树解决问题的步骤：
1. 建树
2. 单点修改/区间修改
3. 区间查询
区间修改后的查询会用到<lazy标记>

[建树]
- 一般是开一个数组，以堆的方式存储数据。
对于node i，左孩子下标为2i，右孩子下标为2i+1，父节点下标为i/2。
- 写代码时使用位运算效率更高。
- 线段树的数组一般要开到 4n 才能确保不出现越界访问。
- 具体实现时，能用Class-object就用Class-object，方便快捷。

[单点修改]
- 当修改数列中下标为i的节点数据时，从root向下DFS：如果当前节点的左孩子的区间[L, R]包含了i (i.e. L <= i <= R)， 就访问左孩子；否则，访问右孩子。直到L=R，也就是搜到了只包含这个数据的节点，此时就可以修改它。同时，不要忘了将包含此数据的大区间的节点值更新。
- 仅存在单点修改的区间查询非常简单，不需要处理<lazy标记>。因为在单点修改时，所有非叶子节点所储存的区间信息已经同步更新。
- 如果要查询的区间(目标区间)完全覆盖当前节点代表的区间，直接返回当前区间的值；否则，继续往下搜索。
- 如果查询区间和左孩子有交集，搜索左孩子。
- 如果查询区间和右孩子有交集，搜索右孩子。
- 最后合并处理左右孩子查询到的数据。

[区间修改]
- <lazy标记>的意义：将此区间标记，表示这个区间的值已经更新，但它的子区间尚未更新，更新的信息就是标记里存的值。
- 如果要修改的区间完全覆盖当前节点代表的区间，直接更新这个区间，打上<lazy标记>。
- 如果没有完全覆盖，并且当前区间已存在<lazy标记>，先下传<lazy标记>到子区间，再清除当前区间的<lazy标记>。
- 如果修改区间和左孩子有交集，搜索左孩子。
- 如果修改区间和右孩子有交集，搜索右孩子。
- 最后将当前区间的值更新。

## 题目
:purple_circle: 求所有子数组中每个子数组不同元素数量的平方和: [2916. Subarrays Distinct Element Sum of Squares II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2916-Subarrays-Distinct-Element-Sum-of-Squares-II)

:red_circle: 分治法解House Robber: [3165. Maximum Sum of Subsequence With Non-adjacent Elements](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3165-Maximum-Sum-of-Subsequence-With-Non-adjacent-Elements)
    - House Robber题型的第二种解法: Divide and Conquer (实现上使用 Segment Tree)

## Resources
- [【neko算法课】线段树 数据结构【6期】](https://www.bilibili.com/video/BV1yF411p7Bt/?spm_id_from=333.337.search-card.all.click&vd_source=0c02ef6f6e7a2b0959d7dd28e9e49da4)