# [689. Maximum Sum of 3 Non-Overlapping Subarrays](https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/description/)

## Solution idea

### Sliding Window - 3-pass

* **Key Idea**: 利用题目给的条件 **3 Non-Overlapping Subarrays**. 想一想如果是 **2 Non-Overlapping Subarrays**的情况下怎么处理？ 解决方法是找一个断点使得左半段与右半段给出最优解. 此题同理, 区别在于此题的"断点"实际上是一段subarray.

* **具体实现: Sliding Window**
把nums分成 L, M, R三段:
```
{X X X X X }   X X X   {X  X X X X }
         i-1   i       i+k
```
使用两个`DP`思想的arrays: `leftMax`, `rightMax`:

**物理意义**:

`leftMax[i]`: the largest k-length subarray sum from `[0:i]`

`rightMax[i]`: the largest k-length subarray sum from `[i:n-1]`



Time complexity = $O(3n)$ = $O(n)$


## Resource
[【每日一题】LeetCode 689. Maximum Sum of 3 Non-Overlapping Subarrays](https://www.youtube.com/watch?v=oo_T4GdRot4&t=1705s&ab_channel=HuifengGuan)