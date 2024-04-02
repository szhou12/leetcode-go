# [3097. Shortest Subarray With OR at Least K II](https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/description/)

## Solution idea
### Binary Search 猜答案 + Sliding Window + Bitwise OR Sum
1. 整体思路比较容易想出来：
    - 最外层是Binary Search猜答案: 左右边界分别是最短可能长度1和最长可能长度n, 然后尽可能往小了猜。如果一个长度mid能找到一个subarray的bitwise OR sum至少是k, 那么找到一个可能解。注意, 本题可能无解, 所以最后要检查收敛解是否满足条件。
    - 给定一个长度mid，如何高效判断是否存在满足条件的subarray? 容易想到使用定长的Sliding Window: 吃进右边元素, 检查滑窗内bitwise OR sum是否满足条件。如果满足，找到了，直接返回。否则，吐出左边元素。
2. Bitwise OR Sum Trick
    - 个人在本题的难点在如何计算bitwise OR sum，更准确的是，如何用bitwise操作吐出左边元素？
    - 小技巧: 
        - 利用OR的性质: 一旦遇到1，之后无论OR谁，都始终保持1
        - 具体实现: 使用一个`count[]`，长度为31。(因为Go中`int`为32位且最高位是signed bit，所以只考虑31位)。`count[]`的物理意义: 遍历滑窗内所有元素，记录每个bit位上1出现的次数。计算sum是遍历`count[]`，遇到非零bit位，就加上`1<<i`。吐的机制则是更新`count[]`: `count[b] -= (nums[left]<<b) & 1`
    
Time complexity = $O(n\log n)$

## Resource
[【每日一题】LeetCode 3097. Shortest Subarray With OR at Least K II](https://www.youtube.com/watch?v=be-rHbISs_E&ab_channel=HuifengGuan)