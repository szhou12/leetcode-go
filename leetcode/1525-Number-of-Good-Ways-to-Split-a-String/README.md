# [1525. Number of Good Ways to Split a String](https://leetcode.com/problems/number-of-good-ways-to-split-a-string/description/)

## Solution idea

### 3-Pass

#### 要点总结
1. 思路相对比较自然的题: 每个分界点, 左手边看有多少distinct characters, 右手边看有多少distinct characters
2. 用 `HashMap` 计算distinct characters

#### 物理意义
*Note: `n=len(s)`*

*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*
```
left[i] := # of distinct chars in s[0:i]
right[i] := # of distinct chars in s[i:n-1]
```

#### 代码整体结构总结
```
Step 1: 构造 left[]: 维护一个hashmap 从左至右 计算distinct characters个数
Step 2: 构造 right[]: 维护一个hashmap 从右至左 计算distinct characters个数
Step 3: iterate 分界点 i, 如果 left[i]==right[i+1]: +1
```

Time complexity = $O(n)$

## Resource
[【每日一题】1525. Number of Good Ways to Split a String, 3/8/2021](https://www.youtube.com/watch?v=NY5YdoyV6tU&ab_channel=HuifengGuan)