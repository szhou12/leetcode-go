# [2093. Minimum Cost to Reach City With Discounts](https://leetcode.ca/2021-12-16-2093-Minimum-Cost-to-Reach-City-With-Discounts/#2093-minimum-cost-to-reach-city-with-discounts)

## Solution idea

### Dijkstra
#### 要点总结: 
1. 任意两个node之间有 "两条边"， 一条edge的weight = toll，另一条edge的weight = toll/2. 
2. Dijkstra算法在expand的时候, 要分别把两条边的情况都Push进去.
3. Dijkstra算法的结果储存在 二维矩阵 = # nodes X # discounts used
4. **node储存双状态**: node 储存 位置信息+已使用discounts次数信息

Time complexity = $O(E \log E)$ where $E = n * m$, $m = $ # of discounts


## Resource
[【LeetCode】2093. Minimum Cost to Reach City With Discounts](https://www.bilibili.com/video/BV1MU4y1Z7Z7/?spm_id_from=333.999.0.0&vd_source=0c02ef6f6e7a2b0959d7dd28e9e49da4)