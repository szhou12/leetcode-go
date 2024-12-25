# [2221. Find Triangular Sum of an Array](https://leetcode.com/problems/find-triangular-sum-of-an-array/description/)

## Solution idea
### Simulation
因为数据量不大(<1000)，按照题意进行全模拟就行：每个`nums[i]`被`(nums[i]+nums[i+1])%10`替换，一共进行n-1轮，最后返回头元素。

Time complexity = $O(n^2)$

### Pascal's Triangle
题目设置的计算机制类似于杨辉三角，只是把杨辉三角倒过来。

如果把杨辉三角倒过来：
```
1  4  6  4  1
  1  3  3  1
    1  2  1
      1  1
        1
```
我们可以做如下理解:
- 最底层的1，由`1 1`衍生而来，并且两个1各自的贡献比例为1:1
- `1 1`由`1 2 1`衍生而来。`1 1`中左侧的1由`[1 2] 1`贡献；右侧的1由`1 [2 1]`贡献。注意到2贡献了两次(复制了两份)，所以三个元素的贡献比例为1:2:1。由传递性，`1 2 1`对最底层的1的贡献比例即为1:2:1。
- 以此类推，可以得出结论：第i行所有元素对最底层的元素的**贡献比例**为第i行的杨辉三角的值。

所以，最后的结果就是`nums[]`所有元素的加权取平均。而每个元素的权重就是$(a+b)^{n-1}$ (`n=len(nums)`)的系数，也就是杨辉三角中第n-1行的值。

Time complexity = $O(n^2)$

## Resource
[【每日一题】LeetCode 2221. Find Triangular Sum of an Array](https://www.youtube.com/watch?v=RzoHl7M9xvM&ab_channel=HuifengGuan)