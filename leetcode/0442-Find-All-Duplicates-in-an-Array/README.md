# [442. Find All Duplicates in an Array](https://leetcode.com/problems/find-all-duplicates-in-an-array/)

## Solution idea

思路与 [448. Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/) 一致： 把元素与index关联起来, 出现一个元素，就在以它作为index的位置做标记 ($* -1$), 再遇到负数，就说明这个元素出现了两次，是我们要找的。

Time complexity = $O(n)$