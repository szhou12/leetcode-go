# Union Find

## 目录
* [知识点 Quick Walkthrough](#知识点-quick-walkthrough)
* [基础题](#基础题)

## 知识点 Quick Walkthrough
### 参考资料
* [Union-Find 算法详解](https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/UnionFind%E7%AE%97%E6%B3%95%E8%AF%A6%E8%A7%A3.md)
* [通俗易懂地讲解《并查集》](https://zhuanlan.zhihu.com/p/125604577)

### 并查集可以解决什么问题?
1. 主要就是集合问题，判断两个节点在不在一个集合
2. 也可以将两个节点添加到一个集合中
3. 找连通: 相互连通的节点属于同一个集合, 相互连通的节点添加到同一个集合

### 并查集主要有三个功能:
1. 找祖宗: 寻找当前节点的根节点，函数：`Find(x int)`，也就是找到当前节点的祖先节点
2. 两家联姻: 将两个节点接入到同一个集合，函数：`Union(u int, v int)`，将两个节点连在同一个根节点上
    * 谁小谁成为祖宗 (目的: 维护一个顺序)
3. 同宗同祖: 判断两个节点是否在同一个集合，函数：`Same(u int, v int)`，就是判断两个节点是不是同一个根节点

## 基础题
* [1971. Find if Path Exists in Graph](https://leetcode.com/problems/find-if-path-exists-in-graph/)

* :red_circle: 从起点到终点的所有路径中最短的一条边: [2492. Minimum Score of a Path Between Two Cities](https://leetcode.com/problems/minimum-score-of-a-path-between-two-cities/description/)

* :red_circle: :secret: 字符串家族: [2157. Groups of Strings](https://leetcode.com/problems/groups-of-strings/description/)
	* Union Find + Bit Mask (encode string to 26-len bit representation)

* :red_circle: :secret: 寻找所有知晓秘密之人: [2092. Find All People With Secret](https://leetcode.com/problems/find-all-people-with-secret/description/)
    * Union Find + Two Pointers (相向而行)

* :red_circle: 谁和谁能成为朋友: [2076. Process Restricted Friend Requests](https://leetcode.com/problems/process-restricted-friend-requests/description/)

* :red_circle: 最小化Hamming Distance: [1722. Minimize Hamming Distance After Swap Operations](https://leetcode.com/problems/minimize-hamming-distance-after-swap-operations/description/)

* :red_circle: :secret: 矩阵以rank表示: [1632. Rank Transform of a Matrix](https://leetcode.com/problems/rank-transform-of-a-matrix/description/)
    * Topological Sort + Union Find

* :yellow_circle: 通过交换可以把字符串按照最小字典序排序: [1202. Smallest String With Swaps](https://leetcode.com/problems/smallest-string-with-swaps/description/)

* :green_circle: 字典序最小的等价字符串: [1061. Lexicographically Smallest Equivalent String](https://leetcode.com/problems/lexicographically-smallest-equivalent-string/description/)

* :green_circle: 等式的可满足性: [990. Satisfiability of Equality Equations](https://leetcode.com/problems/satisfiability-of-equality-equations/description/)

* :red_circle: 分割矩阵能分几份: [959. Regions Cut By Slashes](https://leetcode.com/problems/regions-cut-by-slashes/)

* :yellow_circle: 二维平面上相消元素个数: [947. Most Stones Removed with Same Row or Column](https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/description/)

* :red_circle: 判别二分图: [785. Is Graph Bipartite?](https://leetcode.com/problems/is-graph-bipartite/)

* :red_circle: 情侣手牵手: [765. Couples Holding Hands](https://leetcode.com/problems/couples-holding-hands/description/)

* :yellow_circle: 寻找并合并相同的账户: [721. Accounts Merge](https://leetcode.com/problems/accounts-merge/description/)
    * 修改 UnionFind 模版: `int` -> `string`

* :red_circle: 找出连通图中多余的边II: [685. Redundant Connection II](https://leetcode.com/problems/redundant-connection-ii/)
    * 有向图

* :yellow_circle: 找出连通图中多余的边: [684. Redundant Connection](https://leetcode.com/problems/redundant-connection/)
    * 无向图: 多余的边 就是 判断有没有构成环

* :green_circle: 连通图个数: [547. Number of Provinces](https://leetcode.com/problems/number-of-provinces/description/)

* :yellow_circle: 连通岛屿的个数: [200. Number of Islands](https://leetcode.com/problems/number-of-islands/submissions/)
    * 1-D Index Compression

* :red_circle: 最长连续序列: [128. Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/description/)