# [160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/)

## Solution idea

这题难点在于，由于两条链表的长度可能不同，不管怎么设置fast, slow指针的速度 (步长)， 两个指针并不能同时走到公共节点 (相交节点)

解决这个问题的关键是，通过某些方式，让 fast 和 slow 能够同时到达相交节点

**突破口**: 让在较长的LL上的指针先走，走过 两条链长度差 的步数，此时，两个指针就相当于处在在 “同一起跑线” (笨鸟先飞)


Time complexity = $O(\max\{m, n\})$ where $m$ is length of LL_A, $n$ is length of LL_B


## Resource
[代码随想录-面试题 02.07. 链表相交](https://programmercarl.com/%E9%9D%A2%E8%AF%95%E9%A2%9802.07.%E9%93%BE%E8%A1%A8%E7%9B%B8%E4%BA%A4.html#%E6%80%9D%E8%B7%AF)