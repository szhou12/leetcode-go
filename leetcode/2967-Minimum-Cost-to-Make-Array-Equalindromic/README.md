# [2967. Minimum Cost to Make Array Equalindromic](https://leetcode.com/problems/minimum-cost-to-make-array-equalindromic/description/)

## Solution idea
### Math: Median Theorem + Closest Palindrome
1. 中位数定理：给定数轴上的n个点，找出一个到它们的距离之和尽量小的点。结论：这些点的中位数就是目标点。
2. 本题中，先给数组排序，再根据中位数定理，确定中位数 (注意：中位数不一定必须出现在数组中)。
3. Closest Palindrome: 找到与string表示的目标数值最接近的回文串。参考：[564. Find the Closest Palindrome](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0564-Find-the-Closest-Palindrome)
4. 本题中，与 564 稍微不同的是，不是严格定义的排除自己作为回文串，也就是本身可以是回文串。

Time complexity = $O(n\log n)$

## Resource
[【每日一题】LeetCode 2967. Minimum Cost to Make Array Equalindromic](https://www.youtube.com/watch?v=O2BNWCruFZw&t=639s&ab_channel=HuifengGuan)