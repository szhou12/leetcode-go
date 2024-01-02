# Linked List 的一些思路总结

## Dummy Head 防止 null pointer
* [19. Remove Nth Node From End of List](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0019-Remove-Nth-Node-From-End-of-List)

## 递归 
* [24. Swap Nodes in Pairs](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0024-Swap-Nodes-in-Pairs)

* [206. Reverse Linked List](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0206-Reverse-Linked-List)

## 双指针同向而行，相隔k步，pace一样
* [19. Remove Nth Node From End of List](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0019-Remove-Nth-Node-From-End-of-List)

## 双指针同向而行, fast指针快走，slow指针慢走
* [141. Linked List Cycle](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0141-Linked-List-Cycle)
    * 判断是否有环

* [142. Linked List Cycle II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0142-Linked-List-Cycle-II)
    * 判断是否有环，并找到环的入口
    * 需要额外的**数学推导**，不容易想到

## 综合题
* [143. Reorder List](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0143-Reorder-List)
    * 双指针倍速走，递归，断开与重连

* [234. Palindrome Linked List](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0234-Palindrome-Linked-List)
    * 双指针倍速走找中点，断开, 迭代法翻转链表

## 不等长两条链
* [160. Intersection of Two Linked Lists](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0160-Intersection-of-Two-Linked-Lists)
    * 同一起跑线法：长的LL先走，走的步数为长度的差值；再同时走
