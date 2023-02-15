# [689. Maximum Sum of 3 Non-Overlapping Subarrays](https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/description/)

## Solution idea

### Sliding Window - 3-pass

* **Key Idea**: 利用题目给的条件 **3 Non-Overlapping Subarrays**. 想一想如果是 **2 Non-Overlapping Subarrays**的情况下怎么处理？ 解决方法是找一个断点使得左半段与右半段给出最优解. 此题同理, 区别在于此题的"断点"实际上是一段subarray.
    * *wisdompeak*: 考虑到本题恰好只求三段区间，所以很容易想到three-pass的套路。我们在数组里遍历一段长度为k的滑窗作为中间的区间，假设范围是`[i:i+k-1]`, 那么我们只需要求在`[0:i-1]`内最大的长度为k的区间，以及在`[i+1:n-1]`内最大的长度为k的区间。这两个分量都是可以提前计算好的。我们只要在数组上从前往后跑一遍长度为k的滑窗，就可以记录任意前缀里曾经出现过的最大值，记做`leftMax[i]`；同理，在数组上从后往前跑一遍长度为k的滑窗，就可以记录任意后缀里曾经出现过的最大值，记做`rightMax[i]`。所以我们只要找到全局最大的`leftMax[i-1] + midSum[i:i+k-1] + rightMax[i+k]`即可。
    * *wisdompeak*: 要注意一个细节，因为题意要求，如果有多个总和相同的解，取index位置最小的解。所以我们从左往右滑动L段的时候，只有在`leftMax`**大于**历史最大值的时候才更新`leftIdx`，这样在相同的`leftMax`的时候我们保留的是较小的index。我们在从右往左滑动R段的时候，当`rightMax`**大于等于**历史最大值，就可以更新`rightIdx`，这样在相同的`rightMax`的时候我们保留的是较小的index。我们在从左往右滑动M段的时候，也是只有`leftMax[i-1] + midSum[i:i+k-1] + rightMax[i+k]`**大于**历史最大值的时候才更新输出结果为`{leftIdx[i-1], i, rightIdx[i+k]}`

* **具体实现: Sliding Window**

把nums分成 L, M, R三段:
```
{X X X X X }   X X X   {X  X X X X }
         i-1   i       i+k
```
使用两个`DP`思想的arrays来分别记录L段和R段的信息: `leftMax`, `rightMax`:

**物理意义**:

    leftMax[i]: the largest k-length subarray sum from [0:i]
    rightMax[i]: the largest k-length subarray sum from [i:n-1]

**Objective**: Find index `i` s.t. `max(leftMax[i-1] + midSum[i:i+k-1] + rightMax[i+k])`

**注意**:
	
    leftMax[i]中 i 代表L段的右端点
    rightMax[i]中 i 代表R段的左端点

**注意**:

	L段 从左往右滑动滑窗从而更新 leftMax
	R段 从右往左滑动滑窗从而更新 rightMax

* **Trick**: Use Prefix-sum to compute subarray sum

`subarray-sum[l:r] = prefix-sum[r] - prefix-sum[l-1]`


Time complexity = $O(3n)$ = $O(n)$


## Resource
[【每日一题】LeetCode 689. Maximum Sum of 3 Non-Overlapping Subarrays](https://www.youtube.com/watch?v=oo_T4GdRot4&t=1705s&ab_channel=HuifengGuan)