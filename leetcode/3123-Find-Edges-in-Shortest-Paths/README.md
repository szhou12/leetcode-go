# [3123. Find Edges in Shortest Paths](https://leetcode.com/problems/find-edges-in-shortest-paths/description/)

## Solution idea
### Dijkstra
1. 突破点：题干中明确说了是从单一起点`node 0`开始。同时，constraint的量级在`5*10^4`。结合起来，单一起点，边有权重，并且量级很大的情况下，使用的算法就是Dijkstra。
2. 如何确定一个节点或者一条边是否在最短路径上？
    - Let $d(x, y)$ be the shortest distance from node x to node y.
    - Let $d(0, n-1)$ be the shortest distance from source node 0 to destination node n-1.
3. 判断一个节点`node i`在最短路径上的充要条件: $d(0, i) + d(i, n-1) = d(0, n-1)$
    - Claim: if $d(0, i) + d(i, n-1) = d(0, n-1)$, then node $i$ is on the shortest path from node 0 to node n-1.
    - Prove:
        - Suppose $i$ is not on the shortest path from node 0 to node n-1.
        - Then $d(0, i) + d(i, n-1) > d(0, n-1)$, contradicting the assumption that $d(0, i) + d(i, n-1) = d(0, n-1)$.
4. 判断一条边`edge (a, b)`在最短路径上的充要条件: $d(0, a) + w(a, b) + d(b, n-1) = d(0, n-1)$
    - Claim: if $d(0, a) + w(a, b) + d(b, n-1) = d(0, n-1)$, then edge $(a, b)$ is on the shortest path from node 0 to node n-1.
    - Prove: *我的证明，正确性有待考证*
        - Suppose edge $(a, b)$ is not on the shortest path from node 0 to node n-1.
        - Then there must be other edges connectting $a$ and $b$ instead of direct edge $(a, b)$ that consist of subpaths for the shortest path. Let $K$ represent all the intermediary nodes on these edges. Formally, $d(0, a) + w(a, K) + w(K, b) + d(b, n-1) = d(0, n-1)$ and $w(a, K) + w(K, b) < w(a, b)$. 
        - However, a fundamental property derived from the definition of shortest paths is that any subpath of a shortest path is itself the shortest path between its endpoints.
        - Given the assumption that $d(0, a) + w(a, b) + d(b, n-1) = d(0, n-1)$, since $d(0, n-1)$ is the shortest path, then $d(0, a)$, $w(a, b)$, $d(b, n-1)$ are three shoretest subpaths.
        - It means that edge $(a, b)$ is the shortest path connecting node $a$ and node $b$.
        - It means that $w(a, K) + w(K, b) = w(a, b)$, contradiction. 
5. 代码实现技巧:
    - **Dijkstra跑两遍**: 正向`node 0`为起点跑一遍，逆向`node n-1`为起点跑一遍。
    - 再通过判据 $d(0, a) + w(a, b) + d(b, n-1) = d(0, n-1)$ 检查任意一条边是否在最短路径上。


Claim: if $d(0, i) + d(i, n-1) = d(0, n-1)$, then node $i$ is on the shortest path from node 0 to node n-1.

Prove:

Suppose $i$ is not on the shortest path from node 0 to node n-1.

Then $d(0, i) + d(i, n-1) > d(0, n-1)$, contradicting the assumption that $d(0, i) + d(i, n-1) = d(0, n-1)$.


 
Time complexity = $O(E\log E)$

## Resource
[【每日一题】LeetCode 3123. Find Edges in Shortest Paths](https://www.youtube.com/watch?v=ztiM0oqBKog&ab_channel=HuifengGuan)