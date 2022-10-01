# [131. Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/)

## Solution idea

### DFS

```
# levels = len(s)
# branches = [last cut ... end of string]
```

Time complexity = $O(2^n)$

解释: **切割问题其实是一种组合问题！** 每个位置都有切或者不切两种情况