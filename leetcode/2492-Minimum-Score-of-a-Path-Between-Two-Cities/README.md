# [2492. Minimum Score of a Path Between Two Cities](https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/description/)

## Solution idea
### Union Find

#### 要点总结
1. 注意题目一个重要的前提条件：一条边或一个节点可以重复多次访问。这意味着，我们可以特意"绕远路"只为了找到最短的那条边。这样问题可以转化为 :arrow_right: city 1 与 city N 之间的score本质，就是这两个节点所在连通图里最短的边
2. 解题步骤:
    1. 用Union Find将所有的节点标记联通，将所有节点归到所属的连通图 (集合) 里
    2. 再遍历一遍所有的边，找到 city 1 (or city N) 所属连通图的节点以及对应的最短的边


## Resource
[【每日一题】LeetCode 2492. Minimum Score of a Path Between Two Cities](https://www.youtube.com/watch?v=Tj-L6JTiBdg&ab_channel=HuifengGuan)