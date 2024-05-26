# [3163. String Compression III](https://leetcode.com/problems/string-compression-iii/description/)

## Solution idea
### Sliding Window (Flex)
1. 标准的Sliding Window (Flex) 题目。值的注意的是，本题涉及go语言中concatenate strings的优化问题：loop中使用`strings.Builder`而不是`+`来concatenate strings。这是因为go中string是immutable，每次`+`都会创建一个新的string，从而导致TLE。

Time complexity = $O(n)$