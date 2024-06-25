# [3196. Maximize Total Cost of Alternating Subarrays](https://leetcode.com/problems/maximize-total-cost-of-alternating-subarrays/description/)

## Solution idea
### DP - House Robber
1. 你能在有限时间下，迅速把题意本质与House Robber关联起来吗？
2. 如果按题目要求，把array切成数段subarrays，会是什么样子？
    * 首先，看到一个subarray的内部: 符号的改变交错出现，并且，头元素的符号一定不变。
    * 然后，任意两个相邻的subarrays: subarray1, subarray2，它们的连接处只可能有两种情况：1. subarray1的尾元素符号不变 + subarray2的头元素符号不变；2. subarray1的尾元素符号变过 + subarray2的头元素符号不变。(头元素符号一定要保持不变)
3. 根据上面的两条观察，可以总结出: 对于i号元素，它与i-1号元素的关系只有两种：
    1. 如果i号元素符号保持不变，那它既能接变号的i-1号元素，也能接不变号的i-1号元素。挑大的接上。
    2. 如果i号元素符号变号，那它只能接不变号的i-1号元素，接上。
4. 看出与House Robber的关联来了吗？“变号”在这里就相当于“抢劫”。因为不能有连续两个相邻的元素都变号，也就是不能有连续两个相邻的房子都被抢劫。
5. Apply AP:
    1. Definition: 
        * `DP[i][0] :=` the maximum sum of subarrays in `nums[0:i]` (inclusive) without sign change at i-th element
        * `DP[i][1] :=` the maximum sum of subarrays in `nums[0:i]` (inclusive) with sign change at i-th element
    2. Base Cases:
        * `DP[0][0] = nums[0]`
        * `DP[0][1] = -inf` 表示“无意义”，因为头元素必不能变号
    3. Recurrence:
        * `DP[i][0] = max(DP[i-1][0], DP[i-1][1]) + nums[i]`
        * `DP[i][1] = DP[i-1][0] - nums[i]`
    4. Return:
        * `max(DP[n-1][0], DP[n-1][1])`

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 3196. Maximize Total Cost of Alternating Subarrays](https://www.youtube.com/watch?v=Hpgs7bHNUaU&ab_channel=HuifengGuan)