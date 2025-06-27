# [460. LFU Cache](https://leetcode.com/problems/lfu-cache/description/)

## Solution idea
### HashMap
1. Use two HashMaps: 
```
nodes: {key: Node{key, value, freq, prev, next}, ...}
lists: {freq: head <--> key1 <--> key2 <--> key3 <--> tail, ...}
```

Time complexity = $O(1)$

Space complexity = $O(c)$ where $c$ is the capacity of the cache

## Resource
[DoublyLinkedList + HashMap: Time O(1) + Clean Code](https://leetcode.com/problems/lfu-cache/solutions/6339090/doublylinkedlist-hashmap-time-o-1-clean-code/)