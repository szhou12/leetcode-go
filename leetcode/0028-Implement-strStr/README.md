# [28. Find the Index of the First Occurrence in a String](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/)

## Solution idea
### Sliding Window
Check each `haystack[l:r]` where `[l:r]` is defined by the length of `needle`

Time complexity = $O(n)$ where $n$ is the length of haystack

### KMP
* Step 1: 对 模式串 (pattern string) 使用 KMP 算法，构造 `lsp` 数组
    * base case: `lsp[0] = 0` 因为我们要求"真"前缀和“真”后缀
* Step 2: 配合 `lsp`，对 目标串 (target string) 使用 KMP 算法
    * edge case: 如果`needle`为空 或者 `haystack`为空，则找不到，直接返回合理的值
    * base case: `dp[0] = 1 if haystack[0] == needle[0]`
    * check before loop: `if len(needle) == 1 && dp[0] == 1`, return `0`。因为loop是从`i=1`开始的，所以这里要提前check
* Step 1 和 Step 2 的loop部分逻辑是一样的。稍微不同的是:
    * Step 1 loop: `needle`的后缀和`needle`的前缀比较
    * Step 2 loop: `haystack`的后缀和`needle`的前缀比较

Time complexity = $O(n)$ where $n$ is the length of haystack

## Resource
* KMP: [【每日一题】28. Implement strStr(), 2/3/2021](https://www.youtube.com/watch?v=Kxo1f-SUwUc&ab_channel=HuifengGuan)