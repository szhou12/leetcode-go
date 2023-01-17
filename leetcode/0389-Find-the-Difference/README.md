# [389. Find the Difference](https://leetcode.com/problems/find-the-difference/description/)

## Solution idea

### HashMap / Array of length 26

* 用 HashMap/Array of Length 26 记录 `s`中出现的字母，查询`t`中哪个字母没出现

Time complexity = $O(n)$

### XOR - 同性相消

利用异或的性质，`X^X = 0`，将 s 和 t 依次异或，最终多出来的字符就是最后异或的结果。


## Resource

* array of length 26: [贾考博 LeetCode 389. Find the Difference - 工作随便找系列2](https://www.youtube.com/watch?v=ctd8TZI8uL0&ab_channel=%E8%B4%BE%E8%80%83%E5%8D%9A)

* XOR: [halfrost/LeetCode-Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0389.Find-the-Difference)