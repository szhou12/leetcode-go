# [3027. Find the Number of Ways to Place People II](https://leetcode.com/problems/find-the-number-of-ways-to-place-people-ii/description/)

## Solution idea
### Greedy
1. Constraint并不大，可以直接两层循环暴力搜索。
2. 首先，将所有点按照x坐标升序排序。难点在于，确定一点为Alice后，对于任意一点作为Bob，如何确定是否是合法的Bob？需要满足的条件有两条：
    1. Bob的y坐标有一个上限：不高于Alice的y坐标 `yb <= upper = ya`
    2. Bob的y坐标有一个下限：高于Alice和Bob之间所有点中最高的点 `yb > lower`
        * 实现上，每确定一个Alice，在loop Bob时，需要维护一个`lower`作为往期看过来的最高点。从而判断当前的Bob是否合法。(如图1)
3. 还有一个问题。如果同时有多个x坐标相同，y坐标不同的候选Bob怎么办？
    * Solution: 将所有点先按照x坐标升序排序后，再按照y坐标降序排序。
    * Why? 因为这样就总是先看最高的那个Bob。如果先看的是最低的Bob，那么比它高的Bob们就会被圈进去，不符合要求。(如图2)

### 一图胜千言


Time complexity = $O(n^2)$

## Resource
[【每日一题】LeetCode 3027. Find the Number of Ways to Place People II](https://www.youtube.com/watch?v=cZjZSfRXePU&ab_channel=HuifengGuan)