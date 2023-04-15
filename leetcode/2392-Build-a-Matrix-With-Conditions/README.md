# [2392. Build a Matrix With Conditions](https://leetcode.com/problems/build-a-matrix-with-conditions/description/)

## Solution idea

### Topological Sort
#### 要点总结
1. 难点一：如何把问题转化为 Topological Sort 求解? 或者说，怎么想到是要用 Topological Sort 解题的？
2. 难点二：Topological Sort 得到的order怎么应用在填表(构建matrix)上？
3. 难点三：如何判定无解的情况？
    * 无解的情况：注意到 Topological Sort 一大功能就是检测图中是否有环。这道题无解的情况就发生在构建的有向图中出现了环 (e.g. 1要在2左边，2要在3左边，3要在1左边)
4. 注意: test cases 中有相同edge重复出现的情况


## Resource
[【每日一题】LeetCode 2392. Build a Matrix With Conditions](https://www.youtube.com/watch?v=ManXfvtm1tc&ab_channel=HuifengGuan)