# [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)

## Solution idea

一开始想的是 `DFS` 求 all permuations, 但是 Time Complexity 太高.

然后想的是用 `map` 去记录每一个 char. 

最后想的是 用不着 `map`, 可以直接用一个 26 长度的数组，对于 `s` 中每个char在相应位置 `+1`; 再在对于 `t` 中每个char在相应位置 `-1`

Time complexity = $O(n)$