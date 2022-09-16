# [448. Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/)

## Solution ideas

### Simple idea - HashSet
用一个额外的array来模拟hashset

### 关联元素与index
把每一个元素看作一个index，然后把出现在这个元素作为index的位置上的元素 作标记为 $* -1$。 最后扫一遍所有元素，没有成为负数的元素，说明这个index没有出现过，故加入result中。

Time complexity = $O(n)$

## Similar Problem
[442. Find All Duplicates in an Array](https://leetcode.com/problems/find-all-duplicates-in-an-array/)