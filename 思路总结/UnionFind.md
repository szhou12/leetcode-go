# Union Find

## 参考资料
* [Union-Find 算法详解](https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/UnionFind%E7%AE%97%E6%B3%95%E8%AF%A6%E8%A7%A3.md)

* 并查集可以解决什么问题?
    1. 主要就是集合问题，判断两个节点在不在一个集合
    2. 也可以将两个节点添加到一个集合中

* 并查集主要有三个功能:
    1. 找祖宗: 寻找当前节点的根节点，函数：`Find(x int)`，也就是判断这个节点的祖先节点是哪个
    2. 两家联姻: 将两个节点接入到同一个集合，函数：`Union(u int, v int)`，将两个节点连在同一个根节点上
        * 谁小谁成为祖宗 (目的: 维护一个顺序)
    3. 同宗同祖: 判断两个节点是否在同一个集合，函数：`Same(u int, v int)`，就是判断两个节点是不是同一个根节点

## 基础题
* [684. Redundant Connection](https://leetcode.com/problems/redundant-connection/)
* [1971. Find if Path Exists in Graph](https://leetcode.com/problems/find-if-path-exists-in-graph/)