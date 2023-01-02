# [395. Longest Substring with At Least K Repeating Characters](https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/)

## Solution idea

### Divide & Conquer

**从哪里divide?**

每次在一个区间内找到一个出现频次少于 k 的字母就是切分点

**当前层 干点什么事儿**

把整个当前区间遍历完，统计所有字符和它出现的频次，找到切分点

Time complexity = $O(n^2)$


## Resource

- 分治法/递归: [贾考博 LeetCode 395. Longest Substring with At Least K Repeating Characters](https://www.youtube.com/watch?v=Dn_Wt6lBt4I)
