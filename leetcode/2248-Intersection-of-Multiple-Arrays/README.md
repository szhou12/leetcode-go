# [2248. Intersection of Multiple Arrays](https://leetcode.com/problems/intersection-of-multiple-arrays/description/)

## Solution idea
### Hash Map + Frequency count
1. Keep a count of the number of times each integer occurs in `nums`.
2. Since all integers of `nums[i]` are distinct, if an integer is present in each array, its count will be equal to the total number of arrays.