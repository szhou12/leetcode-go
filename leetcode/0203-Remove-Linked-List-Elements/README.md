# [203. Remove Linked List Elements](https://leetcode.com/problems/remove-linked-list-elements/)

## Solution idea

**Linked List Traversal** + **Dummy Head**

Use **Dummy Head** because the head node may be removed from the list

Traverse input linked list. 

If the current node's value is NOT the target, connect to dummy head's list. Else, skip current node.

**Detail**: for-loop 停在最后一个node, 也就是最后一个node不在loop中处理，让其在post-processing中被处理. 

原因：如果最后一个node恰好是应该被removed，loop在前一次循环中链接了倒数第二个node，而倒数第二个node在input list中链接着最后一个node，也就是说，output list中现在处于链接着最后一个node的状况，但最后一个node应该被剔除。

Time complexity = $O(n)$