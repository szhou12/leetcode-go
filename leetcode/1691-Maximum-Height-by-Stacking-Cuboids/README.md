# [1691. Maximum Height by Stacking Cuboids](https://leetcode.com/problems/maximum-height-by-stacking-cuboids/)

## Solution idea

### DP - 解法一: 同时考虑每个cuboid的6种变体

* 题目要求每个cuboid可以rotate, 所以每个cuboid可以有6种放置方法 ($3! = 6$)
    * 注意: 如果其中有两个变量相等 (e.g. l==w), 则只有3种distinct放置方法 ($\frac{3!}{2!} = 3$)
    * 注意: 如果三个个变量相等 (e.g. l==w==h), 则只有1种distinct放置方法 ($\frac{3!}{3!} = 1$)
* cuboid 6种放置方法, 每一种都加入`cubes`数组中.
* 为了防止任何一个 LIS 中, 同一个cuboid因为其不同变体被重复放入 (当有两个及以上变量相等时，就有可能被LIS算法重复放入), 每一个cuboid的变体额外加入index信息 (positional encoding)
* 预排序规则: 优先级 `width > length > height > index`
    * `width` 升序排列
    * `width`相同时, `length` 升序排列
    * `width`相同时, `length`相同时, `height` 升序排列
    * `width`相同时, `length`相同时, `height` 相同时, 升序排列
* 因为加入index信息，状态转移方程的条件就更新成:
```
width[j] <= width[i] && length[j] <= length[i] && height[j] <= height[i] && index[j] != index[i] for 0 <= j < i
```

* Definition:
```
dp[i] = max height of stacked cuboids ending at cube (one cuboid's variant) i
```

* Base Cases:
```
dp[i] = cubes[i].height
```

* Recurrence:
```
dp[i] = max(dp[j]+cubes[i].height) if width[j] <= width[i] && length[j] <= length[i] && height[j] <= height[i] && index[j] != index[i] for 0 <= j < i
```

* Result
```
res = max(dp[i]) for all i
```

Time complexity = $O(6n*6n)$ because length of `dp[]` equals to length of `cubes[]`, which is 6 times length of cuboids

### DP - 解法二: 每个cuboid只考虑1种变体
* 每个cuboid只需要考虑1种变体: `(min(l, w, h), mid(l, w, h), max(l, w, h))`

* 重要数学推论:
    1. cuboid A $(a_1, a_2, a_3)$ can be on top of cuboid B $(b_1, b_2, b_3)$ `<=>` $a_1 \leq b_1$ and $a_2 \leq b_2$ and $a_3 \leq b_3$ (根据题意)
    2. $a_1 \leq b_1$ and $a_2 \leq b_2$ and $a_3 \leq b_3$ `=>` $min(a_1, a_2, a_3) \leq min(b_1, b_2, b_3)$ and $mid(a_1, a_2, a_3) \leq mid(b_1, b_2, b_3)$ and $max(a_1, a_2, a_3) \leq max(b_1, b_2, b_3)$
        * 如果A的三个变量都小于等于B, 那么A的最小变量一定也小于等于B的最小变量
        * 如果A的三个变量都小于等于B, 那么A的最大变量一定也小于等于B的最大变量
        * A的中间值变量一定也小于等于B的中间值变量 (关系通过画图推导出来)

* 再按照常规LIS去解题

Time complexity = $O(n*n)$

## Resource
[【每日一题】1691. Maximum Height by Stacking Cuboids, 12/16/2020](https://www.youtube.com/watch?v=nyJe6_4_MTs)