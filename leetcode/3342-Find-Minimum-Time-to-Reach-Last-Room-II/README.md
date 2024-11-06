# [3342. Find Minimum Time to Reach Last Room II](https://leetcode.com/problems/find-minimum-time-to-reach-last-room-ii/description/)

## Solution idea
### Dijkstra
- **矩阵走格子类型题**。与[3341. Find Minimum Time to Reach Last Room I](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3341-Find-Minimum-Time-to-Reach-Last-Room-I)的思路基本一致。
- 唯一区别: **node储存一个额外状态**，即，步入下一个格子需要花费的时间：如果当前花费1s，下一个花费2s；如果当前花费2s，下一个花费1s.

Time complexity = $O(mn\log{mn})$