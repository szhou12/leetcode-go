# [2392. Build a Matrix With Conditions](https://leetcode.com/problems/build-a-matrix-with-conditions/description/)

## Solution idea

### Topological Sort
#### 要点总结
1. 难点一：如何把问题转化为 Topological Sort 求解? 或者说，怎么想到是要用 Topological Sort 解题的？
    * 题目中提供了一条重要暗示: `rowConditions`, `colConditions`. 这两个约束条件限定了哪些元素要排在其他元素前面。
    * "遵守一系列restrictions/约束排列出一个顺序" $\Rightarrow $ Topological Sort 是适合解这类题型的思路，因为，“遵守一系列restrictions/约束” 可以具像化为 有向图
    * 具体这道题: "a 要在 b 左边/上边" $\Rightarrow $ "a 要排在 b 前面" $\Rightarrow $ "先到a，再到b" $\Rightarrow $ `a -> b`
2. 难点二：Topological Sort 得到的order怎么应用在填表(构建matrix)上？
    * `rowConditions` 得到的order是每个元素的 x坐标
    * `colConditions` 得到的order是每个元素的 y坐标
    * 合起来，`(x, y)` 就是每个元素的位置
    * 注意：因为本题只需要在`k X k`的方阵中填写k个数字，所以我们可以让所有元素都不同行不同列。这样横向的拓扑关系和纵向的拓扑关系可以独立处理，互不干扰。这就解释了为什么横向order和纵向order只要有解，合起来就是我们需要的坐标。
3. 难点三：如何判定无解的情况？
    * 无解的情况：注意到 Topological Sort 一大功能就是检测图中是否有环。这道题无解的情况就发生在构建的有向图中出现了环 (e.g. 1要在2左边，2要在3左边，3要在1左边)
4. 难点四：test cases 中有相同edge重复出现的情况
    * 实现时，adjacency list representation `next` 使用 `[][]int` 代替 `[]map[int]bool` 防止 degree多算了而next没有对应添加


#### 代码结构
```
Per rowConditions:
Step 1: Reconstruct adj-list representation + Calculate degree
Step 2: Topological Sort on row order

Per colConditions:
Step 1: Reconstruct adj-list representation + Calculate degree
Step 2: Topological Sort on col roder

Step 3: Fill up the matrix by combining row order + col order
```

Time complexity = $O(k)$

## Resource
[【每日一题】LeetCode 2392. Build a Matrix With Conditions](https://www.youtube.com/watch?v=ManXfvtm1tc&ab_channel=HuifengGuan)