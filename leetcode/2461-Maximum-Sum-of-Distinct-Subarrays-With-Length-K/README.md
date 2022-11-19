# [2461. Maximum Sum of Distinct Subarrays With Length K](https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/)

## Solution idea

### Sliding Window
* 一看题目是要求定长为k的subarray，用**模版Flex**解题比较方便

* 更新窗口内数据:
    1. 更新 sum
    2. count的变动依据如下：(镜面对称)
```go
window[rightElement]++
if window[rightElement] == 1 {
    count--
}
```

```go
if window[leftElement] == 1 {
    count--
}
window[leftElement]--
```

* 更新result:
    1. 时机: 每次窗口收缩前
    2. 条件: `count == k`


Time complexity = $O(n)$

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Two_Pointers/2461.Maximum-Sum-of-Distinct-Subarrays-With-Length-K)

