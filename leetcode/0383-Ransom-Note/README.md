# [383. Ransom Note](https://leetcode.com/problems/ransom-note/)

## Solution idea

* 用 HashMap 统计 `magazine`中每个字母出现的个数，然后在看够不够`ransomNote`来使用

* 思路和 [349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/) 一致

* **优化**: 用长度为26的数组 而不是 HashMap 来统计

Time complexity = $O(n)$