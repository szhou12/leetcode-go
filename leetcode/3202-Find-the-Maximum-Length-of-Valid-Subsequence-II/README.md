# [3202. Find the Maximum Length of Valid Subsequence II](https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/description/)

## Solution idea
### DP + Modulo
```
(a + b) % k = r = (a % k + b % k) % k
=> a + b = k * n + r
=> b = k * n + r - a
=> b % k = (r - a + k * n) % k
=>       = [(r - a) % k + (k * n) % k] % k <Distributivity>
=>       = [(r - a) % k + 0] % k
=>       = [(r - a) % k + k % k] % k
=>       = (r - a + k) % k <Distributivity>
```
Let `cur = nums[i] % k`. For a given `r`, we have `prev = (r - cur + k) % k`. 