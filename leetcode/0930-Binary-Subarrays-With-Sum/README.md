# [930. Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum/description/)

## Solution idea

### Sliding Window - 窗口长度可变

* 固定左边界, 右边界延伸至滑窗内的元素的sum恰好为goal. 此时如果知道右边界右边有k个连续的0，那么就意味着以该左边界、元素sum是goal的滑窗就有k+1个
* 对于每个元素，它后面有多少个连续的0，可以提前预处理得到
* 检查左边界**超过**右边界的情况:
    * 如果 goal=0, 就有可能 右边界延伸的loop 无法进入 (因为 sum 初始值也设为0), 这样 右边界无法延伸, 导致出现左边界**超过**右边界的情况
    * e.g. `[0,0,0], sum = 0, goal = 0`
    * Solution: 右边界延伸的loop添加条件: `left >= right` 维护 `[left, right)` 总是有一个元素被包括

Time complexity = $O(n)$

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Hash/930.Binary-Subarrays-With-Sum)