# [3228. Maximum Number of Operations to Move Ones to the End](https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end/description/)

## Solution idea
1. 题目要求“最大化”操作数。通过观察，发现每遇到一个0，前面所有的1都会pass它。所以，我们可以在遍历过程中记录看到的1的数量，每见到一个0，操作数就是当前看过的1的数量。
2. 注意：当出现连续0时，1会一次性“滑”过去，只算作一次操作。所以，需要“跳过”连续的0

Time complexity = $O(n)$


## Resource
[Greedy Solution](https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end/solutions/5508962/greedy/)

[Intuition + Approach || T.C ~ O(N) || S.C ~ O(1)](https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end/solutions/5509105/intuition-approach-t-c-o-n-s-c-o-1/)