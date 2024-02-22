# [3041. Maximize Consecutive Elements in an Array After Modification](https://leetcode.com/problems/maximize-consecutive-elements-in-an-array-after-modification/description/)

## Solution idea
### Sorting + DP (2-D)
1. **Key Observation**: 第一要有的直觉是，排序先！Sort first to make elements with close values group together.
    * NOTE: 题目的数据量级在 $10^5$, 允许 $O(n\log n)$的算法！
2. DP:
    1. Definition:
        * `dp[i][0] :=` len of longest consecutive elements in `nums[:i]` (inclusive) ending at `i` with `nums[i]+0`
        * `dp[i][1] :=` len of longest consecutive elements in `nums[:i]` (inclusive) ending at `i` with `nums[i]+1`
        * Do you understand why we need 2nd dimension to annotate status +0 and +1?
    2. Base case: `dp[0][0] = 1, dp[0][1] = 1`
    3. Recurrence:
        * CASE 0 `另起炉灶`:
            * `dp[i][0] = 1`
            * `dp[i][1] = 1`
        * CASE 1 `nums[i]-nums[i-1] == 2`: 
            * `dp[i][0] = max(dp[i][0], dp[i-1][1]+1)`
        * CASE 2 `nums[i]-nums[i-1] == 1`:
            * `dp[i][0] = max(dp[i][0], dp[i-1][0]+1)`
            * `dp[i][1] = max(dp[i][1], dp[i-1][1]+1)`
        * CASE 3 `nums[i]-nums[i-1] == 0`:
            * `dp[i][1] = max(dp[i][1], dp[i-1][0]+1)`
            * `dp[i][0] = max(dp[i][0], dp[i-1][0])`
            * `dp[i][1] = max(dp[i][1], dp[i-1][1])`

3. 一图胜千言
```
X X X i-1, i X X X
       a   b
case 1: b-a = 2 ==> a+1, b+0
        dp[i][0] = dp[i-1][1] + 1
case 2: b-a = 1 ==> (a+1, b+1) OR (a+0, b+0)
        dp[i][0] = dp[i-1][0] + 1
        dp[i][1] = dp[i-1][1] + 1
case 3: b-a = 0 ==> (a+0, b+1) OR (throw out b, 不用b做贡献, directly inherits len at a+0 or a+1)
        dp[i][1] = dp[i-1][0] + 1
        dp[i][1] = dp[i-1][1]
        dp[i][0] = dp[i-1][0]
case 0: b starts over (另起炉灶)
        dp[i][0] = 1
        dp[i][1] = 1
```

Time complexity = $O(n\log n)$

## Resource
[【每日一题】LeetCode 3041. Maximize Consecutive Elements in an Array After Modification](https://www.youtube.com/watch?v=Pepe16TjyKQ&ab_channel=HuifengGuan)