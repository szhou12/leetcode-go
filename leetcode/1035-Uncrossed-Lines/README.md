# [1035. Uncrossed Lines](https://leetcode.com/problems/uncrossed-lines/description/)

## Solution idea

### DP - Longest Common Subsequence 类型题

* 这是一道可以直接套用 LCS 模版的题

* **难点**在于如何从题目中分析出这是一道 LCS 典型题?
    * 题目要求：绘制一些连接两个数字 `A[i]` 和 `B[j]` 的直线，使得 `A[i] == B[j]`，且**直线不能相交**！
    * **直线不能相交**，这就是说明在字符串A中 找到一个与字符串B相同的子序列，且这个子序列不能改变相对顺序，只要**相对顺序不改变**，链接相同数字的直线就不会相交。

* 例子: `A = [1,4,2]`, `B = [1,2,4]`
    * A和B的最长公共子序列是[1,4]，长度为2
    * 公共子序列的特征就是**相对顺序不变**, 意即, 数字4在字符串A中数字1的后面，那么数字4也应该在字符串B数字1的后面

* **总结**: 题目要求的 1. `A[i] == B[j]`; 2. **直线不能相交** 两个条件刚好就符合 LCS 的定义, 所以题目所求就可以转化为求 LCS

Time complexity = $O(m*n)$, Space complexity = $O(m*n)$

## Resource
[代码随想录-1035.不相交的线](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1035.%E4%B8%8D%E7%9B%B8%E4%BA%A4%E7%9A%84%E7%BA%BF.md)