# [135. Candy](https://leetcode.com/problems/candy/)

## Solution idea

**Key 1**: 注意条件二: Children with a higher rating get more candies than their *neighbors*. 每个人有左、右两个邻居，所以要分别看高于左邻居时得到的糖果 和 高于右邻居时得到的糖果. 两者再取最大值以保证同时高于左、右两个邻居.

### Greedy Algorithm

* Step 1: 先确定右边评分大于左边的情况（从左向右遍历）
    * 此时局部最优：只要右边评分比左边大，右边的孩子就多一个糖果，全局最优：相邻的孩子中，评分高的右孩子获得比左边孩子更多的糖果局部最优可以推出全局最优
* Step 2: 再确定左孩子大于右孩子的情况（从右向左遍历）
    * 此时局部最优
* Step 3: 每个孩子再取最大糖果数以保证同时高于左、右两个邻居
    * 此时全局最优


Time complexity = $O(n)$

## Resource

[代码随想录-135. 分发糖果](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0135.%E5%88%86%E5%8F%91%E7%B3%96%E6%9E%9C.md)