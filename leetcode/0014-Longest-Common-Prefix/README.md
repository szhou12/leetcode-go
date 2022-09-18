# [14. Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/)

## Solution idea

Step 1: sort all words by their length in ascending order; treat first word as `prefix`

Step 2: compare each char between `prefix` and any other words; shorten prefix if current char doesn't match

Time complexity = $O(n\log n + nm)$ where $n$ number of words in `strs` and $m$ length of prefix