# [1579. Remove Max Number of Edges to Keep Graph Fully Traversable](https://leetcode.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable/description/)

## Solution idea
### Union Find + Spanning Tree
#### 思路总结
1. 反思：一开始我的想法是，先以一个人为标准尝试连图，构造一棵生成树。e.g. 先看 只有Alice可以走 和 两个人都可以走的边 连接一个Alice可以走通的树。然后再以Bob为视角，看哪些点没有连接上 以及 哪些边是重复可以去掉的。但是，这样做需要记录哪些边是Alice 用来连接图的，哪些又是候选项，并不方便。例如，如果两个点之间可以用两人都可以走的边连接 又可以用Alice可以走的边连接，这种情况下，我们保留那一个？显然，我们应该优先选择 用两人都可以走的边连接。但是，过一遍所有的边并不能保证我们能够优先选到两人都可以走的边。所以，这个想法太过于"粘腻" - 要先完成一个人的图，再在这个图的基础上做修改，来完成另一个人的图。为什么不独立开来，各干各的？
2. 参考答案：
    * 两个人都能走的边 是“万用”的，先找出这些边，并先用这些边连接节点。此时完成的“半成品图”作为“地基”分别为Alice和Bob各自的构造生成树打底。
        * "可以想象，这一步操作之后整个图变成了若干个割裂的联通块，但是每个联通块内部是最优联通的（也就是用了最少的边）。每个联通块内部，是可以让两个人都自由访问到的。"
    * 在"地基"的基础上，以ALice的视角，看还有哪些节点是还没被连接的，需要用Alice可以走的边来补上。“地基”的边+补上的边 = Alice的生成树必须要有的边。
        * 如果此时边的总数还不到 n - 1, 那就无法完成Alice的生成树。
    * 在"地基"的基础上，以Bob的视角，看还有哪些节点是还没被连接的，需要用Bob可以走的边来补上。“地基”的边+补上的边 = Bob的生成树必须要有的边。
        * 如果此时边的总数还不到 n - 1, 那就无法完成Bob的生成树。
    * 可以删掉的边 = 所有的边 - Alice的生成树必须要有的边 - Bob的生成树必须要有的边 

Time complexity = $O(n)$ where n = number of edges

## Resource
[【每日一题】1579. Remove Max Number of Edges to Keep Graph Fully Traversable, 9/9/2020](https://www.youtube.com/watch?v=AplZXP_LxgU&ab_channel=HuifengGuan)