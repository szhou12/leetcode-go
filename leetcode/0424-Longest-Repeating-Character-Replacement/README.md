# [424. Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/description/)

## Solution idea
### Sliding Window (Flex模版)
#### 思路总结

##### 解法一 (朴素解法)
1. 题目要求窗口内的字符全是同一个字母，朴素的解法是“挨个试”：把26个大写字母挨个当作窗口内的目标字母，滑一遍字符串，在变换次数的限制下，看以当前字母为窗口目标的最长窗口是多少。
2. 窗口为 左闭右闭 区间，即`[left, right]`
3. 右边界可延伸条件为：(变换次数 k > 0 OR 当前右边界所指元素是目标字母不需变换)
    * 注：一开始我写的时候只有 k > 0. 但发现 k = 0 的情况时发现这样无法使有连续目标字母的子串伸长。分析 k = 0 的情况帮我确认了这个条件。

Time complexity = $O(26n) = O(n)$

##### 解法二 (优雅解法)
1. 滑窗框住一段区间。计算区间内占多数的字母，其他占少数的字母就是需要被替换掉的。
2. 一个合法的窗口满足：`total window size L - Majority element A <= k`
3. 怎么找 Majority?
    * 转化为 记录频次。(aka. count[])


## Resource
[【每日一题】LeetCode 424. Longest Repeating Character Replacement 10/6 2021](https://www.youtube.com/watch?v=wXicFFUVdd0&ab_channel=HuifengGuan)