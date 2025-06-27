# [146. LRU Cache](https://leetcode.com/problems/lru-cache/description/)

## Solution idea
### Hashmap + Doubly Linked-List
1. Use a dummy head and dummy tail to simplify the node insertion and removal
2. Use a hashmap to store the key-node pairs
3. Use a doubly linked list to store the nodes in the order of their recent usage
4. When a node is used, remove it from the current position and add it to the end of the list
5. When the cache is full, remove the least recently used node from the head of the list

Time complexity = $O(1)$ for both get and put operations
Space complexity = $O(capacity)$ for the cache

