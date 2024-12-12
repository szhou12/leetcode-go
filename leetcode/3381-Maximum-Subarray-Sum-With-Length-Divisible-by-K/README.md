# [3381. Maximum Subarray Sum With Length Divisible by K](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/description/)

## Solution idea
### DP (Maximum Subarray Sum) + Prefix Sum
- 本题考查对 Classic Kadane's Algorithm ([53. Maximum Subarray](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0053-Maximum-Subarray)) 本质有多了解。
    1. Let's review **Classic Kadane's Algorithm**: 本质上是从任意长度的subarray中找到和最大的subarray，长度至少为1。根据`DP[i]`的定义，最小成立的单元为i号元素自己，我们考虑它能否与前面的和“接续”上。$$DP_i = max(DP_{i-1}, 0) + x_i$$
    2. 推论: 拓展到本题，额外条件是subarray的长度必须是k的倍数，所以最小成立的单元是长度正好为k的subarray。如果定义i号元素为末尾元素，那么，最小单元就是从i-k+1到i的这段subarray，我们考虑这一段能否与前面的和“接续”上。既然预留了一段k长的subarray, 那么“接续”的和也要往前跳过k格，看到`DP[i-k]`。注意到，我们也需要前面的和表示的是k的倍数的和，而不是任意长度的和，这个check我们通过初始化为-inf来解决，也就是说，当`DP[i-k]`在与0取max时，我们同时保证了前面的和的“有效性”(大于-inf说明是长度为k倍的和)以及“可接续性”(大于0说明接上能获得更大收益，否则不如“另起炉灶”)。我们还需要修改状态转移方程i的取值范围为`[k : n-1]`，base case为`sum([0 : k-1])`. $$DP_i = max(DP_{i-k}, 0) + \sum_{i-k+1}^{i}x_i$$
    3. 使用 Prefix Sum 快速计算某一段subarray的和。
- DP template:
```
Definition:
    DP[i] := Maximum k-sized subarry sum ending at i.
Base Case:
    DP[k-1] = sum(nums[0:k-1]) (inclusive) = prefixSum[k-1]
Recurrence:
    DP[i] = max(DP[i-k], 0) + (prefixSum[i] - prefixSum[i-k]) for k <= i < n
```

Time complexity = $O(n)$

## Resource
- [Prefix Sum Kadane](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/solutions/6124522/prefix-sum-kadane/)
    - Look for *David Espinosa* comment
- [DP in 5 Lines of Python3](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/solutions/6124606/dp-in-5-lines-of-python3/)
- [7ms - O(N) - Kadane with benefits](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/solutions/6124512/7ms-o-n-kadane-with-benefits/)