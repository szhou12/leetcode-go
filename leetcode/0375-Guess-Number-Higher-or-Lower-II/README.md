# [375. Guess Number Higher or Lower II](https://leetcode.com/problems/guess-number-higher-or-lower-ii/description/)

## Solution idea

Also known as **Min-max** problem; interval DP problem.

Define `DP[i][j]` as the min cost to guess for an interval `[i...j]`.

### Top-down: Recursion + DP Memoization (easy to understand)


* **Key 1**: Take **mx** of left-half subarray and right-half subarry. This is because we need a cost that can cover both halves

* **Key 2**: Out of all split points, choose the one that gives the **min** cost.

Time complexity = $O(n^2*n) = O(n^3)$

Space complexity = $O(n^2)$

**Time complexity Explanation**:

With memoization, the algorithm has no repeated recursion, meaning that we split into and traverse each continuous subarray only once. For an array of lenth $n$, the total number of continuous subarray is $O(1+2+3+\cdots+n) = O(\frac{n(n+1)}{2}) = O(n^2)$. For each subarray, we traverse each element fo find the best splitting point, which is $O(n)$. Therefore, in total, the time complexity = $O(n^3)$.

### Bottom-up: DP

如果解题角度是，从小区间开始先算再往大区间算，则 DP 的两层for-loop 为:
1. 外层loop是循环 length：从最小的length的区间开始，逐步增长
2. 内层loop是循环 start position: 这里要注意，最大的start position 加上length后不能越界, 终止条件要注意用 `j-i+1=len => j=i+len-1 => i <= i+len-1`

DP 的loop 搭建好了以后，要考虑边界条件：
* 考虑 循环的边界取值带入 update DP 的步骤时，会不会发生越界。要把越界的值提前算好 (默认值设成0？还是无限大/小？此题设成0)。

Time complexity = $O(n^2*n) = O(n^3)$

Space complexity = $O(n^2)$

## Resource

Top-down: [【递归 + 动态规划】leecode 375 Guess Number Higher or Lower II](https://www.youtube.com/watch?v=YTGHiGH_oTg&ab_channel=guoguowg)

Bottom-up: [【每日一题】375. Guess Number Higher or Lower II, 04/27/2019](https://www.youtube.com/watch?v=VfJPDNG0nYM&ab_channel=HuifengGuan); [wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Dynamic_Programming/375.Guess-Number-Higher-or-Lower-II)

Time complexity analysis: [DP and memo solution explanation](https://leetcode.com/problems/guess-number-higher-or-lower-ii/solutions/197981/dp-and-memo-solution-explanation/?orderBy=most_votes&page=3)

Total number of continuous subarrays: [The total number of subarrays](https://math.stackexchange.com/questions/1194584/the-total-number-of-subarrays)