# [58. Length of Last Word](https://leetcode.com/problems/length-of-last-word/)

## Solution idea

首先，**从后往前**扫，找到last word's last character;

然后，**从last character出发往前**扫，找到 empty space just right before last word's first character;

last word长度等于两者相减.

Time complexity = $O(n)$

## Resources
[58. Length of Last Word](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0058.Length-of-Last-Word)