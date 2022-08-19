# [219. Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/)

## Solution idea

Use `HashMap` where the `key` stores each unique `num`, the `value` stores a struct `item`

Struct `item` contains two fields: `lastIndex` is the index that `num` appears last time; `distance` is the distance that distance between index values of `num` (updated only when the distance between the current index of `num` and `lastIndex` is smaller than the previously stored `distance`)

Time complexity = $O(n)$