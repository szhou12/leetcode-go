# [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/description/)

## Solution idea

### Sliding Window
* 整体code框架与 [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/description/) 基本一致
* 不同的地方在于:
    1. 缩短左边界的条件: 本题要求 sliding window 是定长, 一旦左边界和右边界的距离 >= 要求的定长, 就需要缩短左边界
    2. update result 的条件: 每个字符出现次数都符合要求, 添加一个result

* 因为 sliding window 是定长, 写法也可以是单指针, 写法类似 [1052. Grumpy Bookstore Owner](https://leetcode.com/problems/grumpy-bookstore-owner/description/)

Time complexity = $O(n)$

## Resource
[滑动窗口算法](https://labuladong.github.io/algo/2/20/27/)
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Hash/438.Find-All-Anagrams-in-a-String)