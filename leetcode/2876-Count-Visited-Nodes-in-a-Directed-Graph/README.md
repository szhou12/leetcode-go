# [2876. Count Visited Nodes in a Directed Graph](https://leetcode.com/problems/count-visited-nodes-in-a-directed-graph/description/)

## Solution idea
### Topological Sort + Find Cycle Length + DFS
1. Need to notice that out-degree of each node = 1 by the description that `edges[i]` has ONLY one value.
2. With out-degree = 1, starting from any node, there are only two endings: 1. dead-end 2. go into a cycle. Situation 1 is not possible in this problem bc if dead-end, then the last visited node will have no out-edges (out-degree == 0). Contradictory.
3. Therefore, there definitely will be at least one cycle in the graph (there might be several CCs) -> "百川汇大海"
4. 解法: 1. Topological sort to remove “百川” (毛刺); 2. calculate cycle length from those nodes whose in-degree != 0; 3. DFS to calculate length of “百川” (毛刺)

Time complexity = $O(n)$ (Topological Sort) + $O(n)$ (Find Cycle Length) + $O(n)$ (DFS) = $O(n)$

## Resource
[【每日一题】LeetCode 2876. Count Visited Nodes in a Directed Graph](https://www.youtube.com/watch?v=DZXJ1jK78kY&list=PLwdV8xC1EWHrtgsYCcDTXIMVaHSlsnLzL&index=7&ab_channel=HuifengGuan)