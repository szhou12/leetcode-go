# [395. Longest Substring with At Least K Repeating Characters](https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/)

## Solution idea

### Divide & Conquer

**从哪里divide?**

每次在一个区间内找到一个出现频次少于 k 的字母就是切分点

**当前层 干点什么事儿**

把整个当前区间遍历完，统计所有字符和它出现的频次，找到切分点

Time complexity = $O(n^2)$


### Sliding Window - 模版Flex

**Caution**: 这道题不能简单套用**模版Flex**, 需要额外添加一个条件防止右边界无限延伸 (无限延伸是由题目要求的条件导致的)

**模版Flex能work的两个必要条件**
1. sliding window 内 每个字母个数 `>= k` (题目要求)
2. sliding window 内 固定只框住 `m` 个不同的字母 (额外添加条件) ($m = 1 \cdots 26$)

Time complexity = $O(26n)$

## Resource

- 分治法/递归: [贾考博 LeetCode 395. Longest Substring with At Least K Repeating Characters](https://www.youtube.com/watch?v=Dn_Wt6lBt4I)

- [【每日一题】395. Longest Substring with At Least K Repeating Characters, 12/22/2020](https://www.youtube.com/watch?v=_MJKUvM-4fM&ab_channel=HuifengGuan)
