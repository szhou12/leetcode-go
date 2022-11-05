# [763. Partition Labels](https://leetcode.com/problems/partition-labels/)

## Solution idea

### Greedy Algorithm

* **Key Idea**: 统计字符串中所有字符的起始和结束位置，记录这些区间，**将区间按左边界从小到大排序，找到边界将区间划分成组，互不重叠**。
    * Similar Questions: [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/), [452. Minimum Number of Arrows to Burst Balloons](https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/)

* **怎么想到的呢?**
    * 题目定义每种字符能且只能出现在一个partition里，也就是说，每个字符不能跨partition，也就是暗示了这是一道不能跨越区间的题
    * partition的定义转化为: 排序后, 所有overlapped intervals的集合, 也就是说, partition与partition之间互相不重叠

* **为什么这道题是按左边界增序排列呢？**
    * 这是因为sort by left end 可以迅速确认每个partition最小左边界 (i.e. 每个partition第一个遇到的interval), 这样, 只需要在loop中不断找partition的最大右边界就行。

* 物理意义:
    * `left`: 目前parition的最小左边界
    * `right`: 目前parition的最大右边界
    * 找到一个新partition的判断条件: 
        * 第i-th轮interval的左边界 > 目前parition的最大右边界
        * 说明之后所有intervals都不与目前parition重叠, 所以, 从第i-th轮interval起, 开始一个新的partition

* 实现细节:
    * 最后一个partition不会在loop中被添加，需要在loop结束后，额外添加最后一个partition

Time complexity = $O(n) + O(m\log m + m) = O(n + m\log m)$ where $n$ = length of input string, $m$ = number of distinct characters in the input string

## Resource

还有其他解法，但不容易想到且不好归纳类型题，目前不做研究，在参考链接中给出

[代码随想录-763.划分字母区间](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0763.%E5%88%92%E5%88%86%E5%AD%97%E6%AF%8D%E5%8C%BA%E9%97%B4.md)

[halfrost/LeetCode-Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0763.Partition-Labels)