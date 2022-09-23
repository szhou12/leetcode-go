# [19. Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

## Solution idea

### Two pointers: 同向而行

**Key 1**: 两个指针相隔k步时同时往前走，第一个指针走到底时，第二个指针指向的就是倒数第k个

**Key 2**: 加一个dummy head 防止因为删掉的是第一个指针而导致null pointer exception

Time complexity = $O(n)$