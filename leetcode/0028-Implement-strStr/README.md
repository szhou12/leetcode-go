# [28. Find the Index of the First Occurrence in a String](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/)

## Solution idea
### Sliding Window
Check each `haystack[l:r]` where `[l:r]` is defined by the length of `needle`

Time complexity = $O(n)$ where $n$ is the length of haystack

### KMP
* Step 1: 对 模式串 (pattern string) 使用 KMP 算法，构造 `lsp` 数组
* Step 2: 配合 `lsp`，对 目标串 (target string) 使用 KMP 算法

## Resource
* KMP: [【每日一题】28. Implement strStr(), 2/3/2021](https://www.youtube.com/watch?v=Kxo1f-SUwUc&ab_channel=HuifengGuan)