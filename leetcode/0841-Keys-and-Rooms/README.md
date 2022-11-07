# [841. Keys and Rooms](https://leetcode.com/problems/keys-and-rooms/description/)

## Solution idea

### DFS
* 题目要求只要能从`room 0`出发到达所有其他rooms就行
    * 不强求从`room 0`出发，必须遍历所有rooms，最后到达`room n-1`
* input graph可能有环，所以用 hashmap `visited` 记录遍历过的节点以避免回到遍历过的节点
    * 同时，`visited`记录了从`room 0`可以遍历到的所以rooms, 最后只用检查哪个`room`不在 `visited`中就可以决定返回true/false

Time complexity = $O(V+E)$

## Resource
[代码随想录-841.钥匙和房间](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0841.%E9%92%A5%E5%8C%99%E5%92%8C%E6%88%BF%E9%97%B4.md)