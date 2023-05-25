# [1168. Optimize Water Distribution in a Village](https://leetcode.ca/2019-02-10-1168-Optimize-Water-Distribution-in-a-Village/)

## Solution idea
### MST = Union Find + Kruskal
#### 思路总结
1. 这道题的解法很巧妙，需要用到转化思想。
2. 注意到：每个节点有两种操作，一是自建井，二是从别的节点引水。两种操作看似不一样，其实底层逻辑是一样的，即，它们都等同于“从别处引水” (make a connection to a water source)。自建井 相当于 从一个虚拟的大水池节点引水，引水的费用就是建井的费用。
3. 这样，我们就把两种操作都转化成同类型的操作了，即，从别处引水。这样，我们就可以用 Kruskal 算法来求解了。
4. 实际操作上，只需要新增一个 node 0 (虚拟的大水池节点)，把每个节点自建井的费用wells[i]想象成连接[0, i]的边的费用。
5. 然后使用 Kruskal 算法求解即可。

Time complexity = $O(E\log E)$

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Union_Find/1168.Optimize-Water-Distribution-in-a-Village)