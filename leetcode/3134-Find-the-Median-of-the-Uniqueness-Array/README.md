# [3134. Find the Median of the Uniqueness Array](https://leetcode.com/problems/find-the-median-of-the-uniqueness-array/description/)

## Solution idea
### Binary Search (Guess K) + Sliding Window
1. 看到涉及subarray首先想到的是DP，但是constraint给的是`1 <= nums.length <= 10^5`。一个长度为n的数组，subarray的总数是$\frac{n(n+1)}{2}$，这个数量级是$O(n^2)$，先罗列出所有subarrays再去找中位数是不现实的。DP的思想需要罗列所有subarrays，所以不适合。
2. 根据constraint，算法的时间复杂度应该在$O(n\log n)$或者$O(n)$。适合的算法一般是binary search 或者 greedy (巧了，本题两个都涉及到了)。
3. 题目求median，相当于暗示一个有序数组，一个有序数组暗示了使用binary search。搜索域是什么？根据题目给出的例子，把distinct subarray以不同元素个数从小打到排列: e.g., `[1, 1, 1, 2, 2, 3]`. 使用binary search猜答案，左右边界定义为不同元素的个数，左边界是1，右边界是n。收敛条件：每次猜一个mid，计算number of subarrays with at most mid distinct elements。
4. 如何计算number of subarrays with at most mid distinct elements? 使用sliding window。解法参考[0992](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0992-Subarrays-with-K-Different-Integers)

Time complexity = $O(n\log n)$

## Resource
[【每日一题】LeetCode 3134. Find the Median of the Uniqueness Array](https://www.youtube.com/watch?v=AHwygCfqyGo&ab_channel=HuifengGuan)