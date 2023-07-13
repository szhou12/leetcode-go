# [2730. Find the Longest Semi-Repetitive Substring](https://leetcode.com/problems/find-the-longest-semi-repetitive-substring/description/)

## Solution idea
### Sliding Window (Flex)
#### 思路总结
1. 理解题意: semi-repetitive == at most one consecutive pair of the same digits == 子串中最多只能有一对儿挨着的相同字符
2. 我的解法: 用一个 boolean flag `pairFound` 来标记是否滑窗内已经存在 a consecutive pair of the same digits
3. 吃的逻辑: 如果 right 未出界:
    1. 如果 `pairFound == true` 并且 即将要吃的元素与前一个元素相同，则停止吃/延伸窗口右边界
    2. 如果 `pairFound == false` 并且 即将要吃的元素与前一个元素相同，则 `pairFound = true`，并可以继续吃/延伸窗口右边界
4. 吐的逻辑: 如果 当前left与下一个left的元素相同，更新 `pairFound = false`
5. 标准答案: 
    * 用 `count` 记录窗口内 consecutive pairs 的个数。按题目要求，`count` 始终要小于2
    * 继续吃/延伸右边界的条件: 即将要吃的元素给`count`的贡献使`count`总数仍然小于2

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 2730. Find the Longest Semi-Repetitive Substring](https://www.youtube.com/watch?v=NFS-vlKWO_o&ab_channel=HuifengGuan)