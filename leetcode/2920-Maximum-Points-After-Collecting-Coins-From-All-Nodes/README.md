# [2920. Maximum Points After Collecting Coins From All Nodes](https://leetcode.com/problems/maximum-points-after-collecting-coins-from-all-nodes/description/)

## Solution idea
### DFS + Memoization
1. 破题：本题容易看出是一道 tree traversal 的题目，这也就意味着可以使用 Recursion (DFS) 来递归解决。
2. 细节:
    1. “砍半”的操作会影响到后代。也就是，如果当前节点决定进行“砍半”操作，那么它的所有后代节点都要进行一次“砍半”。如何累计一个节点在进行操作前，它需要先“砍半”多少次？可以使用一个变量 `times` 来累计传到当前节点时，它的祖先们进行了多少次的“砍半”操作。
    2. 使用 DFS 会造成重复计算。避免重复计算的一个常用方法是 Memoization。如何构造这个 `memo`？注意到，`dfs()` 函数里递归时，变量是 `cur` 和 `times`，因此如此定义：`memo[x][t] := max points that can be collected from subtree rooted at x, with t cut-in-half operations accumulated from x's ancestors`
        * 注意到题目的constraint里给到：`1 <= coins[i] <= 10^4`。这个暗示了，我们有一个最多“砍半”操作数：$13 < \log_{2} 10^4 < 14$。也就是说，一个节点“砍半”次数多于13次时，就是对0进行砍半，也就是没有任何操作。因此，`memo[x][t]`的第二维的长度只需要14。 

## Resource
[【每日一题】LeetCode 2920. Maximum Points After Collecting Coins From All Nodes](https://www.youtube.com/watch?v=ifMo3JfFinI&ab_channel=HuifengGuan)