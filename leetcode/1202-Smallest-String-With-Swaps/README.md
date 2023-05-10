# [1202. Smallest String With Swaps](https://leetcode.com/problems/smallest-string-with-swaps/description/)

## Solution idea
### Union Find
#### 要点总结
1. 解法不难想到，题目给出：每一个pair可以无限次swap，就可以想象成一个 connected component 中的任意两个节点可以随意 swap，也就是具有了 transitive。所以，第一步就是利用 Union Find 找出所有的 connected component 并把 节点 归类到对应的 component 中去。
    * Component 储存 每个字符的index
2. 稍微棘手一些的是如何排序：每个 component 中，通过成员的index找到对应的字符，sort, 再对应index把排好序的字符依次放回 source string

## Resource
[【每日一题】1202. Smallest String With Swaps, 10/06/2019](https://www.youtube.com/watch?v=oCifdls4XpE)