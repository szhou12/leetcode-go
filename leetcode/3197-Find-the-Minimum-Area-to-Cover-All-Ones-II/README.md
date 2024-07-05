# [3197. Find the Minimum Area to Cover All Ones II](https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-ii/description/)

## Solution idea
1. 需要三个不重叠的最小区域，能够覆盖所有1的格子。
2. 首先，可以把整片区域划分成不重叠的三个子区域一共有下图所示的六种分法。
3. 然后，对每个子区域，使用[3195](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3195-Find-the-Minimum-Area-to-Cover-All-Ones-I)的方法找出最小覆盖所有1的矩形。这样，就能保证满足题目要求。

Time complexity = $O(\max\{m^2, n^2, mn\}\cdot mn)$

### 一图胜千言
![autodraw 6_28_2024 (1)](https://github.com/szhou12/leetcode-go/assets/35708194/658f0d8c-d4cd-4fa1-ba4a-28bd92d67895)


## Resource
[【每日一题】LeetCode 3197. Find the Minimum Area to Cover All Ones II](https://www.youtube.com/watch?v=EZcmzaMmRwM&t=59s&ab_channel=HuifengGuan)
