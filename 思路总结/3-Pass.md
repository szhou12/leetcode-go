# 3-Pass

* *此种解法有点像 Dynamic Programming, 又有点像 Greedy Algorithm, 实现又常涉及 Sliding Window*

* **重点1**: 要清楚明白 从左往右扫和从右往左扫 时候, 分别构造的array的物理意义

* **重点2**: 分界点怎么定义(它可以是一个元素, 也可以是一段subarray)? 分界点左手边需要什么？分界点右手边需要什么? 左右两边怎么 concatenate?\

* **破题着眼点** - 如何下手这类题有两个着眼点 (如果第1个分析起来没有头绪就换第2个):
    1. 直接定义 **从左往右扫构造的array 和 从右往左扫构造的array** 的物理意义
    2. 定义 **分界点** 的物理意义开始。一般把分界点定义为题目所求, 然后中心开花向两边拓展

## 求总数/Count类型题

* 接雨水: [42. Trapping Rain Water](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0042-Trapping-Rain-Water)
    * 入手着眼点: **分界点** 的物理意义

* 移除一个元素使偶数位的和与奇数位的和相等: [1664. Ways to Make a Fair Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1664-Ways-to-Make-a-Fair-Array)

* L段和R段不同字符数相等的分界点个数: [1525. Number of Good Ways to Split a String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1525-Number-of-Good-Ways-to-Split-a-String)
    * **3-Pass + HashMap**

* 求两个字符串只差一个字符的substring匹配个数: [1638. Count Substrings That Differ by One Character](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1638-Count-Substrings-That-Differ-by-One-Character)
    * **3-Pass + 2D DP(双序列型的DP)**
    * 入手着眼点: **分界点** 的物理意义 
        * 分界点定义为substring中differ的那个char, 看两边保持元素对应相同的条件下最远可以走多远

* 所有subarray的最小值的和: [907. Sum of Subarray Minimums](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0907-Sum-of-Subarray-Minimums)
    * **3-Pass + Stack (find next smaller element)**
    * 入手着眼点: 从 **分界点** 的物理意义 出发

* 是否存在递增的triplet: [334. Increasing Triplet Subsequence](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0334-Increasing-Triplet-Subsequence)


## 求Max or Min/优化类型题

* 找三个不重叠k-长subarray的最大和的index: [689. Maximum Sum of 3 Non-Overlapping Subarrays]()

* 找最小可裁定/删除的subarray长度: [2565. Subsequence With the Minimum Score]()
    * **3-Pass + Binary Search 猜答案**

* 最大化array上任意两个segments: [2555. Maximize Win From Two Segments]()
    * two intervals --> 找分界点

* 移动盒子里的球分几步-求最小步数: [1769. Minimum Number of Operations to Move All Balls to Each Box]()

* 0/1交替出现所需最小翻转次数: [1888. Minimum Number of Flips to Make the Binary String Alternating]()

* 转变成mountain array的最小操作: [1671. Minimum Number of Removals to Make Mountain Array]()
    * **3-Pass + 1D DP(LIS)**

* 移除指定元素的最少时间: [2167. Minimum Time to Remove All Cars Containing Illegal Goods]()

* L段Sum - R段Sum 最小值: [2163. Minimum Difference in Sums After Removal of Elements]()
    * **3-Pass + Priority Queue**

* 构造ab形式字符串的最少操作数: [1653. Minimum Deletions to Make String Balanced]()
