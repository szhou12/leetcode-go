# [151. Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/)

## Solution idea

### Two Pointers

* Step 1: remove leading spaces, trailing spaces, and extra spaces betweeen words
    * Implementation 是 [27. Remove Element](https://leetcode.com/problems/remove-element/) 的进阶版

* Step 2a: reverse all characters all at once. 一整个字符串翻转

* Step 2b: reverse all chars per word. 每个词每个词单独翻转

Time complexity = $O(n^2)$

## Resource

[代码随想录-151.翻转字符串里的单词](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0151.%E7%BF%BB%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2%E9%87%8C%E7%9A%84%E5%8D%95%E8%AF%8D.md)
