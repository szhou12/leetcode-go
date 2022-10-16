# [796. Rotate String](https://leetcode.com/problems/rotate-string/)

## Solution idea

* 方法很巧妙，但是不容易想到 (背下来！！！)：
    * 判断 `goal` 是不是 `s`的 rotated string, 只需要 concat `s+s`, 然后再判断 `goal`是不是`s+s`的substring即可
    * The 'shift' operation can be regarded as a 'sliding' operation on `s+s`

Example:
```
s = "abcde", goal = "cdeab"
s+s = "abcdeabcde"
s+s = "ab cdeab cde"
           goal         
```

Time complexity = $O(n)$