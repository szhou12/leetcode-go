# [1918. Kth Smallest Subarray Sum](https://leetcode.ca/2021-08-01-1918-Kth-Smallest-Subarray-Sum/)

**Locked Question**

## Solution idea

### Binary Search - Guess k-th element + Two Pointers

* **Key** 转化思想: subarray sum = prefix-sum diff

e.g. `sum[i:j]` as the subarray sum for `nums` from `i` to `j` (左闭右闭)
```
sum[i:j] = prefixSum[j] - prefixSum[i-1]
```

* prefix-sum 的特性: **non-decreasing order** $\Rightarrow $ 这个特性使得可以用 **two-pointers**

* 为什么 prefix-sum array 要比 `nums` length + 1?
    * 方便计算 `i=0` 为start的 subarray sum. 也就是说, 否则计算diff时，无法包括前缀是`nums`第一个元素的diff
    * e.g. `sum[0:2] = prefixSum[2] - prefixSum[0-1] = prefixSum[2] - prefixSum[-1]` 所以 `prefixSum` 在头部多垫一个格子放0, 使得 `nums[i]` 对应 `prefixSum[i+1]` so that `sum[i:j] = prefixSum[j+1] - prefixSum[i+1-1] = prefixSum[j+1] - prefixSum[i]`

* prefix-sum array例子:

```
nums = [2, 1, 3]

prefixSum = [0, 2, 3, 6]

sum[0:2] = 2+1+3 = 6 = prefixSum[3] - prefixSum[1-1] =  prefixSum[3] - prefixSum[0]
```

Time complexity = $O(n\log D)$ where $D = $ difference between the largest subarray sum of `nums` - smallest subarray sum of `nums`

## Resource
[【LeetCode】1918. Kth Smallest Subarray Sum](https://www.bilibili.com/video/BV1Vb4y1h7fW/?spm_id_from=333.337.search-card.all.click&vd_source=0c02ef6f6e7a2b0959d7dd28e9e49da4)

