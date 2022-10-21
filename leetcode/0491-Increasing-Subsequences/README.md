# [491. Increasing Subsequences](https://leetcode.com/problems/increasing-subsequences/)

## Solution idea

### DFS - 需要去重的All Subsets问题

* 注意！此题用All Subsets的第一种写法并不好写, 因为需要当前层去重
    * 第一种写法: 每层 # branches = 2 (加与不加); Base Case添加合法的当前子集
    * 第二种写法: 每层 # branches = 从 startIndex 到 最后一个, 依次需要试的元素; 合法的当前子集沿着 recursion tree 路径随时添加

* 每一层需要**新**生成一个 **Map** 去重, 跳过当前层加过的元素
    * 注意！层与层之间的相同元素**不能**去重

Time complexity = $O(2^n)$

## Resource
[代码随想录-491.递增子序列](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0491.%E9%80%92%E5%A2%9E%E5%AD%90%E5%BA%8F%E5%88%97.md)