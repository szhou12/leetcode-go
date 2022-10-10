# [234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)

## Solution idea

### Naive Idea
最简单的做法是把遍历一遍，把所有值存在一个数组里，再用双指针相向而行.

### Advanced Idea

1. 快慢指针找到中间节点 Middle Node

2. 从Middle Node处断开，即，Middle Node 成为第二条链的head
    * 翻转第二条链，Iterative Way (用三个指针`pre`, `cur`, `next`做传递)

3. 第一条链从头开始，第二条链从 new head 开始，同时同速度走, 每一步对应的value要相等.
    * 注意，当有奇数个Nodes时，第二条链末位多出一个链，不要紧，能顺利走完第一条链即可.

Time complexity = $O(n)$

## Resource

[代码随想录-234.回文链表](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0234.%E5%9B%9E%E6%96%87%E9%93%BE%E8%A1%A8.md)
