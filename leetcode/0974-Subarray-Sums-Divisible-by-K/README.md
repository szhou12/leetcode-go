# [974. Subarray Sums Divisible by K](https://leetcode.com/problems/subarray-sums-divisible-by-k/description/)

## Solution idea
### Math + Prefix Sum + Hash Map (frequency map)
1. 80%的题看到subarray sum，都要想到用 Prefix Sum 来解决
2. 数学: `for x = a - b, x % k = 0 iff a % k == b % k`
3. Mod trick: `(sum % k + k) % k` to avoid sum % k < 0
4. Prefix Sum + Hash Map 时初始化Map技巧: `dict[0] = 1` 表示 prefixSum % k = 0, 也就是 prefixSum = 0, 这个0指代index=-1时的prefixSum[-1], value=1表示j=-1算一个count，虚拟地认为index=-1是存在的。

Time complexity = $O(n)$

Space complexity = $O(k)$

## Resource
[【每日一题】LeetCode 974. Subarray Sums Divisible by K](https://www.youtube.com/watch?v=6lik5HreRHI&ab_channel=HuifengGuan)