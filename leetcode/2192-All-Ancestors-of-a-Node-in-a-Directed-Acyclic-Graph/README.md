# [2192. All Ancestors of a Node in a Directed Acyclic Graph](https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/description/)

## Solution idea

### Topological Sort
#### 要点总结
1. 本题比较容易想到用 Topological Sort 解
2. 难点：在于每走到一个node，如何记录它的所有祖先？
    * 不要想太多，用最直接的方法：用`HashMap`当作set，直接给每一个节点配一个集合ancestor来记录它的祖先。
    * queue 每次Pop出来一个node `cur`，就把`cur`一系的血脉继承给下一个要到达的node `nei`


#### 代码结构
```
Step 1: Reconstruct adj list representation + Calculage in-degrees
Step 2: Topological Sort - record cur's ancestors for nei when making the next move
Step 3: Convert set to list & Sort list in increasing order
```

Time complexity = $O(E) + O(V+E) + O(n^2 \log n) = O(n^2 \log n)$

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/BFS/2192.All-Ancestors-of-a-Node-in-a-Directed-Acyclic-Graph)