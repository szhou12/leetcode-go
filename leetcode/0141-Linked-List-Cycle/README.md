# [141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)

## Solution idea

### Two Pointers

**双指针同向而行, fast指针走两步，slow指针走一步**

* 为什么fast 走两个节点，slow走一个节点，有环的话，一定会在环内相遇呢，而不是永远的错开呢？
    * **fast指针一定先进入环中，如果fast 指针和slow指针相遇的话，一定是在环中相遇，这是毋庸置疑的**
        * 为什么fast指针和slow指针一定会相遇呢？
            * 因为fast是快速走，slow是慢速走. 相当于，在一个无限长的跑道上 (cycle)，fast追击slow的问题. fast走得快，总有一天会追到slow.

Time complexity = $O(n)$

## Resource

[代码随想录-141. 环形链表](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0141.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8.md)
