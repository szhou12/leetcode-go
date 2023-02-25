# [2163. Minimum Difference in Sums After Removal of Elements](https://leetcode.com/problems/minimum-difference-in-sums-after-removal-of-elements/description/)

## Solution idea

### 3-Pass + Priority Queue

#### 要点总结



#### 物理意义
*Note: `n=len(nums)/3`*

*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*

```
leftMin[i] := min sum of n-len subseq in nums[0:i] = pick n smallest elements in nums[0:i] by using max heap
rightMax[i] := max sum of n-len subseq in nums[i:3n-1] = pick n largest elements in nums[i:3n-1] by using min heap
```



#### 代码整体结构总结

## Resource

[【每日一题】LeetCode 2163. Minimum Difference in Sums After Removal of Elements](https://www.youtube.com/watch?v=Fdsdm8rbA9E&ab_channel=HuifengGuan)