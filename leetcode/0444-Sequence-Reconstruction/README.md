# [444. Sequence Reconstruction](https://leetcode.ca/2017-02-16-444-Sequence-Reconstruction/)

## Solution idea
### Topological Sort
#### 要点总结
1. When does Topological Sort gives uniquely sorted array?
    * if there is `> 1` nodes at same level (degree == 0), then sorted array is NOT unique
    * 如何判别？`queue` 每时每刻都只会有一个元素，每 pop 1个元素，它只会 enqueue 1个元素

Time complexity = $O(V+E)$

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/BFS/444.Sequence-Reconstruction)