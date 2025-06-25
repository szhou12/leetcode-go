# [408. Valid Word Abbreviation](https://leetcode.com/problems/valid-word-abbreviation/description/)

## Problem Description
Check if `word` can be abbreviated to `abbr`.

```
Example 1: word = "internationalization", abbr = "i12iz4n"
Output: true
Explanation: "internationalization" -> "i [nternational] iz [atio] n" -> "i12iz4n"
```
```
Example 2: word = "apple", abbr = "a2e"
Output: false
Explanation: "apple" cannot be abbreviated to "a2e"
```

## Solution idea
### Two Pointers

Time complexity = $O(n)$