# [2977. Minimum Cost to Convert String II](https://leetcode.com/problems/minimum-cost-to-convert-string-ii/description/)

## Solution idea
### Floyd-Warshall + DP + Trie
1. Floyd算法: `original[]`, `changed[]`, `cost[]` 部分与[2976. Minimum Cost to Convert String I](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2976-Minimum-Cost-to-Convert-String-I)一样。差别在于，本题的 node 代表一个字符串，而不是一个字符。需要先对字符串进行“离散化”，也就是，每一个字符串赋予一个整数ID。实现上，用 HashMap 映射 字符串 到 ID。
2. DP: 一维区间型DP
    * Definition: `dp[i] = ` minimum cost to convert `source[0:i]` to `target[0:i]`
    * Base case: `dp[0] = 0` meaning there is no cost to convert an empty string to another empty string
    * Recurrence: `dp[i] = min(dp[j-1] + dist[j][i])` for all `j` in `[1, i]` where `dist[j][i]` is the shortest path from `source[j:i]` (inclusive) (represented as node id `j`) to `target[j:i]` (inclusive) (represented as node id `i`) in the Floyd graph.
3. Trie: “高效查询是否一个字符串存在于一个字符串池中”
    * 如何判断 字符串 `source[j:i]` 和 `target[j:i]` 存在于 Floyd graph nodes (构成字符串池) 中？
        * TLE Solution: HashMap 构造 字符串池, 但是本题会造成 TLE
        * Optimized Solution: Trie 构造 字符串池。注意，为了进一步优化，`j`的loop方向改为逆向: `j := i -> 1` (这样做的结果是，Trie中查询是按照字符串从右往左的顺序)；所以，为保持顺序一致，在一开始，提前reverse `original[]`, `changed[]`所有字符串。


Time complexity = $O(V\cdot h + n^3 + n^2 \cdot h)$
## Resource
[【每日一题】LeetCode 2977. Minimum Cost to Convert String II](https://www.youtube.com/watch?v=pQ_gRovgx70&ab_channel=HuifengGuan)