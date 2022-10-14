# [406. Queue Reconstruction by Height](https://leetcode.com/problems/queue-reconstruction-by-height/)

## Solution idea

### Greedy Algorithm

* **Key**: 本题有两个维度，h和k，看到这种题目一定要想如何**先确定一个维度**，然后再按照另一个维度重新排列。

* 先按h排序呢，还是先按照k排序呢？
    * 如果按照k来从小到大排序，排完之后，会发现k的排列并不符合条件，身高也不符合条件，两个维度哪一个都没确定下来。
    * 那么按照身高h来排序呢，身高一定是从大到小排（身高相同的话则k小的站前面），让高个子在前面。

* 按照身高排序之后，优先按身高高的people的k来插入，后序插入节点也不会影响前面已经插入的节点，最终按照k的规则完成了队列。

* **局部最优可推出全局最优**
    * 局部最优：优先按身高高的people的k来插入。插入操作过后的people满足队列属性
    * 全局最优：最后都做完插入操作，整个队列满足题目队列属性

* 刷题或者面试的时候，手动模拟一下感觉**可以局部最优推出整体最优，而且想不到反例**，那么就试一试**贪心**

Time complexity = $O(n\log n + n^2)$

* Similar Question: [135. Candy](https://leetcode.com/problems/candy/)

## Resource

[代码随想录-406.根据身高重建队列](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0406.%E6%A0%B9%E6%8D%AE%E8%BA%AB%E9%AB%98%E9%87%8D%E5%BB%BA%E9%98%9F%E5%88%97.md)