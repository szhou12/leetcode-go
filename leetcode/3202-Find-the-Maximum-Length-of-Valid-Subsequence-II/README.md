# [3202. Find the Maximum Length of Valid Subsequence II](https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/description/)

## Solution idea
### DP + Modulo
1. Modulo:
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

Why write as `(r - a + k)` instead of `(r - a)`? The step ensures that even if `r - a` is negative, adding `k` (which is the modulus) brings the total back into a positive range, and taking modulo `k` ensures it wraps around correctly if it exceeds `k`.

2. DP:
    1. Definition: `DP[cur][r] :=` the max length of subsequence ending at **value** `cur = (nums[i] % k)` that satisfies each pair `(sub[j-1] + sub[j]) % k == r`.
    2. Base case: `DP[X][r] = 0` where `X = 0, ..., k-1`.
    3. Recurrence: `DP[cur][r] = max(DP[cur][i], DP[prev][r] + 1)` where `prev = (r - cur + k) % k`.


Time complexity = $O(n\cdot k)$

Space complexity = $O(k\cdot k)$

## Resource
[Yes it's a 2d DP !!](https://leetcode.com/problems/find-the-maximum-length-of-valid-subsequence-ii/solutions/5389322/yes-it-s-a-2d-dp/)