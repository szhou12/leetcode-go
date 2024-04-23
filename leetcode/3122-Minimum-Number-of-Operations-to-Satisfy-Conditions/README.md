# [3122. Minimum Number of Operations to Satisfy Conditions](https://leetcode.com/problems/minimum-number-of-operations-to-satisfy-conditions/description/)

## Solution idea
### 2-D DP
1. 第一眼看过去，容易想到的是greedy algorithm: 先以列为单位，看每一列，找出频次最高的数，把其他数变成这个数。再以行为单位，调整使每相邻的数不同。但是，这样做有很多corner cases要处理。比如：1. 如果一列中所有数频次都相同选哪一个？2.先调整列再调整行会增加operation次数(因为看不到全局，本可以在调整列时就一次性完成的操作额外留到了调整行的步骤)
2. 突破点: 看到题目的constraint `1 <= n, m <= 1000`。说明，每个格子一一看过去需要`10^6`次操作，这是可以接受的。同时，`0 <= grid[i][j] <= 9`暗示了如果枚举`0~9`，时间复杂度在`10^7`，刚刚好可以接受。所以，可能我们需要通过**枚举**来解决问题。正式的，每一列我们尝试枚举`0~9`，看哪一个数要求的操作数最少。如何保证列与列之间的数不同？通过DP的转移方程来传递信息。
3. Definition: `dp[i][p] := min operations to fill up first i cols with i-th col being set up with number p`
    - `i`: col numbers so far; `p`: 0~9
    - Base case: 0-th col doesn't depend on previous cols. So `dp[0][p] = min ops of first 0 col = cost to set 0-th col to p`
    - Recurrence: `dp[i][p] = min{dp[i-1][q] + cost(i, p) for q=0..9} for p=0..9 and p!=q`
4. 可优化的地方：DP的解法中可优化，每一列不用枚举所有0~9，只用找出当前列频次最高，次高，第三高的三个数，枚举这三个数就可以。为什么？题目要求当前列与邻接的列的数不同。所以三个数足够找出一个数使其与左右两列的数不同，同时它的频次又是top 3，所以相较于其他频次低的数，它导致的操作数肯定更少。

Time complexity = $O(mn\cdot 10 \cdot 10)$

## Resource
[【每日一题】LeetCode 3122. Minimum Number of Operations to Satisfy Conditions](https://www.youtube.com/watch?v=L2j-KCX4Y5Y&ab_channel=HuifengGuan)