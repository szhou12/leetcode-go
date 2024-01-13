# [2953. Count Complete Substrings](https://leetcode.com/problems/count-complete-substrings/description/)

## Solution idea
### Sliding Window: Flex + Fix
1. 同时满足两个条件的滑窗："挑软柿子捏"
    * 条件2比较好解决 - 每个相邻的字符的差的绝对值不超过2。先以条件2为基准，把整个字符串分割成若干个区间。
    * 使用 Sliding Window Flex 模版。
    * 每个区间以内再找符合条件1的子区间个数。
2. 隐含条件：输入的字符串只包含26个英文小写字母。
    * 与条件1结合：一个区间内，若有1个字母，那么符合条件的子区间的长度必为 1*k；若有2个字母，那么符合条件的子区间的长度必为 2*k；若有3个字母，那么符合条件的子区间的长度必为 3*k；...；若有26个字母，那么符合条件的子区间的长度必为 26*k。
    * 所以，每个区间先统计拥有不同字母的种类。假设有 T 种，那么满足条件1的子区间长度可以是：1*k, 2*k, 3*k, ..., T*k。
    * 每个子区间长度 (1*k, 2*k, 3*k, ..., T*k) 依次使用 Sliding Window Fix 模版的滑窗。
    * 每滑动一次，统计滑窗内字母的频次。若存在一个字母的频次不是k次，那么此时的子区间不满足条件1；反之，找到一个满足条件1的子区间。

Time complexity = $O(n*26*26)$

Explanation: The whole array is splited into `m` chunks (each chunk has max size of `n` while the total size of all chunks together is `n`). Each chunk iterates at most 26 times of window size while each window size iterates at most 26 times to check each character's frequency.

## Recource
[【每日一题】LeetCode 2953. Count Complete Substrings](https://www.youtube.com/watch?v=4DYbP4shsno&ab_channel=HuifengGuan)