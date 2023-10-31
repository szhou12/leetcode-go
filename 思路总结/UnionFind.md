# Union Find

## 目录
* [知识点 Quick Walkthrough](#star-知识点-quick-walkthrough)
* [基础题](#基础题)
* [MST](#minimum-spanning-tree-mst)

## :star: 知识点 Quick Walkthrough
### 参考资料
* [Union-Find 算法详解](https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/UnionFind%E7%AE%97%E6%B3%95%E8%AF%A6%E8%A7%A3.md)
* [通俗易懂地讲解《并查集》](https://zhuanlan.zhihu.com/p/125604577)

### :bulb: 并查集可以解决什么问题?
1. 主要就是集合问题，判断两个节点在不在一个集合
2. 也可以将两个节点添加到一个集合中
3. 找连通: 相互连通的节点属于同一个集合, 相互连通的节点添加到同一个集合

### :bulb: 并查集主要有三个功能:
1. 找祖宗: 寻找当前节点的根节点，函数：`Find(x int)`，也就是找到当前节点的祖先节点
2. 两家联姻: 将两个节点接入到同一个集合，函数：`Union(u int, v int)`，将两个节点连在同一个根节点上
    * 谁小谁成为祖宗 (目的: 维护一个顺序)
3. 同宗同祖: 判断两个节点是否在同一个集合，函数：`Same(u int, v int)`，就是判断两个节点是不是同一个根节点
    * 可以用 `uf.Find(u) == uf.Find(v)` 代替

### :mag: 小技巧
#### 二维坐标节点的处理
* UnionFind 模版储存 `int`，那么如果题目中节点是二维的，怎么处理?
    1. 压缩为 1-D: 如果是在**矩阵**中，例如，一个节点表示为 `node=(i, j)`，则可以压缩为1-D储存在UnionFind中: `node = i*n + j`，其中`n = len(matrix[0])`是列的个数。
        * 例题: 1632, 959, 200, 947 (直角坐标系但当作矩阵处理，因为知道横纵坐标上限并且节点都在第一象限)
    2. 使用节点的index: 如果是在**直角坐标系**中，例如，一个节点i表示为 `node[i] = (x_i, y_i)`，则可以直接使用index `i` 储存在UnionFind中。
        * 例题: 1584

## 基础题

* :red_circle: :secret: 连通共享公约数的数字: [2709. Greatest Common Divisor Traversal](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2709-Greatest-Common-Divisor-Traversal)
    * Union Find + Prime Factorization
    * 时间优化: 用 共享的 prime factor(质因数) 作为 Union 的媒介/桥梁

* :green_circle: **起点到终点是否有通路:** [1971. Find if Path Exists in Graph](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1971-Find-if-Path-Exists-in-Graph)

* :red_circle: **从起点到终点的所有路径中最短的一条边:** [2492. Minimum Score of a Path Between Two Cities](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2492-Minimum-Score-of-a-Path-Between-Two-Cities)

* :red_circle: :secret: **字符串家族:** [2157. Groups of Strings](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2157-Groups-of-Strings)
	* Union Find + Bit Mask (encode string to 26-len bit representation)

* :red_circle: :secret: **寻找所有知晓秘密之人:** [2092. Find All People With Secret](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2092-Find-All-People-With-Secret)
    * Union Find + Two Pointers (相向而行)

* :red_circle: **谁和谁能成为朋友:** [2076. Process Restricted Friend Requests](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2076-Process-Restricted-Friend-Requests)

* :red_circle: **最小化Hamming Distance:** [1722. Minimize Hamming Distance After Swap Operations](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1722-Minimize-Hamming-Distance-After-Swap-Operations)

* :red_circle: :secret: **矩阵以rank表示:** [1632. Rank Transform of a Matrix](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1632-Rank-Transform-of-a-Matrix)
    * Topological Sort + Union Find

* :yellow_circle: **通过交换可以把字符串按照最小字典序排序:** [1202. Smallest String With Swaps](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1202-Smallest-String-With-Swaps)

* :green_circle: **字典序最小的等价字符串:** [1061. Lexicographically Smallest Equivalent String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1061-Lexicographically-Smallest-Equivalent-String)

* :green_circle: **等式的可满足性:** [990. Satisfiability of Equality Equations](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0990-Satisfiability-of-Equality-Equations)

* :red_circle: **分割矩阵能分几份:** [959. Regions Cut By Slashes](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0959-Regions-Cut-By-Slashes)

* :yellow_circle: **二维平面上相消元素个数:** [947. Most Stones Removed with Same Row or Column](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0947-Most-Stones-Removed-with-Same-Row-or-Column)

* :red_circle: **判别二分图:** [785. Is Graph Bipartite?](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0785-Is-Graph-Bipartite)

* :red_circle: **情侣手牵手:** [765. Couples Holding Hands](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0765-Couples-Holding-Hands)

* :yellow_circle: **寻找并合并相同的账户:** [721. Accounts Merge](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0721-Accounts-Merge)
    * 修改 UnionFind 模版: `int` -> `string`

* :red_circle: **找出连通图中多余的边II:** [685. Redundant Connection II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0685-Redundant-Connection-II)
    * 有向图

* :yellow_circle: **找出连通图中多余的边:** [684. Redundant Connection](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0684-Redundant-Connection)
    * 无向图: 多余的边 就是 判断有没有构成环

* :green_circle: **连通图个数:** [547. Number of Provinces](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0547-Number-of-Provinces)

* :yellow_circle: **连通岛屿的个数:** [200. Number of Islands](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0200-Number-of-Islands)
    * 1-D Index Compression

* :red_circle: **最长连续序列:** [128. Longest Consecutive Sequence](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0128-Longest-Consecutive-Sequence)

* :yellow_circle: **找出不包含边界元素的连通图:** [130. Surrounded Regions](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0130-Surrounded-Regions)

## Minimum Spanning Tree (MST)
* :red_circle: :secret: **寻找MST中的关键边和伪关键边:** [1489. Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1489-Find-Critical-and-Pseudo-Critical-Edges-in-Minimum-Spanning-Tree)
    * Union Find + Kruskal

* :red_circle: :secret: **生成连通图所需最少的边/可以删除最多多少边:** [1579. Remove Max Number of Edges to Keep Graph Fully Traversable](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1579-Remove-Max-Number-of-Edges-to-Keep-Graph-Fully-Traversable)
    * Union Find + Spanning Tree
    * 边没有权重，所以这道题只求 spanning tree 而不用 MST

* :yellow_circle: **连通所有二维节点的最小花费:** [1584. Min Cost to Connect All Points](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1584-Min-Cost-to-Connect-All-Points)
    * Union Find + Kruskal

* :green_circle: :lock: 最小花费连接所有城市: [1135. Connecting Cities With Minimum Cost](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1135-Connecting-Cities-With-Minimum-Cost)
    * Union Find + Kruskal

* :red_circle: :lock: 最小花费给村庄供给水源: [1168. Optimize Water Distribution in a Village](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1168-Optimize-Water-Distribution-in-a-Village)
    * Union Find + Kruskal
