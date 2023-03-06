# [334. Increasing Triplet Subsequence](https://leetcode.com/problems/increasing-triplet-subsequence/description/)

## Solution idea

### 3-Pass

#### 要点总结
1. 3-Pass类型题里面比较容易上手的题目: 以中间的元素向左右两边看, 如果满足: 左边L段最小元素 < 中间元素 < 右边R段最大元素; 那么就存在题目要求的triplet

#### 物理意义

*Note: `n=len(nums)`*

*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*

```go
// Definition
leftMin[i] := min element in nums[0:i]
// Base Case
leftMin[0] = nums[0]
// Recurrence
leftMin[i] = min(nums[i], leftMin[i-1])

// Definition
rightMax[i] := max element in nums[i:n-1]
// Base Case
rightMax[n-1] = nums[n-1]
// Recurrence
rightMax[i] = max(nums[i], rightMax[i+1])
```

#### 代码整体结构总结
```
Step 1: 构造 leftMin[]: DP方法 从左至右 总是保留最小值
Step 2: 构造 rightMax[]: DP方法 从右至左 总是保留最大值
Step 3: iterate 分界点, 判断 是否存在 leftMin[i-1] < nums[i] < rightMax[i+1]
```

Time complexity = $O(n)$

## Resource
[【每日一题】334. Increasing Triplet Subsequence, 12/18/2020](https://www.youtube.com/watch?v=-wtypYo-K-o&ab_channel=HuifengGuan)