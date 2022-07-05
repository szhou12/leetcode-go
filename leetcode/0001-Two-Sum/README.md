# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## Solution idea

Use `map` whose `key` stores array elements and `value` stores corresponding index.

Loop through the input array: if the complement (`target - num`) of the current `num` exists in the `map`, return their indices.

Otherwise, store current element and its index into the map.

Time complexity = $O(n)$