# [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)

## Solution idea

### Two Pointers

* **思路:** 求 Subarray 只需要确定左、右边界，左右边界圈出来的长度就是Subarray的个数
* **如何确定左右边界？**
    * 固定一个边界，另一个边界不停探索至极限
    * 到达极限后，收缩固定的那个边界，另一个边界再不停探索至极限
    * 循环往复，直至出界
* 落实到这道题, 比较intuitive的做法:
    * 固定左边界，右边界不停往前推至极限 (ie. product >= k)
        * 未到极限的条件: 右边界还未出界 **并且** 累乘上当前右边界的值还未超出要求
    * 到达极限后，收缩左边界(右移一位)，右边界再不停往前推至极限
* **一个重要的细节**
    * slow指针会超过fast指针的情况 (根据物理意义, slow指针要永远 $ \leq $ fast指针)
    * 发生这种状况的例子: 
        * [10, 2, 1], k = 4
        * slow & fast一开始都指向10
        * 按照上述逻辑, fast目前已至极限, slow右移一位, 那么此时, slow指向2, fast还指向10; 而且因为收缩slow, 砍掉slow之前指向的值时会导致 product = 0
        * 解决办法：
            * fast开始探索前, 先检查是否此时slow超越了fast
            * 如果是, fast跳到slow处，并且product重置为初始值1


Time complexity = $O(n)$

## Resource

[【每日一题】713. Subarray Product Less Than K, 9/19/2020](https://www.youtube.com/watch?v=WOSWdl4Fl00&ab_channel=HuifengGuan)