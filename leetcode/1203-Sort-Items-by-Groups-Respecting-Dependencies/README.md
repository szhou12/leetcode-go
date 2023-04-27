# [1203. Sort Items by Groups Respecting Dependencies](https://leetcode.com/problems/sort-items-by-groups-respecting-dependencies/description/)

## Solution idea

### Topological Sort
#### 思路总结
1. 想到解决思路并不困难，只要把图画出来，就可以很自然的想到嵌套的两层 Topological Sort: 先每一个group内进行sort，然后group之间进行sort
2. 难点: 在于实现上比较麻烦
    * `next` 和 `degree` 都使用 `HashMap` 会比较方便
    * `group = -1`的元素是“自由人”，不属于任一group，可以随意安排 (我一开始把它们都当成了一个group)。实现上，需要relabel每一个`group = -1`的元素，使它们的 group id 各自不同


## Resource
[【每日一题】1203. Sort Items by Groups Respecting Dependencies, 7/22/2020](https://www.youtube.com/watch?v=NP9ia_QU0l8&ab_channel=HuifengGuan)