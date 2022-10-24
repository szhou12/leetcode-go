# Two Pointers

## Subarray 类型

* 一般像是求满足条件的subarray个数

* 总体思路:
    * 求 Subarray 只需要确定左、右边界，左右边界圈出来的长度就是Subarray的个数

* 具体实现的思路:
    * 如何确定左右边界？
    * 用快慢指针
    * 固定一个边界，另一个边界不停探索至极限
    * 到达极限后，收缩固定的那个边界，另一个边界再不停探索至极限
    * 循环往复，直至出界

* 求满足条件的Subarray个数: [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)
    * 固定左边界，右边界不停探索至极限
    * 到达极限后，收缩左边界(右移一位)，右边界再探索至极限
    * **注意！！！** 探索前先检查 slow超过fast的情况