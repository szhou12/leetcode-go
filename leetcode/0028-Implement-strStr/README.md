# [28. Find the Index of the First Occurrence in a String](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/)

## Solution idea
### Sliding Window
Check each `haystack[l:r]` where `[l:r]` is defined by the length of `needle`

Time complexity = $O(n)$ where $n$ is the length of haystack