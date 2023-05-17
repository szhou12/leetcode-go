# [721. Accounts Merge](https://leetcode.com/problems/accounts-merge/description/)

## Solution idea
### Union Find
#### 思路总结
1. 一开始我的想法是：两两account比较，如果有相同的email，就 union account index，这样，可以继续使用 UnionFind模版。缺点是：两两比较不够efficient，同时，需要多个map记录index到email的映射等，倒来倒去容易出错
    * UnionFind 里存储 account index (`int`)
2. 参考答案：实际上，不需要两两account比较，loop每一个account时，把他的所有email union起来就够了 (i.e. 每两个相邻邮箱的归并)。这样，如果下一个account存在相同的email时，它会负责把这两个account union起来。归并完之后，再遍历一次所有的邮箱，按照其Father分类就是答案所要求的分类。每个分类对应的人名，就是Father account对应的人名。
    * UnionFind 里存储 email address (`string`)
    * UnionFind 模版需要稍加修改：把 `int` 改成 `string` 即可

Time complexity = $O(n*m)$ where n = # of accounts, m = # of emails per account
## Resource
[【每日一题】721. Accounts Merge, 10/03/2019](https://www.youtube.com/watch?v=SaDPCgT-EbQ&ab_channel=HuifengGuan)
