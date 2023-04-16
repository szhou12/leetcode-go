# [2392. Build a Matrix With Conditions](https://leetcode.com/problems/build-a-matrix-with-conditions/description/)

## Solution idea

### Topological Sort
#### 要点总结
1. 难点一：如何把问题转化为 Topological Sort 求解? 或者说，怎么想到是要用 Topological Sort 解题的？
    * 题目中提供了一条重要暗示: `rowConditions`, `colConditions`. 这两个约束条件限定了哪些元素要排在其他元素前面。
    * "遵守一系列restrictions/约束排列出一个顺序" $\Rightarrow $ Topological Sort 是适合解这类题型的思路，因为，“遵守一系列restrictions/约束” 可以具像化为 有向图
    * 具体这道题: "a 要在 b 左边/上边" $\Rightarrow $ "a 要排在 b 前面" $\Rightarrow $ "先到a，再到b" $\Rightarrow $ `a -> b`
2. 难点二：Topological Sort 得到的order怎么应用在填表(构建matrix)上？
3. 难点三：如何判定无解的情况？
    * 无解的情况：注意到 Topological Sort 一大功能就是检测图中是否有环。这道题无解的情况就发生在构建的有向图中出现了环 (e.g. 1要在2左边，2要在3左边，3要在1左边)
4. 注意: test cases 中有相同edge重复出现的情况


## Resource
[【每日一题】LeetCode 2392. Build a Matrix With Conditions](https://www.youtube.com/watch?v=ManXfvtm1tc&ab_channel=HuifengGuan)