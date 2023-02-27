# [2163. Minimum Difference in Sums After Removal of Elements](https://leetcode.com/problems/minimum-difference-in-sums-after-removal-of-elements/description/)

## Solution idea

### 3-Pass + Priority Queue

#### 要点总结
1. 找分界点分成L段和R段。L段挑n个元素，R段挑n个元素。这是什么？不就是长度为n的sub-sequence吗. 又因为题目条件是每个元素都是整数，题目求的是L段sum-R段sum (注意: 不是差的绝对值)，所以，L段的n个元素尽可能挑小的，R段的n个元素尽可能挑大的。


#### 物理意义
*Note: `n=len(nums)/3`*

*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*

```
leftMin[i] := min sum of n-len subseq in nums[0:i] = pick n smallest elements in nums[0:i] by using max heap
rightMax[i] := max sum of n-len subseq in nums[i:3n-1] = pick n largest elements in nums[i:3n-1] by using min heap
```



#### 代码整体结构总结
```
Step 1: 构造 leftMin[]: index i 从左往右, 用**MaxHeap**保留n个最小元素
Step 2: 构造 rightMax[]: index i 从右往左, 用**MinHeap**保留n个最大元素
Step 3: iterate 分界点, 取 min(leftMin[i] - rightMax[i+1])
```

## Resource

[【每日一题】LeetCode 2163. Minimum Difference in Sums After Removal of Elements](https://www.youtube.com/watch?v=Fdsdm8rbA9E&ab_channel=HuifengGuan)