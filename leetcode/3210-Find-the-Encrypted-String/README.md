# [3210. Find the Encrypted String](https://leetcode.com/problems/find-the-encrypted-string/description/)

## Solution idea
### Circular String + Modulo
1. 基本上就是circular string的思想：从input string任意位置开始，找到可以“接上”的新string。
    1. 确定起始位置：`k`可以远远大于input string的长度，所以利用取模，`k % n`将起始位置映射在`[0, n-1]`之间。
    2. 剩下的就是Circular String的做法：将input string复制一倍，然后从确定的起始位置开始，取长度`n`的新string。
