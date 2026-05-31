# [475. Heaters](https://leetcode.com/problems/heaters/)

## Solution idea

### Binary Search - Find Lower Bound
#### 要点总结
站在每个house的角度，看它左右两边最近的heaters的位置，取两者最短即为该house所需最短半径。注意边界情况，情况一：house在所有heaters的右边，情况二：house在所有heaters的左边。每个house的最短半径中取最大值即为可以覆盖所有house的最短半径。

Time complexity = $O(n \log m)$ where $n$ is the number of houses and $m$ is the number of heaters, Space complexity = $O(1)$

## Resource
[【每日一题】475. Heaters, 10/22/2019](https://www.youtube.com/watch?v=25LSSsAGLDw)