# Linked List 的一些思路总结

## Dummy Head 防止 null pointer
* [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/),

## 递归 
* [24. Swap Nodes in Pairs](https://leetcode.com/problems/swap-nodes-in-pairs/)

* [206. Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

## 双指针同向而行，相隔k步，pace一样
* [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## 双指针同向而行, fast指针快走，slow指针慢走
* [141. Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/)
    * 判断是否有环

* [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)
    * 判断是否有环，并找到环的入口
    * 需要额外的**数学推导**，不容易想到

## 综合题
* [143. Reorder List](https://leetcode.com/problems/reorder-list/)
    * 双指针倍速走，递归，断开与重连

* [234. Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/)
    * 双指针倍速走找中点，断开, 迭代法翻转链表

## 不等长两条LL
* [160. Intersection of Two Linked Lists](https://leetcode.com/problems/intersection-of-two-linked-lists/)
    * 同一起跑线法：长的LL先走，走的步数为长度的差值；再同时走