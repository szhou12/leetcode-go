# [3213. Construct String with Minimum Cost](https://leetcode.com/problems/construct-string-with-minimum-cost/description/)

## Solution idea
### DFS + Memoization + Trie
```
 already chopped | about to chop (recusion)
    [X   X  X  X] {X  X  X | X  X}
                  cur    i
```
1. Trie: 在一个包含多个单词的字典里面找匹配substring `target[cur:i]` (inclusive) 的单词，容易想到使用 Trie 来储存这个字典。这样，就可以线性检查 (a在不在, ab在不在, abc在不在...)substring在不在字典里面。如果不在就不用再往下找了。
2. DFS: 怎么写DFS? Given a prefix `target[0:cur-1]` (假定这个前缀存在于Trie, 或者说已经“切好”了)，recursively DFS on `target[cur:i], cur <= i < n`, checking if `target[cur:i]` exists in the Trie (相当于把`target[cur:i]`切出来), if so, record cost and continue DFS on `target[i+1:n]` (切剩下的); otherwise, return `INF` (因为这个分支"切"不出来)。
3. Memo: DFS的好伙伴。记忆“未来的”结果 - 每次 往下DFS on `target[cur:n-1]`之前，如果发现以前已经处理过 (memo中已经记录了值)，就不用再重复DFS了。所以, memo的定义是：`memo[i] := min cost to 'chop' target[i:n-1]`
4. **注意！**这个解法无法AC `target="aaaaaaaa", words={"a", "aa", "aaa"}`的case，因为可以切的地方太多了。



## Resource
[【每日一题】LeetCode 3213. Construct String with Minimum Cost](https://www.youtube.com/watch?v=xiJ1fuN8Yfc&ab_channel=HuifengGuan)