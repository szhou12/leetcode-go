# [959. Regions Cut By Slashes](https://leetcode.com/problems/regions-cut-by-slashes/)

## Solution idea
### Union Find
#### 思路总结
1. 解题之前先问自己一个问题：什么时候划分会多出一个封闭区间？提示：把每个点分成两类 - 在边界上的"外围点"为一类，"内部点"为一类
    * 任意两个外围点之间连线，就会多出一个封闭区间。
    * 一个外围点 + 一个内部点之间连线，不会多出一个封闭区间，但内部点会被同化进外围点的聚类 (component) 中。
2. 初始状态：
    * 外围点：聚为一类
    * 内部点：未被聚类，单独为类


## Resource
[【每日一题】959. Regions Cut By Slashes, 2/12/2020](https://www.youtube.com/watch?v=46wmnzVyTv0&ab_channel=HuifengGuan)